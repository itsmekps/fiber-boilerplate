package middleware

import (
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Set Credentials
		expectedUsername := "admin"
		expectedPassword := "password"

		// Get Authorization Header
		authHeader := c.Get("Authorization")

		// check if auth header exists
		if authHeader == "" {
			return c.Status(401).SendString("Unauthorized")
		}

		// Verify credentials
		if authHeader != "Basic "+BasicAuth(expectedUsername, expectedPassword) {
			return c.Status(401).SendString("Unauthorized")
		}

		// Continue to next handler
		return c.Next()
	}
}

func BasicAuth(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}
