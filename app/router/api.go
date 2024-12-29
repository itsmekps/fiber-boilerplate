package router

import (
	"github.com/gofiber/fiber/v2"
)

func ApiRouter(app *fiber.App) {

	apiGroup := app.Group("/api")
	// Register user routes
	AuthRouter(apiGroup)
	UserRouter(apiGroup)

	// Register other routes here...
}
