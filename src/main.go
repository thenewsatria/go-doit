package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	testing := os.Getenv("DEV_TEST_ENV")

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"testing": testing,
			"mantap":  "bingit",
		})
	})
	log.Fatal(app.Listen(":3000"))
}
