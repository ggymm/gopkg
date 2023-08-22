package auth

import (
	"github.com/pkg/errors"
	"time"
)

func NotInit() bool {
	return auth == nil
}

func GetTokenName() string {
	if NotInit() {
		return ""
	}
	return auth.tokenName
}

func GetDefaultTimeout() time.Duration {
	if NotInit() {
		return 0
	}
	return auth.tokenTimeout
}

func Login(id int64, config ...LoginConfig) (string, error) {
	if NotInit() {
		return "", errors.New(ErrAuthNotInit)
	}
	var cfg = LoginConfig{
		Device:  "web",
		Timeout: auth.tokenTimeout,
	}
	if len(config) > 0 {
		cfg = config[0]
	}
	return auth.Login(id, cfg)
}

func Check(token string) (bool, error) {
	if NotInit() {
		return false, errors.New(ErrAuthNotInit)
	}
	return auth.CheckToken(token)
}

func GetSession(id int64) (interface{}, error) {
	return nil, nil
}

func SaveSession(id int64, value interface{}) error {
	if NotInit() {
		return errors.New(ErrAuthNotInit)
	}
	return auth.SetSessionData(id, value)
}

func GetSessionData(token string) (interface{}, error) {
	if NotInit() {
		return nil, errors.New(ErrAuthNotInit)
	}
	return auth.GetSessionData(token)
}
