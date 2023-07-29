package recover

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"

	"github.com/ggymm/gopkg/constant"
	"github.com/ggymm/gopkg/log"
)

func NewRecover() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {

		defer func() {
			if r := recover(); r != nil {

				log.Error().Err(errors.WithStack(r.(error))).Msg(constant.ServerPanic)

				err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"msg":     constant.ServerPanic,
					"success": false,
				})
			}
		}()

		return c.Next()
	}
}
