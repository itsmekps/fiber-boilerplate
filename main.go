package main

import (
	"fiber-boilerplate/app/database"
	"fiber-boilerplate/app/logger"
	"fiber-boilerplate/app/middleware"
	"fiber-boilerplate/app/repository"
	"fiber-boilerplate/app/router"
	"fiber-boilerplate/app/service"
	"fiber-boilerplate/config"
	"fiber-boilerplate/internal/bootstrap"
)

func main() {
	// Initialize logger
	log := logger.NewLogger()

	v, err := config.InitConfig()
	if err != nil {

		log.Fatal(err)
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize repositories
	repos := repository.InitRepositories(db)

	// Initialize services
	service.InitServices(repos)

	app := bootstrap.InitWebServer()
	app.Use(middleware.LogMiddleware())
	// app.Use(func(c *fiber.Ctx) error {
	// 	log.Logger.WithFields(logrus.Fields{
	// 		"method": c.Method(),
	// 		"path":   c.Path(),
	// 	}).Info("Request received")
	// 	return c.Next()
	// })

	middleware.RegisterMiddleware(app)
	// router.RegisterAPIRoutes(app)
	// router.SetupAPI(app, userService)
	router.ApiRouter(app)
	log.Fatal(app.Listen(":" + v.GetString("Port")))
}
