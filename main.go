package main

import (
	"log"

	"fiber-boilerplate/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	v, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Fiber app
	app := fiber.New()

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, World!")
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	// Start the server
	log.Fatal(app.Listen(":" + v.GetString("PORT")))
}
