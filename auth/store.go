package auth

type storeType int

const (
	Redis storeType = iota
	RoseDB
)

const (
	NeverExpire    = -1 // NeverExpire 永不过期
	NotValueExpire = -2 // NotValueExpire 没有值过期
)

const (
	StoreExecGetError    = "store get error"
	StoreExecPutError    = "store put error"
	StoreExecUpdateError = "store update error"
	StoreExecDeleteError = "store delete error"
)

type store interface {
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte, timeout int64) error

	Delete(key []byte) error
	Update(key []byte, value []byte) error

	CheckTimeout(key []byte) (int64, error)
	UpdateTimeout(key []byte, timeout int64) error
}
