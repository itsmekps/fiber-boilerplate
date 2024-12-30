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
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	repos := repository.InitRepositories(db)

	// Initialize services
	service.InitServices(repos)

	app := bootstrap.InitWebServer()
	app.Use(middleware.LogMiddleware())
	app.Use(cors.New(config.CORSConfig()))
	middleware.RegisterMiddleware(app)

	router.ApiRouter(app)
	startServer(app)
}

func startServer(app *fiber.App) {
	// Initialize logger
	log := logger.NewLogger()

	v, err := config.InitConfig()
	if err != nil {

		log.Fatal(err)
	}
	log.Fatal(app.Listen(":" + v.GetString("Port")))
	// defer db.MySQL.Close()
}
