package main

import (
	"fmt"
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
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
	})

	app.Get("/resources/:rid/latest", func(c *fiber.Ctx) error {
		current_version := c.Query("current_version")
		channel := c.Query("channel", "stable")
		os := c.Query("os")
		arch := c.Query("arch")

		return c.Status(fiber.StatusOK).JSON(Response{
			Code:    0,
			Message: "server is under maintenance",
			Data: Data{
				VersionName:   current_version,
				VersionNumber: 0,
				Channel:       channel,
				OS:            os,
				Arch:          arch,
				ReleaseNote:   "",
			},
		})
	})

	addr := fmt.Sprintf(":%d", 8000)
	if err := app.Listen(addr); err != nil {
		log.Fatal("failed to start server")
	}
}
