package main

import (
	"fmt"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/ggymm/gopkg/auth"
	"github.com/ggymm/gopkg/log"
)

func main() {
	log.Init("")

	if err := auth.Init(auth.Config{
		Store: auth.Local,
		LocalConfig: auth.LocalConfig{
			Dir: "./temp/auth/storage",
		},

		Concurrent:    true,
		ShareToken:    false,
		MaxLoginCount: math.MaxInt,

		TokenName:      "ninelock-token",
		TokenTimeout:   time.Duration(30) * time.Minute,
		AutoRenewToken: true,
	}); err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		Views: html.New("./temp/auth/", ".html"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		status, err := auth.WebCheck(c)
		if err != nil {
			return err
		}
		if status {
			return c.Redirect("/index")
		} else {
			return c.Redirect("/login")
		}
	})

	app.Get("/index", func(c *fiber.Ctx) error {
		status, err := auth.WebCheck(c)
		if err != nil {
			return err
		}
		if status {
			return c.Render("index", fiber.Map{
				"User":  "Admin",
				"Token": "1234567890",
			})
		} else {
			return c.Redirect("/login")
		}
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		// 执行登陆
		token, err := auth.WebLogin(1, c)
		if err != nil {
			return err
		}
		fmt.Println("login success, token:", token)
		return c.Redirect("/index")
	})

	app.Post("/login/timeout", func(c *fiber.Ctx) error {
		// 执行登陆
		token, err := auth.WebLogin(1, c, auth.LoginConfig{
			Device:  "web",
			Timeout: time.Duration(1) * time.Minute,
		})
		if err != nil {
			return err
		}
		fmt.Println("login success, token:", token)
		return c.Redirect("/index")
	})

	panic(app.Listen(":8888"))
}
