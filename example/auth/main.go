package main

import (
	"fmt"
	"math"
	"runtime/debug"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"github.com/ggymm/gopkg/auth"
	"github.com/ggymm/gopkg/log"
)

type SessionData struct {
	UserId   int64
	UserName string
}

func main() {
	log.Init("")

	if err := auth.Init(auth.Config{
		LogPath: "./temp/auth/log",

		Store: auth.Local,
		LocalConfig: auth.LocalConfig{
			Dir: "./temp/auth/storage",
		},

		Concurrent:    true,
		ShareToken:    true,
		MaxLoginCount: math.MaxInt,

		TokenName:      "ninelock-token",
		TokenTimeout:   time.Duration(30) * time.Minute,
		AutoRenewToken: true,
	}); err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		Views: html.New("./example/auth/", ".html"),
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
			fmt.Println(string(debug.Stack()))
			fmt.Println("index error:", err)
			return err
		}
		if status {
			return c.Render("index", nil)
		} else {
			return c.Redirect("/login")
		}
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil)
	})

	app.Get("/session", func(c *fiber.Ctx) error {
		data, err := auth.WebGetSession(c)
		if err != nil {
			return err
		}
		fmt.Println(data)
		return c.Render("session", data)
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

	app.Post("/login/web", func(c *fiber.Ctx) error {
		// 执行登陆
		token, err := auth.WebLogin(1, c, auth.LoginConfig{
			Device:  "web",
			Timeout: -1,
		})
		if err != nil {
			fmt.Println(string(debug.Stack()))
			fmt.Println("login web error:", err)
			return err
		}
		fmt.Println("login web success, token:", token)
		return c.Redirect("/index")
	})

	app.Post("/login/mobile", func(c *fiber.Ctx) error {
		// 执行登陆
		token, err := auth.WebLogin(1, c, auth.LoginConfig{
			Device:  "mobile",
			Timeout: -1,
		})
		if err != nil {
			fmt.Println(string(debug.Stack()))
			fmt.Println("login mobile error:", err)
			return err
		}
		fmt.Println("login mobile success, token:", token)
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
		fmt.Println("login web timeout success, token:", token)
		return c.Redirect("/index")
	})

	app.Post("/login/session", func(c *fiber.Ctx) error {
		// 执行登陆
		token, err := auth.WebLogin(1, c, auth.LoginConfig{
			Device:  "web",
			Timeout: time.Duration(1) * time.Minute,
		})
		if err != nil {
			return err
		}
		fmt.Println("login web session success, token:", token)

		// 保存 session data
		err = auth.SaveSession(1, SessionData{
			UserId:   1,
			UserName: "Admin",
		})
		if err != nil {
			return err
		}
		return c.Redirect("/index")
	})

	panic(app.Listen(":8888"))
}
