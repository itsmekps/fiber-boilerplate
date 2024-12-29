package router

import (
	"fiber-boilerplate/app/dtos"
	"fiber-boilerplate/app/handlers/user"
	"fiber-boilerplate/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router) {
	validate := validator.New()
	var getUserDetailsRequest dtos.GetUserDetailsRequest
	var loginRequest dtos.LoginRequest
	userGroup := router.Group("/users")
	{
		userGroup.Post("/login", middleware.ValidateRequest(validate, &loginRequest), user.Login)
		userGroup.Get("/details", user.GetUserDetails)
		userGroup.Get("/:id", middleware.ValidateRequest(validate, &getUserDetailsRequest), user.GetUser)
	}
}
