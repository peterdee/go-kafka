package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"go-kafka-consumer/broker"
	"go-kafka-consumer/constants"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal("Could not load .env file!")
	}

	brokerError := broker.CreateConnection(
		os.Getenv(constants.ENV_NAMES.BrokerAddress),
	)
	if brokerError != nil {
		log.Fatal(brokerError)
	}

	broker.ReadMessages()
	broker.CloseConnection()
}
