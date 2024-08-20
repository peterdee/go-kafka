package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"

	"go-kafka-publisher/constants"
)

func main() {
	connection, connectionError := kafka.DialLeader(
		context.Background(),
		"tcp",
		"localhost:9092",
		constants.DEFAULT_TOPIC_NAME,
		0,
	)
	if connectionError != nil {
		log.Fatal("failed to dial leader:", connectionError)
	}

	connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, writeError := connection.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if writeError != nil {
		log.Fatal("failed to write messages:", writeError)
	}

	if closeError := connection.Close(); closeError != nil {
		log.Fatal("failed to close writer:", closeError)
	}
}
