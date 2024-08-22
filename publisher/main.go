package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"

	"go-kafka-publisher/broker"
	"go-kafka-publisher/constants"
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

	broker.WriteMessages(kafka.Message{Key: []byte("testkey"), Value: []byte("test")})
	broker.CloseConnection()
}
