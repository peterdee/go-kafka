package main

import (
	"go-kafka-consumer/broker"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal("Could not load .env file!")
	}

	_, connectionError := broker.CreateConnection(os.Getenv("BROKER_ADDRESS"))
	if connectionError != nil {
		log.Fatal(connectionError)
	}

	broker.ReadMessages()
	broker.CloseConnection()
}
