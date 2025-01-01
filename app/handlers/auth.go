package handlers

import (
	"fiber-boilerplate/app/dtos"
	appErrors "fiber-boilerplate/app/errors"
	"fiber-boilerplate/app/service/mongodb"
	"fiber-boilerplate/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Login returns a access token by username and password
func Login(c *fiber.Ctx) error {
	// Retrieve validated email and password from request context
	validatedRequest := c.Locals("validatedRequest").(*dtos.LoginRequest)

	// Extract email and password from the validated request
	email := validatedRequest.Email
	password := validatedRequest.Password

	// Fetch user by email from the database
	user, err := mongodb.UserServiceInstance.GetUserByEmail(email)
	fmt.Println("user details: ", user)

	// Return error if user is not found or retrieval fails
	if err != nil || user == nil {
		return appErrors.USER_NOT_FOUND.Respond(c)
	}

	// Validate the provided password against the stored hash
	if !utils.CheckPasswordHash(password, user.Password) {
		return appErrors.INVALID_PASSWORD.Respond(c)
	}

	// Generate a JWT token for the authenticated user
	jwt, err := utils.GenerateToken(user)
	if err != nil {
		return appErrors.INTERNAL_SERVER_ERROR.Respond(c)
	}

	// Respond with the generated access token
	return c.JSON(fiber.Map{"access_token": jwt,
		"token_type": "Bearer"})
}
