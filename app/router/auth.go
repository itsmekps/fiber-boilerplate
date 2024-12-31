package router

import (
	"fiber-boilerplate/app/dtos"
	"fiber-boilerplate/app/handlers"
	"fiber-boilerplate/app/messages"
	"fiber-boilerplate/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(router fiber.Router) {
	// Initialize a new instance of the validator
	validate := validator.New()
	// Create a sub-group of routes under the "/auth" path
	authGroup := router.Group("/auth")
	{
		// Validate login request body
		authGroup.Post("/login", middleware.ValidateRequestDTO(validate, &dtos.LoginRequest{}, messages.Login), handlers.Login)
	}
}
