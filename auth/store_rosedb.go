package auth

import (
	"github.com/ggymm/gopkg/utils/cast"
	"github.com/ggymm/rosedb"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"strings"
	"time"

	"github.com/ggymm/gopkg/utils"
)

type RoseDBStore struct {
	log zerolog.Logger

	data   *rosedb.DB
	expire *rosedb.DB
}

type RoseDBConfig struct {
	DirPath string
}

func newRoseDB(c RoseDBConfig, log zerolog.Logger) (store *RoseDBStore, err error) {
	store = &RoseDBStore{
		log: log,
	}
	if strings.HasSuffix(c.DirPath, "/") {
		c.DirPath = c.DirPath[:len(c.DirPath)-1]
	}

	ops := rosedb.DefaultOptions
	ops.DirPath = c.DirPath
	store.data, err = rosedb.Open(ops)
	if err != nil {
		return nil, err
	}

	ops = rosedb.DefaultOptions
	ops.DirPath = c.DirPath + "-expire"
	store.expire, err = rosedb.Open(ops)

	// 定时清理过期数据
	go func() {
		// 遍历 expire 数据库
		iterOptions := rosedb.DefaultIteratorOptions
		iter := store.expire.NewIterator(iterOptions)
		defer iter.Close()
		for {

			for ; iter.Valid(); iter.Next() {
				key := iter.Key()

				// 清理过期数据
				store.clearTimeout(key)
			}

			// 每隔 1 分钟执行一次清理任务
			time.Sleep(time.Minute)
		}
	}()
	return store, err
}

func (r *RoseDBStore) errLog(err error) *zerolog.Event {
	return r.log.Error().Stack().Err(errors.WithStack(err))
}

func (r *RoseDBStore) checkTimeout(key []byte) int64 {
	r.clearTimeout(key)

	expire, err := r.expire.Get(key)
	if err != nil || expire == nil {
		if !errors.Is(err, rosedb.ErrKeyNotFound) { // 不打印找不到值的错误
			r.errLog(err).Msg(GetExpireError)
		}
		return NotValueExpire
	}

	timeout := cast.ToInt64(expire)
	if timeout == NeverExpire {
		return NeverExpire
	}

	// 计算剩余过期时间
	timeout = timeout - utils.CurrentTimestamp()
	if timeout < 0 {
		_ = r.data.Delete(key)
		_ = r.expire.Delete(key)
		return NotValueExpire
	}
	return timeout
}

func (r *RoseDBStore) clearTimeout(key []byte) {
	expire, err := r.expire.Get(key)
	if err != nil || expire == nil {
		if !errors.Is(err, rosedb.ErrKeyNotFound) { // 不打印找不到值的错误
			r.errLog(err).Msg(GetExpireError)
		}
		return
	}

	// 清除数据
	// 非永不过期的数据，且已过期
	timeout := cast.ToInt64(expire)
	if timeout != NeverExpire && timeout < utils.CurrentTimestamp() {
		_ = r.data.Delete(key)
		_ = r.expire.Delete(key)
	}
}

func (r *RoseDBStore) Get(key []byte) ([]byte, error) {
	if r.checkTimeout(key) == NotValueExpire {
		return nil, nil
	}
	value, err := r.data.Get(key)
	if err != nil {
		r.errLog(err).Msg(GetValueError)
		return nil, errors.New(StoreExecGetError)
	}
	return value, nil
}

func (r *RoseDBStore) Put(key []byte, value []byte, timeout int64) (err error) {
	if timeout == 0 || timeout < NotValueExpire {
		return errors.New(InvalidTimeout)
	}

	// 保存数据
	err = r.data.Put(key, value)
	if err != nil {
		r.errLog(err).Msg(PutValueError)
		return errors.New(StoreExecPutError)
	}

	// 保存过期时间
	if timeout != NeverExpire {
		timeout = utils.CurrentTimestamp() + timeout*1000
	}
	err = r.expire.Put(key, cast.ToBytes(timeout))
	if err != nil {
		r.errLog(err).Msg(PutExpireError)
		return errors.New(StoreExecPutError)
	}
	return nil
}

func (r *RoseDBStore) Update(key []byte, value []byte) error {
	if r.checkTimeout(key) == NotValueExpire {
		return nil
	}
	err := r.data.Put(key, value)
	if err != nil {
		r.errLog(err).Msg(PutValueError)
		return errors.New(StoreExecUpdateError)
	}
	return nil
}

func (r *RoseDBStore) Delete(key []byte) (err error) {
	err = r.data.Delete(key)
	if err != nil {
		r.errLog(err).Msg(DeleteValueError)
		return errors.New(StoreExecDeleteError)
	}
	err = r.expire.Delete(key)
	if err != nil {
		r.errLog(err).Msg(DeleteExpireError)
		return errors.New(StoreExecDeleteError)
	}
	return nil
}

func (r *RoseDBStore) CheckTimeout(key []byte) (int64, error) {
	return r.checkTimeout(key), nil
}

func (r *RoseDBStore) UpdateTimeout(key []byte, timeout int64) (err error) {
	if timeout != NeverExpire {
		timeout = utils.CurrentTimestamp() + timeout*1000
	}
	err = r.expire.Put(key, cast.ToBytes(timeout))
	if err != nil {
		r.errLog(err).Msg(PutExpireError)
		return errors.New(StoreExecUpdateError)
	}
	return nil
}
