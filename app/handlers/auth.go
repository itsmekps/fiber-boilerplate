package handlers

import (
	"errors"
	"fiber-boilerplate/app/dtos"
	"fiber-boilerplate/app/service/mongodb"
	"fiber-boilerplate/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Login returns a access token by username and password
func Login(c* fiber.Ctx) error {
	// validatedRequest := c.Locals("validatedRequest").(*dtos.LoginRequest)

	var loginRequest dtos.LoginRequest

	email := loginRequest.Email
	password := loginRequest.Password
	
	fmt.Println("email: ", email)
	fmt.Println("password: ", password)

	user, err := mongodb.UserServiceInstance.GetUserByEmail(email)
	fmt.Println("user details: ", user)

	if err != nil || user == nil {
		fmt.Println("inside service error: ")

        return errors.New("user not found")
    }

	if !utils.CheckPasswordHash(password, user.Password) {
        return errors.New("invalid password")
    }

	// utils.GenerateToken(user.ID)

	return c.JSON(fiber.Map{"success": true, "data": fiber.Map{"user": user}})
}