package router

import (
	"fiber-boilerplate/app/dtos"
	user "fiber-boilerplate/app/handlers"
	"fiber-boilerplate/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router) {
	validate := validator.New()

	// userGroup := router.Group("/users")
	userGroup := router.Group("/users", middleware.AuthMiddleware())
	
	{
		// Validate details request without parameters
		userGroup.Get("/details", user.GetUserDetails)

		userGroup.Get("/:id", middleware.ValidateRequestDTO(validate, &dtos.GetUserByMongoID{}), user.GetUser)
		// userGroup.Get("/", user.GetUser)
	}
}
