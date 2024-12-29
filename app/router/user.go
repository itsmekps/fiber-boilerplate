// package router

// import (
// 	"fiber-boilerplate/app/dtos"
// 	"fiber-boilerplate/app/handlers/user"
// 	"fiber-boilerplate/app/middleware"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/gofiber/fiber/v2"
// )

// func UserRouter(router fiber.Router) {
// 	validate := validator.New()
// 	var getUserDetailsRequest dtos.GetUserDetailsRequest
// 	var loginRequest dtos.LoginRequest
// 	userGroup := router.Group("/users")
// 	{
// 		userGroup.Post("/login", middleware.ValidateRequest(validate, &loginRequest), user.Login)
// 		userGroup.Get("/details", user.GetUserDetails)
// 		userGroup.Get("/:id", middleware.ValidateRequest(validate, &getUserDetailsRequest), user.GetUser)
// 	}
// }

package router

import (
	"fiber-boilerplate/app/dtos"
	user "fiber-boilerplate/app/handlers"
	"fiber-boilerplate/app/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router) {
	validate := validator.New()
	// var getUserDetailsRequest dtos.GetUserByMongoIDRequest
	// var loginRequest dtos.LoginRequest
	userGroup := router.Group("/users")
	{
		// Validate login request body
		// userGroup.Post("/login", middleware.ValidateRequestDTO(validate, &dtos.LoginRequest{}), user.Login)

		// Validate details request without parameters
		userGroup.Get("/details", user.GetUserDetails)

		// Validate route with integer ID
		// userGroup.Get("/:id", middleware.ValidateRequest(validate, &getUserDetailsRequest, map[string]string{
		// 	"id": "objectid",
		// }), user.GetUser)

		userGroup.Get("/:id", middleware.ValidateRequestDTO(validate, &dtos.GetUserByMongoID{}), user.GetUser)

		// Validate route with MongoDB ObjectID
		// userGroup.Get("/mongo/:id", middleware.ValidateRequest(validate, nil, map[string]string{
		// 	"id": "objectid",
		// }), user.GetMongoUser)
	}
}
