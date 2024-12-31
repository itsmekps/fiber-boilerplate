package router

import (
	player "fiber-boilerplate/app/handlers"
	"fiber-boilerplate/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PlayRouter(router fiber.Router) {

	// Initialize a new instance of the validator
	// validate := validator.New()

	// Create a sub-group of routes under the "/players" path, protected by the AuthMiddleware for authentication
	playerGroup := router.Group("/players", middleware.AuthMiddleware())

	{
		// get list of all players based on filters/query params
		playerGroup.Get("/", player.GetPlayersList)

		// get details of a player by playerID
		playerGroup.Get("/:id", player.GetPlayerDetails)
	}
}
