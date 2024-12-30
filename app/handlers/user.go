package handlers

import (
	"fiber-boilerplate/app/dtos"
	"fiber-boilerplate/app/service/mongodb"

	"github.com/gofiber/fiber/v2"
)

// GetUser returns a user by ID
func GetUser(c *fiber.Ctx) error {
	// Retrieve the validated request object from Fiber's context
	// The validated request contains the parsed and validated ObjectID from the route parameter
	validatedRequest := c.Locals("validatedRequest").(*dtos.GetUserByMongoID)
	
	// Fetch the user from the database using the validated ObjectID
	// The service layer is responsible for interacting with MongoDB
	user, err := mongodb.UserServiceInstance.GetUser(validatedRequest.ID)
	if err != nil {
		// Return the error if the user retrieval fails (e.g., user not found or DB error)
		return err
	}

	// Create a sanitized response object excluding sensitive fields (e.g., password)
	response := dtos.UserResponse{
		ID:       user.ID.Hex(), // Convert the ObjectID to its string representation (Hex format)
		Username: user.FirstName, // User's username
		Email:    user.Email,    // User's email address
	}

	// Send the JSON response back to the client, wrapping the user data in a success object
	return c.JSON(fiber.Map{
		"success": true,                  // Indicate the operation was successful
		"data":    fiber.Map{"user": response}, // Include the sanitized user data
	})
}


func GetUserDetails(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "API is working!"})
}
