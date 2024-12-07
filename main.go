package main

import (
	"log"

	"fiber-boilerplate/app/database"
	"fiber-boilerplate/app/middleware"
	"fiber-boilerplate/app/router"
	"fiber-boilerplate/config"
	"fiber-boilerplate/internal/bootstrap"
)

func main() {
	v, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := bootstrap.InitWebServer()
	middleware.RegisterMiddleware(app)
	router.RegisterAPIRoutes(app)
	log.Fatal(app.Listen(":" + v.GetString("Port")))
}
