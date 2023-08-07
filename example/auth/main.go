package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ggymm/gopkg/auth"
	"github.com/ggymm/gopkg/log"
)

func main() {
	app := fiber.New()

	log.Init("")

	err := auth.Init(auth.Config{
		Store: auth.RoseDB,
		RoseDBConfig: auth.RoseDBConfig{
			DirPath: "./temp/auth/storage",
		},
	})
	if err != nil {
		panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/login.html")
	})
	app.Static("/", "./temp/auth/")

	v1 := app.Group("api/v1")
	{
		authApi := v1.Group("auth")
		{
			authApi.Post("login", func(c *fiber.Ctx) error {
				account := c.FormValue("account")
				password := c.FormValue("password")

				log.Info().Str("account", account).Str("password", password).Msg("login")

				if account == "admin" {

				}

				return c.SendString("login success")
			})
			authApi.Get("check", func(c *fiber.Ctx) error {

				return c.SendString("trace file")
			})
			authApi.Get("logout", func(c *fiber.Ctx) error {

				return c.SendString("trace file")
			})
			authApi.Get("kickout", func(c *fiber.Ctx) error {

				return c.SendString("trace file")
			})

			authApi.Get("userinfo", func(c *fiber.Ctx) error {

				return c.SendString("trace file")
			})
		}
	}

	panic(app.Listen(":8888"))

}
