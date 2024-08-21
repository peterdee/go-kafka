package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"

	"go-kafka-publisher/broker"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal("Could not load .env file!")
	}

	_, connectionError := broker.CreateConnection(os.Getenv("BROKER_ADDRESS"))
	if connectionError != nil {
		log.Fatal(connectionError)
	}

	broker.WriteMessages(kafka.Message{Key: []byte("testkey"), Value: []byte("test")})
	broker.CloseConnection()
}
