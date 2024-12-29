package user

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
	validatedRequest := c.Locals("validatedRequest").(*dtos.LoginRequest)

	// Access email and password
	email := validatedRequest.Email
	password := validatedRequest.Password
	fmt.Println("email: ", email)
	user, err := mongodb.UserServiceInstance.GetUserByEmail(email)
	fmt.Println("user details: ", user)

	if err != nil || user == nil {
		fmt.Println("inside service error: ")

        return errors.New("user not found")
    }

	if !utils.CheckPasswordHash(password, user.Password) {
        return errors.New("invalid password")
    }

	return user
}

// GetUser returns a user by ID
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	// Access the GetUser method
	// user, err := mysql.UserServiceInstance.GetUser(id)
	user, err := mongodb.UserServiceInstance.GetUser(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "data": fiber.Map{"user": user}})
}

func GetUserDetails(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "API is working!"})
}

// GetUser Swagger documentation
// @Summary Get a user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
