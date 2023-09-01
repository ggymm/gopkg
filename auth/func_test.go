package auth

import (
	"math"
	"testing"
)

func TestLogin(t *testing.T) {
	err := Init(Config{
		LogPath: "D:/temp/auth/auth.log",

		Store: Local,
		LocalConfig: LocalConfig{
			Dir: "D:/temp/auth/storage",
		},

		Concurrent:    true,
		ShareToken:    false,
		MaxLoginCount: math.MaxInt,

		TokenName:      "ninelock-token",
		AutoRenewToken: true,
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("auth token name:", GetTokenName())
	t.Log("auth token timeout:", GetDefaultTimeout())

	// 登录
	token, err := Login(1, LoginConfig{
		Device:  "web",
		Timeout: 5,
	})
	if err != nil {
		t.Errorf("login error: %s", err)
		return
	}
	t.Log("login success token:", token)

	// 检查登录状态
	status, err := Check(token)
	if err != nil {
		t.Errorf("check login error: %s", err)
		return
	}
	if !status {
		t.Error("check login error: status is false")
		return
	}
	t.Log("check login success")

	// 获取 session
	session, err := GetSession(token)
	if err != nil {
		t.Errorf("get session error: %s", err)
		return
	}
	t.Log("get session success:", session)

	// 保存 session
	err = SaveSession(1, session)
	if err != nil {
		t.Errorf("save session error: %s", err)
		return
	}

	// 重新获取 session
	session, err = GetSession(token)
	if err != nil {
		t.Errorf("get session error: %s", err)
		return
	}
	t.Log("get session success:", session)

	// 退出登录
	err = Logout(1)
	if err != nil {
		t.Errorf("logout error: %s", err)
		return
	}
}
