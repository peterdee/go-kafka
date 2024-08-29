package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"

	eventsAPIs "go-kafka-publisher/apis/events"
	indexAPIs "go-kafka-publisher/apis/index"
	"go-kafka-publisher/broker"
	"go-kafka-publisher/constants"
	"go-kafka-publisher/utilities"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal("Could not load .env file!")
	}

	broker.CreateWriter(os.Getenv(constants.ENV_NAMES.BrokerAddress))

	app := fiber.New(fiber.Config{ErrorHandler: utilities.GlobalErrorHandler})

	app.Use(cors.New())
	app.Use(favicon.New(favicon.Config{
		File: "./static/favicon.ico",
	}))
	app.Use(logger.New())

	eventsAPIs.Initialize(app)
	indexAPIs.Initialize(app)

	port := utilities.GetEnv(constants.ENV_NAMES.Port, constants.DEFAULT_PORT)
	app.Listen(fmt.Sprintf(":%s", port))
}
