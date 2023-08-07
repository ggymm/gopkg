package auth

import (
	"github.com/ggymm/gopkg/log"
	"testing"
	"time"
)

func newRoseDBStore() *RoseDBStore {
	s, err := newRoseDB(RoseDBConfig{
		DirPath: "C:/Product/gopkg/temp/auth/storage",
	}, log.InitCustom("C:/Product/gopkg/temp/auth/auth.log"))
	if err != nil {
		panic(err)
		return nil
	}
	return s
}

func TestRoseDBStore_Get_Put_Update(t *testing.T) {
	var (
		err   error
		key   = []byte("key")
		value = []byte("value")
	)
	s := newRoseDBStore()

	// ====== 测试保存、获取数据 ======
	err = s.Put(key, value, NeverExpire)
	if err != nil {
		t.Error(err)
		return
	}
	value, err = s.Get([]byte("key"))
	if err != nil {
		t.Error(err)
		return
	}
	if string(value) != "value" {
		t.Error("get value error")
		return
	}
	t.Log("get value success")

	// ====== 测试更新数据 ======
	err = s.Update(key, []byte("value-update"))
	if err != nil {
		t.Error(err)
		return
	}
	value, err = s.Get([]byte("key"))
	if err != nil {
		t.Error(err)
		return
	}
	if string(value) != "value-update" {
		t.Error("update value error")
		return
	} else {
		t.Log("update value success")
	}
}

func TestRoseDBStore_Delete(t *testing.T) {
	var (
		err   error
		key   = []byte("key")
		value = []byte("value")
	)
	s := newRoseDBStore()

	// 保存数据
	err = s.Put(key, value, NeverExpire)
	if err != nil {
		t.Error(err)
		return
	}

	// ====== 测试删除数据 ======
	err = s.Delete(key)
	if err != nil {
		t.Error(err)
		return
	}

	// ====== 验证是否删除成功 ======
	value, err = s.Get([]byte("key"))
	if err != nil {
		t.Error(err)
		return
	}
	if value != nil {
		t.Error("delete value error")
		return
	}
	t.Log("delete value success")
}

func TestRoseDBStore_Expire(t *testing.T) {
	var (
		err   error
		key   = []byte("key")
		value = []byte("value")
	)
	s := newRoseDBStore()

	// 保存数据
	err = s.Put(key, value, 5)
	if err != nil {
		t.Error(err)
		return
	}

	// ====== 验证是否保存成功 ======
	value, err = s.Get([]byte("key"))
	if err != nil {
		t.Error(err)
		return
	}
	if string(value) != "value" {
		t.Error("get value error")
		return
	}
	t.Log("get value success")

	// ====== 验证是否过期 ======
	time.Sleep(6 * time.Second)
	value, err = s.Get([]byte("key"))
	if err != nil {
		t.Error(err)
		return
	}
	if value != nil {
		t.Error("expire value error")
		return
	}
	t.Log("expire value success")
}
