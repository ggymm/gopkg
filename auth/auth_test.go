package auth

import (
	"math"
	"testing"
	"time"
)

func InitAuth() {
	err := Init(Config{
		LogPath: "C:/Product/gopkg/temp/auth/auth.log",

		Store: RoseDB,
		RoseDBConfig: RoseDBConfig{
			DirPath: "C:/Product/gopkg/temp/auth/storage",
		},

		Concurrent:    true,
		ShareToken:    false,
		MaxLoginCount: math.MaxInt,

		TokenPrefix:    "ninelock-token",
		AutoRenewToken: true,
	})
	if err != nil {
		panic(err)
		return
	}
}

func TestAuth_Login(t *testing.T) {
	var (
		id     int64 = 1
		device       = "mobile"
	)

	InitAuth()

	token, err := auth.Login(id, LoginConfig{
		Device:  device,
		Timeout: NeverExpire,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("token:", token)

	t.Log("login status:", auth.CheckLogin(id))
	t.Log("login status:", auth.CheckToken(token))
}

func TestAuth_Login_Timeout(t *testing.T) {
	var (
		id     int64 = 1
		device       = "mobile"
	)

	InitAuth()

	token, err := auth.Login(id, LoginConfig{
		Device:  device,
		Timeout: 5,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("token:", token)

	time.Sleep(6 * time.Second)

	t.Log("login status:", auth.CheckLogin(id))
	t.Log("login status:", auth.CheckToken(token))
}

func TestAuth_Login_RenewToken(t *testing.T) {
	var (
		id     int64 = 1
		device       = "mobile"
	)

	InitAuth()

	token, err := auth.Login(id, LoginConfig{
		Device:  device,
		Timeout: 5,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("token:", token)

	// 休眠 3 秒，续签 token
	time.Sleep(3 * time.Second)
	t.Log("login status:", auth.CheckToken(token))

	// 休眠 3 秒，检查 token 是否过期（没过期）
	time.Sleep(3 * time.Second)
	// 休眠 6 秒，检查 token 是否过期（已经过期）
	// time.Sleep(6 * time.Second)
	t.Log("login status:", auth.CheckLogin(id))
	t.Log("login status:", auth.CheckToken(token))
}
