package main

import (
	"context"
	"fiber-boilerplate/app/database"
	"fiber-boilerplate/app/logger"
	"fiber-boilerplate/app/middleware"
	"fiber-boilerplate/app/repository"
	"fiber-boilerplate/app/router"
	"fiber-boilerplate/app/service"
	"fiber-boilerplate/config"
	"fiber-boilerplate/internal/bootstrap"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Initialize logger
	log := logger.NewLogger()

	v, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize MySQL database
	mysqlDB, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlDB.Close()

	// Initialize MongoDB database
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://kpskispl:shuklaCLG%40123!@clg.bg5vb.mongodb.net/?retryWrites=true&w=majority&appName=CLG"))
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(context.TODO())

	mongoDB := mongoClient.Database("your_database_name")

	// Initialize repositories
	repos, err := repository.InitRepositories("both", mysqlDB, mongoDB)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize services
	service.InitServices(repos.MySQL)
	// service.InitServices(repos.MongoDB)

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
