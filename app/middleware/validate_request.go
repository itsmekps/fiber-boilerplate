// middleware/validate_request.go
package middleware

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// ValidationError returns a validation error response
func ValidationError(c *fiber.Ctx, field string, error string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Validation failed",
		"details": []fiber.Map{
			{
				"field": field,
				"error": error,
			},
		},
	})
}

// ValidateRequest validates the request body and path parameters
func ValidateRequest(validate *validator.Validate, request interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Only parse the request body for non-GET requests
		if c.Method() != fiber.MethodGet {
			err := c.BodyParser(request)
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
			}
		}

		// Validate the request
		err := validate.Struct(request)
		if err != nil {
			var details []struct {
				Field string `json:"field"`
				Error string `json:"error"`
			}

			for _, err := range err.(validator.ValidationErrors) {
				details = append(details, struct {
					Field string `json:"field"`
					Error string `json:"error"`
				}{
					Field: err.Field(),
					Error: err.Tag(),
				})
			}

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   "Validation failed",
				"details": details,
			})
		}

		// Manually validate path parameters for GET requests
		if c.Method() == fiber.MethodGet {
			id := c.Params("id")
			if id == "" {
				return ValidationError(c, "id", "ID is required")
			}

			_, err := strconv.Atoi(id)
			if err != nil {
				return ValidationError(c, "id", "ID must be an integer")
			}
		}

		return c.Next()
	}
}
