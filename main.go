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
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db, err := database.InitDB()
	// defer db.MySQL.Close()

	// db_mysql, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// db_mongo, err := database.InitDB()
	// // Initialize MongoDB database
	// mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://kpskispl:shuklaCLG%40123!@clg.bg5vb.mongodb.net/?retryWrites=true&w=majority&appName=CLG"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db_mongo.Close()
	// defer mongoClient.Disconnect(context.TODO())
	// mongoDB := mongoClient.Database("your_database_name")

	// connections := database.DBConnections{
	// 	MySQL: db.db_mysql,
	// }

	// Initialize repositories
	repos := repository.InitRepositories(db)

	fmt.Printf("%v\n", repos)
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
	//

	middleware.RegisterMiddleware(app)
	// router.RegisterAPIRoutes(app)
	// router.SetupAPI(app, userService)
	router.ApiRouter(app)
	startServer(app, db)
	// log.Fatal(app.Listen(":" + v.GetString("Port")))
	// defer db.MySQL.Close()
}

func startServer(app *fiber.App, db database.DBConnections) {
	// Initialize logger
	log := logger.NewLogger()

	v, err := config.InitConfig()
	if err != nil {

		log.Fatal(err)
	}
	log.Fatal(app.Listen(":" + v.GetString("Port")))
	// defer db.MySQL.Close()
}
