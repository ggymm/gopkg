package auth

import (
	"github.com/rs/zerolog"
	"time"
)

type RedisStore struct {
	log zerolog.Logger
}

type RedisConfig struct {
	Addr     string
	Password string
	Database int
}

func newRedis(config RedisConfig, log zerolog.Logger) (store *RedisStore, err error) {
	store = &RedisStore{
		log: log,
	}
	return store, err
}

func (r *RedisStore) Get(key []byte) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisStore) Put(key []byte, value []byte, timeout time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisStore) Delete(key []byte) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisStore) Update(key []byte, value []byte) error {
	//TODO implement me
	panic("implement me")
}

func (r *RedisStore) CheckTimeout(key []byte) (time.Duration, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisStore) UpdateTimeout(key []byte, timeout time.Duration) error {
	//TODO implement me
	panic("implement me")
}
