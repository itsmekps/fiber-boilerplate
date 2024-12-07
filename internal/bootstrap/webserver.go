package bootstrap

import (
	"github.com/gofiber/fiber/v2"
)

func InitWebServer() *fiber.App {
	app := fiber.New()
	return app
}