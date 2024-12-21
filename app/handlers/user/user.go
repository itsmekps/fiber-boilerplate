package user

import (
	"fiber-boilerplate/app/service/mongodb"

	"github.com/gofiber/fiber/v2"
)

// GetUser returns a user by ID
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	// Access the GetUser method
	// user, err := mysql.UserServiceInstance.GetUser(id)
	user, err := mongodb.UserServiceInstancem.GetUser(id)
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
