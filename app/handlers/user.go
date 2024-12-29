package handlers

import (
	"fiber-boilerplate/app/dtos"
	"fiber-boilerplate/app/service/mongodb"

	"github.com/gofiber/fiber/v2"
)

// GetUser returns a user by ID
func GetUser(c *fiber.Ctx) error {
	validatedRequest := c.Locals("validatedRequest").(*dtos.GetUserByMongoID)
	// id, err := c.ParamsInt("id")

	// id := c.Params("id")
	// if err != nil {
	// 	return err
	// }

	// Convert the ID to a MongoDB ObjectID
	// objID, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Invalid ID format",
	// 	})
	// }
	// fmt.Println("objID handler: ", objID)
	// Access the GetUser method
	// user, err := mysql.UserServiceInstance.GetUser(id)
	user, err := mongodb.UserServiceInstance.GetUser(validatedRequest.ID)
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
