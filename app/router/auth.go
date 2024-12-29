package router

import (
	"fiber-boilerplate/app/dtos"
	"fiber-boilerplate/app/handlers"
	"fiber-boilerplate/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(router fiber.Router) {
	validate := validator.New()
	// var getUserDetailsRequest dtos.GetUserByMongoIDRequest
	// var loginRequest dtos.LoginRequest
	authGroup := router.Group("/auth")
	{
		// Validate login request body
		authGroup.Post("/login", middleware.ValidateRequestDTO(validate, &dtos.LoginRequest{}), handlers.Login)

	}
}
