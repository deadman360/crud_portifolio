package main

import (
	"context"
	"log"

	"github.com/deadman360/crud_portifolio/src/configuration/database/mongodb"
	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"github.com/deadman360/crud_portifolio/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting user aplication")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDbConnection(context.Background())
	if err != nil {
		log.Fatal("Error trying to Connect with database")
	}

	userController := InitDependecies(database)

	router := gin.Default()

	userGroup := router.Group("/user")

	routes.InitRoutes(userGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
