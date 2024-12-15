// middleware/validate_request.go
package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Details []struct {
		Field string `json:"field"`
		Error string `json:"error"`
	} `json:"details"`
}

func ValidateRequest(validate *validator.Validate, request interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.BodyParser(request)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request payload")
		}

		err = validate.Struct(request)
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

			return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
				Error:   "Validation failed",
				Details: details,
			})
		}

		return c.Next()
	}
}
