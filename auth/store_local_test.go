package auth

import (
	"testing"
	"time"

	"github.com/ggymm/gopkg/logger"
)

func newLocalStore() *LocalStore {
	s, err := newLocal(LocalConfig{
		Dir: "C:/Product/gopkg/temp/auth/storage",
	}, logger.InitCustom("C:/Product/gopkg/temp/auth/auth.log"))
	if err != nil {
		panic(err)
		return nil
	}
	return s
}

func TestLocalStore_Get_Put_Update(t *testing.T) {
	var (
		err   error
		key   = []byte("key")
		value = []byte("value")
	)
	s := newLocalStore()

	// 保存数据
	err = s.Put(key, value, NeverExpire)
	if err != nil {
		t.Error("put value error", err)
		return
	}
	value, err = s.Get(key)
	if err != nil {
		t.Error("get value error", err)
		return
	}
	if string(value) != "value" {
		t.Error("get value error data not match")
		return
	}
	t.Log("get value success")

	// 更新数据
	err = s.Update(key, []byte("value-update"))
	if err != nil {
		t.Error("update value error", err)
		return
	}

	// 获取数据
	value, err = s.Get(key)
	if err != nil {
		t.Error("get value error", err)
		return
	}
	if string(value) != "value-update" {
		t.Error("update value error data not match")
		return
	}
	t.Log("update value success")
}

func TestLocalStore_Delete(t *testing.T) {
	var (
		err   error
		key   = []byte("key")
		value = []byte("value")
	)
	s := newLocalStore()

	// 保存数据
	err = s.Put(key, value, NeverExpire)
	if err != nil {
		t.Error("put value error", err)
		return
	}

	// 删除数据
	err = s.Delete(key)
	if err != nil {
		t.Error("delete value error", err)
		return
	}

	// 获取数据
	value, err = s.Get(key)
	if err != nil {
		t.Error("get value error", err)
		return
	}
	if value != nil {
		t.Error("delete value error data is not nil")
		return
	}
	t.Log("delete value success")
}

func TestLocalStore_Expire(t *testing.T) {
	var (
		err   error
		key   = []byte("key")
		value = []byte("value")
	)
	s := newLocalStore()

	// 保存数据
	err = s.Put(key, value, time.Duration(5)*time.Second)
	if err != nil {
		t.Error("put value error", err)
		return
	}

	// 获取数据
	value, err = s.Get(key)
	if err != nil {
		t.Error("get value error", err)
		return
	}
	if string(value) != "value" {
		t.Error("get value error data not match")
		return
	}
	t.Log("get value success")

	// 等待过期后，再次获取数据
	time.Sleep(6 * time.Second)
	value, err = s.Get(key)
	if err != nil {
		t.Error("get value error", err)
		return
	}
	if value != nil {
		t.Error("expire value error data is not nil")
		return
	}
	t.Log("expire value success")
}

func TestLocalStore_UpdateTimeout(t *testing.T) {
	var (
		err   error
		key   = []byte("key")
		value = []byte("value")
	)
	s := newLocalStore()

	// 保存数据
	err = s.Put(key, value, time.Duration(5)*time.Second)
	if err != nil {
		t.Error("put value error", err)
		return
	}

	// 获取数据
	value, err = s.Get(key)
	if err != nil {
		t.Error("get value error", err)
		return
	}
	if string(value) != "value" {
		t.Error("get value error data not match")
		return
	}

	// 更新过期时间
	time.Sleep(4 * time.Second)
	err = s.UpdateTimeout(key, time.Duration(5)*time.Second)
	if err != nil {
		t.Error("update value timeout error", err)
		return
	}

	// 在数据过期前获取数据
	time.Sleep(2 * time.Second)
	value, err = s.Get(key)
	if err != nil {
		t.Error("get value error", err)
		return
	}
	if value == nil {
		t.Error("update timeout error data is nil")
		return
	}

	// 等待过期后，再次获取数据
	time.Sleep(5 * time.Second)
	value, err = s.Get(key)
	if err != nil {
		t.Error("get value error", err)
		return
	}
	if value != nil {
		t.Error("update timeout error data is not nil")
		return
	}
	t.Log("update timeout success")
}
