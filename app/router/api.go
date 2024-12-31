package router

import (
	"github.com/gofiber/fiber/v2"
)

func ApiRouter(app *fiber.App) {
	// Create a main API group under the "/api" path to organize all API endpoints
	apiGroup := app.Group("/api")
	// Register auth routes
	AuthRouter(apiGroup)
	// Register user routes
	UserRouter(apiGroup)
	// Register player routes
	PlayRouter(apiGroup)
}
