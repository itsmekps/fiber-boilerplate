package router

import (
	"fiber-boilerplate/app/dtos"
	user "fiber-boilerplate/app/handlers"
	"fiber-boilerplate/app/messages"
	"fiber-boilerplate/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router) {
	validate := validator.New()

	// Create a sub-group of routes under the "/users" path
	userGroup := router.Group("/users")

	{
		// Validate details request without parameters
		userGroup.Get("/details", user.GetUserDetails)

		userGroup.Get("/:id", middleware.ValidateRequestDTO(validate, &dtos.GetUserByMongoID{}, messages.GetUser), user.GetUser)
		// userGroup.Get("/", user.GetUser)
	}
}
