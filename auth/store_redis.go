package auth

import "github.com/rs/zerolog"

type RedisStore struct {
	log zerolog.Logger
}

type RedisConfig struct {
	Addr     string
	Password string
	Database int
}

func newRedis(c RedisConfig, log zerolog.Logger) (store *RedisStore, err error) {
	store = &RedisStore{
		log: log,
	}
	return store, err
}

func (r *RedisStore) Get(key []byte) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisStore) Put(key []byte, value []byte, timeout int64) error {
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

func (r *RedisStore) CheckTimeout(key []byte) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisStore) UpdateTimeout(key []byte, timeout int64) error {
	//TODO implement me
	panic("implement me")
}
