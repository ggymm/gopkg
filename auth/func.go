package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/ggymm/gopkg/utils"
)

func notInit() bool {
	return auth == nil
}

// ------------------------------------------------ no ctx ----------------------------------------------------------------

func Login(id int64, config ...LoginConfig) (string, error) {
	if notInit() {
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
	if notInit() {
		return false, errors.New(ErrAuthNotInit)
	}
	return false, nil
}

func GetSession(id int64) (interface{}, error) {
	return nil, nil
}

func SaveSession(id int64, value interface{}) error {
	if notInit() {
		return errors.New(ErrAuthNotInit)
	}
	return auth.SetSessionData(id, value)
}

// ------------------------------------------------ fiber web ctx ----------------------------------------------------------------

func WebLogin(id int64, ctx *fiber.Ctx, config ...LoginConfig) (string, error) {
	if notInit() {
		return "", errors.New(ErrAuthNotInit)
	}

	// 获取配置
	var cfg = LoginConfig{
		Device:  "web",
		Timeout: auth.tokenTimeout,
	}
	if len(config) > 0 {
		cfg = config[0]
	}

	// 执行登陆
	token, err := auth.Login(id, cfg)
	if err != nil {
		return "", err
	}

	// 写入 cookie
	cookie := fasthttp.AcquireCookie()
	cookie.SetKey(auth.tokenName)
	cookie.SetValue(token)
	cookie.SetPath("/")
	// cookie.SetDomain(ctx.Hostname()) // 会自动设置
	if cfg.Timeout == -1 {
		// 设置永不过期
		cookie.SetMaxAge(utils.YearToSecond(1))
		// cookie.SetExpire(time.Now().AddDate(1, 0, 0))
	} else if cfg.Timeout > 0 {
		cookie.SetMaxAge(int(cfg.Timeout.Seconds()))
		// cookie.SetExpire(time.Now().Add(time.Duration(cfg.Timeout) * time.Second))
	}
	ctx.Response().Header.SetCookie(cookie)
	fasthttp.ReleaseCookie(cookie)

	// 写入 response header
	ctx.Response().Header.Set(auth.tokenName, token)
	return token, nil
}

func WebCheck(ctx *fiber.Ctx) (bool, error) {
	if notInit() {
		return false, errors.New(ErrAuthNotInit)
	}

	// 从请求体中获取 token
	token := ctx.Get(auth.tokenName)

	// 从请求头中获取 token
	if len(token) == 0 {
		token = ctx.GetReqHeaders()[auth.tokenName]
	}

	// 从 cookie 中获取 token
	if len(token) == 0 {
		token = ctx.Cookies(auth.tokenName)
	}

	return auth.CheckToken(token)
}

func WebGetSession(ctx *fiber.Ctx) (interface{}, error) {
	if notInit() {
		return nil, errors.New(ErrAuthNotInit)
	}

	// 从请求体中获取 token
	token := ctx.Get(auth.tokenName)

	// 从请求头中获取 token
	if len(token) == 0 {
		token = ctx.GetReqHeaders()[auth.tokenName]
	}

	// 从 cookie 中获取 token
	if len(token) == 0 {
		token = ctx.Cookies(auth.tokenName)
	}

	// 获取 session
	data, err := auth.GetSessionData(token)
	if err != nil {
		return nil, err
	}
	return data, nil
}
