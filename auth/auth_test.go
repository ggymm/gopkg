package auth

import (
	"math"
	"testing"
	"time"
)

func InitAuth() {
	err := Init(Config{
		LogPath: "C:/Product/gopkg/temp/auth/auth.log",

		Store: Local,
		LocalConfig: LocalConfig{
			Dir: "C:/Product/gopkg/temp/auth/storage",
		},

		Concurrent:    true,
		ShareToken:    false,
		MaxLoginCount: math.MaxInt,

		TokenName:      "ninelock-token",
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

	token, err := Auth.Login(id, LoginConfig{
		Device:  device,
		Timeout: NeverExpire,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("token:", token)

	var status bool
	status, err = Auth.CheckLogin(id)
	t.Log("login status:", status, err)

	status, err = Auth.CheckToken(token)
	t.Log("login status:", status, err)
}

func TestAuth_Login_Timeout(t *testing.T) {
	var (
		id     int64 = 1
		device       = "mobile"
	)

	InitAuth()

	token, err := Auth.Login(id, LoginConfig{
		Device:  device,
		Timeout: 5,
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("token:", token)

	time.Sleep(6 * time.Second)

	var status bool
	status, err = Auth.CheckLogin(id)
	t.Log("login status:", status, err)

	status, err = Auth.CheckToken(token)
	t.Log("login status:", status, err)
}

func TestAuth_Login_RenewToken(t *testing.T) {
	var (
		id     int64 = 1
		device       = "mobile"
	)

	InitAuth()

	token, err := Auth.Login(id, LoginConfig{
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

	var status bool
	status, err = Auth.CheckToken(token)
	t.Log("login status:", status, err)

	// 休眠 3 秒，检查 token 是否过期（没过期）
	time.Sleep(3 * time.Second)
	// 休眠 6 秒，检查 token 是否过期（已经过期）
	// time.Sleep(6 * time.Second)
	status, err = Auth.CheckLogin(id)
	t.Log("login status:", status, err)

	status, err = Auth.CheckToken(token)
	t.Log("login status:", status, err)
}
