package main

import (
	"fmt"
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Get("/resources/:rid/latest", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": 0,
			"msg":  "server is under maintenance",
			"data": fiber.Map{
				"version_name":   "",
				"version_number": 0,
			},
		})
	})

	addr := fmt.Sprintf(":%d", 8000)
	if err := app.Listen(addr); err != nil {
		log.Fatal("failed to start server")
	}
}
