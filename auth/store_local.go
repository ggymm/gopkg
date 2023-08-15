package auth

import (
	"time"

	"github.com/nutsdb/nutsdb"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type LocalStore struct {
	log zerolog.Logger

	db     *nutsdb.DB
	bucket string
}

type LocalConfig struct {
	Dir string
}

func newLocal(config LocalConfig, log zerolog.Logger) (store *LocalStore, err error) {
	store = &LocalStore{
		log: log,
	}

	store.db, err = nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(config.Dir), // 数据库会自动创建这个目录文件
	)
	store.bucket = "bucket-auth"
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (store *LocalStore) errLog(err error) *zerolog.Event {
	return store.log.Error().Stack().Err(errors.WithStack(err))
}

func (store *LocalStore) errIgnore(err error) bool {
	return errors.Is(err, nutsdb.ErrNotFoundKey) ||
		errors.Is(err, nutsdb.ErrNotFoundBucket) ||
		errors.Is(err, nutsdb.ErrKeyNotFound)
}

func (store *LocalStore) Get(key []byte) ([]byte, error) {
	var value []byte
	err := store.db.View(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get(store.bucket, key)
		if err != nil {
			if store.errIgnore(err) {
				return nil
			}
			store.errLog(err).Msg(storeGetError)
			return err
		}
		value = entry.Value
		return nil
	})
	if err != nil {
		store.errLog(err).Msg(storeGetError)
		return nil, err
	}
	return value, nil
}

func (store *LocalStore) Put(key []byte, value []byte, timeout time.Duration) error {
	if timeout == 0 || timeout < NotValueExpire {
		return errors.New(InvalidTimeout)
	}

	return store.db.Update(func(tx *nutsdb.Tx) error {
		if timeout == NeverExpire {
			err := tx.Put(store.bucket, key, value, nutsdb.Persistent)
			if err != nil {
				store.errLog(err).Msg(storePutError)
				return err
			}
			return nil
		} else {
			err := tx.Put(store.bucket, key, value, uint32(timeout.Seconds()))
			if err != nil {
				store.errLog(err).Msg(storePutError)
				return err
			}
			return nil
		}
	})
}

func (store *LocalStore) Delete(key []byte) error {
	return store.db.Update(func(tx *nutsdb.Tx) error {
		err := tx.Delete(store.bucket, key)
		if err != nil {
			if store.errIgnore(err) {
				return nil
			}
			store.errLog(err).Msg(storeDeleteError)
			return err
		}
		err = tx.Delete(store.bucket, key)
		if err != nil {
			store.errLog(err).Msg(storeDeleteError)
			return err
		}
		return nil
	})
}

func (store *LocalStore) Update(key []byte, value []byte) error {
	return store.db.Update(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get(store.bucket, key)
		if err != nil {
			if store.errIgnore(err) {
				return nil
			}
			store.errLog(err).Msg(storeUpdateError)
			return err
		}
		err = tx.Put(store.bucket, key, value, entry.Meta.TTL)
		if err != nil {
			store.errLog(err).Msg(storeUpdateError)
			return err
		}
		return nil
	})
}

func (store *LocalStore) CheckTimeout(key []byte) (time.Duration, error) {
	// 秒 -> time.Duration
	var timeout time.Duration
	err := store.db.View(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get(store.bucket, key)
		if err != nil {
			if store.errIgnore(err) {
				return nil
			}
			store.errLog(err).Msg(storeGetError)
			return err
		}

		timeout = time.Unix(int64(entry.Meta.TTL), 0).Sub(time.Now())
		return nil
	})
	return timeout, err
}

func (store *LocalStore) UpdateTimeout(key []byte, timeout time.Duration) error {
	if timeout == 0 || timeout < NotValueExpire {
		return errors.New(InvalidTimeout)
	}

	return store.db.Update(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get(store.bucket, key)
		if err != nil {
			if store.errIgnore(err) {
				return nil
			}
			store.errLog(err).Msg(storeGetError)
			return err
		}

		if timeout == NeverExpire {
			err = tx.Put(store.bucket, key, entry.Value, nutsdb.Persistent)
			if err != nil {
				store.errLog(err).Msg(storePutError)
				return err
			}
			return nil
		} else {
			err = tx.Put(store.bucket, key, entry.Value, uint32(timeout.Seconds()))
			if err != nil {
				store.errLog(err).Msg(storePutError)
				return err
			}
			return nil
		}
	})
}
