package auth

import (
	"github.com/pkg/errors"
)

func NotInit() bool {
	return Auth == nil
}

func Login(id int64, config ...LoginConfig) (string, error) {
	if NotInit() {
		return "", errors.New(ErrAuthNotInit)
	}
	var cfg = LoginConfig{
		Device:  "web",
		Timeout: Auth.tokenTimeout,
	}
	if len(config) > 0 {
		cfg = config[0]
	}
	return Auth.Login(id, cfg)
}

func Check(token string) (bool, error) {
	if NotInit() {
		return false, errors.New(ErrAuthNotInit)
	}
	return false, nil
}

func GetSession(id int64) (interface{}, error) {
	return nil, nil
}

func SaveSession(id int64, value interface{}) error {
	if NotInit() {
		return errors.New(ErrAuthNotInit)
	}
	return Auth.SetSessionData(id, value)
}
