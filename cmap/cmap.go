package cmap

import (
	"fmt"
	"sync"
)

var ShardCount = 32

type Stringer interface {
	fmt.Stringer
	comparable
}

// CMap A "thread" safe map of type string:Anything.
// To avoid lock bottlenecks this map is dived to several (ShardCount) map shards.
type CMap[K comparable, V any] struct {
	shards   []*Shared[K, V]
	sharding func(key K) uint32
}

// Shared A "thread" safe string to anything map.
type Shared[K comparable, V any] struct {
	items        map[K]V
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func create[K comparable, V any](sharding func(key K) uint32) CMap[K, V] {
	m := CMap[K, V]{
		sharding: sharding,
		shards:   make([]*Shared[K, V], ShardCount),
	}
	for i := 0; i < ShardCount; i++ {
		m.shards[i] = &Shared[K, V]{items: make(map[K]V)}
	}
	return m
}

// New Creates a new concurrent map.
func New[V any]() CMap[string, V] {
	return create[string, V](fnv32)
}

// GetShard returns shard under given key
func (m CMap[K, V]) GetShard(key K) *Shared[K, V] {
	return m.shards[uint(m.sharding(key))%uint(ShardCount)]
}

func (m CMap[K, V]) MSet(data map[K]V) {
	for key, value := range data {
		shard := m.GetShard(key)
		shard.Lock()
		shard.items[key] = value
		shard.Unlock()
	}
}

// Set Sets the given value under the specified key.
func (m CMap[K, V]) Set(key K, value V) {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// UpsertCb Callback to return new element to be inserted into the map
// It is called while lock is held, therefore it MUST NOT
// try to access other keys in same map, as it can lead to deadlock since
// Go sync.RWLock is not reentrant
type UpsertCb[V any] func(exist bool, valueInMap V, newValue V) V

// Upsert Insert or Update - updates existing element or inserts a new one using UpsertCb
func (m CMap[K, V]) Upsert(key K, value V, cb UpsertCb[V]) (res V) {
	shard := m.GetShard(key)
	shard.Lock()
	v, ok := shard.items[key]
	res = cb(ok, v, value)
	shard.items[key] = res
	shard.Unlock()
	return res
}

// SetIfAbsent Sets the given value under the specified key if no value was associated with it.
func (m CMap[K, V]) SetIfAbsent(key K, value V) bool {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	_, ok := shard.items[key]
	if !ok {
		shard.items[key] = value
	}
	shard.Unlock()
	return !ok
}

// Get retrieves an element from map under given key.
func (m CMap[K, V]) Get(key K) (V, bool) {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// Count returns the number of elements within the map.
func (m CMap[K, V]) Count() int {
	count := 0
	for i := 0; i < ShardCount; i++ {
		shard := m.shards[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// Has Looks up an item under specified key
func (m CMap[K, V]) Has(key K) bool {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// See if element is within shard.
	_, ok := shard.items[key]
	shard.RUnlock()
	return ok
}

// Remove removes an element from the map.
func (m CMap[K, V]) Remove(key K) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// RemoveCb is a callback executed in a map.RemoveCb() call, while Lock is held
// If returns true, the element will be removed from the map
type RemoveCb[K any, V any] func(key K, v V, exists bool) bool

// RemoveCb locks the shard containing the key, retrieves its current value and calls the callback with those params
// If callback returns true and element exists, it will remove it from the map
// Returns the value returned by the callback (even if element was not present in the map)
func (m CMap[K, V]) RemoveCb(key K, cb RemoveCb[K, V]) bool {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	v, ok := shard.items[key]
	remove := cb(key, v, ok)
	if remove && ok {
		delete(shard.items, key)
	}
	shard.Unlock()
	return remove
}

// Pop removes an element from the map and returns it
func (m CMap[K, V]) Pop(key K) (v V, exists bool) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	v, exists = shard.items[key]
	delete(shard.items, key)
	shard.Unlock()
	return v, exists
}

// IsEmpty checks if map is empty.
func (m CMap[K, V]) IsEmpty() bool {
	return m.Count() == 0
}

// Tuple Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type Tuple[K comparable, V any] struct {
	Key K
	Val V
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m CMap[K, V]) IterBuffered() <-chan Tuple[K, V] {
	chs := snapshot(m)
	total := 0
	for _, c := range chs {
		total += cap(c)
	}
	ch := make(chan Tuple[K, V], total)
	go fanIn(chs, ch)
	return ch
}

// Clear removes all items from map.
func (m CMap[K, V]) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// Returns an array of channels that contains elements in each shard,
// which likely takes a snapshot of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
func snapshot[K comparable, V any](m CMap[K, V]) (chs []chan Tuple[K, V]) {
	// When you access map items before initializing.
	if len(m.shards) == 0 {
		panic(`CMap is not initialized. Should run New() before usage.`)
	}
	chs = make([]chan Tuple[K, V], ShardCount)
	wg := sync.WaitGroup{}
	wg.Add(ShardCount)
	// Foreach shard.
	for index, shard := range m.shards {
		go func(index int, shard *Shared[K, V]) {
			// Foreach key, value pair.
			shard.RLock()
			chs[index] = make(chan Tuple[K, V], len(shard.items))
			wg.Done()
			for key, val := range shard.items {
				chs[index] <- Tuple[K, V]{key, val}
			}
			shard.RUnlock()
			close(chs[index])
		}(index, shard)
	}
	wg.Wait()
	return chs
}

// fanIn reads elements from channels `chs` into channel `out`
func fanIn[K comparable, V any](chs []chan Tuple[K, V], out chan Tuple[K, V]) {
	wg := sync.WaitGroup{}
	wg.Add(len(chs))
	for _, ch := range chs {
		go func(ch chan Tuple[K, V]) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}

// Items returns all items as map[string]V
func (m CMap[K, V]) Items() map[K]V {
	tmp := make(map[K]V)

	// Insert items to temporary map.
	for item := range m.IterBuffered() {
		tmp[item.Key] = item.Val
	}

	return tmp
}

// IterCb Iterator callback for every key,value found in
// maps. RLock is held for all calls for a given shard
// therefore callback sess consistent view of a shard,
// but not across the shards
type IterCb[K comparable, V any] func(key K, v V)

// IterCb Callback based iterator, cheapest way to read
// all elements in a map.
func (m CMap[K, V]) IterCb(fn IterCb[K, V]) {
	for idx := range m.shards {
		shard := (m.shards)[idx]
		shard.RLock()
		for key, value := range shard.items {
			fn(key, value)
		}
		shard.RUnlock()
	}
}

// Keys returns all keys as []string
func (m CMap[K, V]) Keys() []K {
	count := m.Count()
	ch := make(chan K, count)
	go func() {
		// Foreach shard.
		wg := sync.WaitGroup{}
		wg.Add(ShardCount)
		for _, shard := range m.shards {
			go func(shard *Shared[K, V]) {
				// Foreach key, value pair.
				shard.RLock()
				for key := range shard.items {
					ch <- key
				}
				shard.RUnlock()
				wg.Done()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()

	// Generate keys
	keys := make([]K, 0, count)
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}
