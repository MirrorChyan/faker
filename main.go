package main

import (
	"fmt"
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    Data   `json:"data"`
}

type Data struct {
	VersionName   string `json:"version_name"`
	VersionNumber int    `json:"version_number"`
	Channel       string `json:"channel"`
	OS            string `json:"os,omitempty"`
	Arch          string `json:"arch,omitempty"`
	ReleaseNote   string `json:"release_note"`
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		ProxyHeader: fiber.HeaderXForwardedFor,
	})

	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}?${queryParams} | ${error}\n",
	}))

	app.Get("/resources/:rid/latest", func(c *fiber.Ctx) error {
		currentVersion := c.Query("current_version")
		channel := c.Query("channel", "stable")
		os := c.Query("os")
		arch := c.Query("arch")

		return c.JSON(Response{
			Code:    0,
			Message: "server is under maintenance, so i am faker",
			Data: Data{
				VersionName: currentVersion,
				Channel:     channel,
				OS:          os,
				Arch:        arch,
			},
		})
	})

	addr := fmt.Sprintf(":%d", 8000)
	if err := app.Listen(addr); err != nil {
		log.Fatal("failed to start server")
	}
}
