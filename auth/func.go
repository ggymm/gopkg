package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func notInit() bool {
	return auth == nil
}

func Login(id int64, c LoginConfig) (string, error) {
	if notInit() {
		return "", errors.New(ErrAuthNotInit)
	}
	return auth.Login(id, c)
}

func LoginWeb(id int64, ctx *fiber.Ctx) error {
	if notInit() {
		return errors.New(ErrAuthNotInit)
	}
	return nil
}
