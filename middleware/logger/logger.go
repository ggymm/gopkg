package logger

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/ggymm/gopkg/constant"
	"github.com/ggymm/gopkg/log"
)

func needLog(contentType []byte) bool {
	var typeList = []string{
		"application/json",
		"application/xml",
		"text/xml",
		"text/plain",
		"application/x-www-form-urlencoded",
	}

	// 判断请求类型 是否在 contentTypes 中
	for _, t := range typeList {
		if string(contentType) == t {
			return true
		}
	}
	return false
}

func NewLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// next middleware
		if err := c.Next(); err != nil {
			return err
		}

		// 计算请求耗时
		elapsed := time.Since(start)
		costTime := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)

		var params map[string]string
		var reqBody []byte

		// 判断请求类型 是否在 contentTypes 中

		if needLog(c.Request().Header.ContentType()) {
			params = c.Queries()
			reqBody = c.Request().Body()
		}

		var respBody []byte
		if needLog(c.Response().Header.ContentType()) {
			respBody = c.Response().Body()
		}

		// 记录请求日志
		log.Info().
			// 请求基本参数
			Str("ip", c.IP()).
			Str("path", c.Path()).
			Str("method", c.Method()).
			Str("costTime", costTime).

			// 请求参数
			Str("params", fmt.Sprintf("%v", params)).
			Str("reqBody", fmt.Sprintf("%s", reqBody)).

			// 响应参数
			Str("respBody", fmt.Sprintf("%s", respBody)).
			Msg(constant.APITrace)

		return nil
	}
}
