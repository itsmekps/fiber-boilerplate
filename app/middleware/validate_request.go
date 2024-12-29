package middleware

import (
	"reflect"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ValidationError returns a structured JSON validation error response
func ValidationError(c *fiber.Ctx, field, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Validation failed",
		"details": []fiber.Map{
			{
				"field": field,
				"error": message,
			},
		},
	})
}

// ValidateRequestDTO dynamically validates request body and parameters based on the type struct
func ValidateRequestDTO(validate *validator.Validate, dto interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		value := reflect.ValueOf(dto).Elem()

		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			fieldValue := value.Field(i)

			// Check for param or JSON tags
			paramTag := field.Tag.Get("param")
			jsonTag := field.Tag.Get("json")

			if paramTag != "" {
				// Validate route parameters
				paramValue := c.Params(paramTag)
				if paramValue == "" {
					return ValidationError(c, paramTag, "Parameter is required")
				}

				// Validate parameter type
				switch field.Type.Kind() {
				case reflect.Int:
					intValue, err := strconv.Atoi(paramValue)
					if err != nil {
						return ValidationError(c, paramTag, "Parameter must be an integer")
					}
					fieldValue.SetInt(int64(intValue))
				case reflect.String:
					fieldValue.SetString(paramValue)
				case reflect.Struct:
					if field.Type == reflect.TypeOf(primitive.ObjectID{}) {
						objID, err := primitive.ObjectIDFromHex(paramValue)
						if err != nil {
							return ValidationError(c, paramTag, "Parameter must be a valid MongoDB ObjectID")
						}
						fieldValue.Set(reflect.ValueOf(objID))
					} else {
						return ValidationError(c, paramTag, "Unsupported parameter type")
					}
				
				default:
					return ValidationError(c, paramTag, "Unsupported parameter type")
				}
			}

			if jsonTag != "" && (c.Method() == fiber.MethodPost || c.Method() == fiber.MethodPatch || c.Method() == fiber.MethodPut) {
				// Parse request body for body-based fields
				if err := c.BodyParser(dto); err != nil {
					return ValidationError(c, "body", "Invalid request payload")
				}
			}
		}

		// Validate the DTO using struct tags
		if err := validate.Struct(dto); err != nil {
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

		// Store validated request for handler access
		// c.Locals("validatedRequest", dto)

		return c.Next()
	}
}
