package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"go-kafka-consumer/broker"
	"go-kafka-consumer/constants"
)

func main() {
	envSource := os.Getenv(constants.ENV_NAMES.ENV_SOURCE)
	if envSource != "env" {
		envError := godotenv.Load()
		if envError != nil {
			log.Fatal("Could not load .env file!")
		}
	}

	broker.CreateReader(os.Getenv(constants.ENV_NAMES.BROKER_ADDRESS))

	// broker.ReadMessages()
	// broker.CloseConnection()
}
