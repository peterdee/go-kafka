package broker

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"

	"go-kafka-publisher/constants"
)

var Connection *kafka.Conn

func checkConnection() error {
	_, checkError := Connection.Brokers()
	return checkError
}

func CloseConnection() error {
	if checkError := checkConnection(); checkError != nil {
		return checkError
	}
	if closeError := Connection.Close(); closeError != nil {
		return closeError
	}
	log.Print("Connection to broker closed")
	return nil
}

func CreateConnection(address string) error {
	var connectionError error
	for i := 1; i <= 5; i += 1 {
		Connection, connectionError = kafka.DialLeader(
			context.Background(),
			"tcp",
			address,
			constants.DEFAULT_TOPIC_NAME,
			0,
		)
		if connectionError != nil {
			if i < 5 {
				log.Printf("Could not connect to broker, reconnecting in %d seconds...", i)
				time.Sleep(time.Duration(i) * time.Second)
			}
		} else {
			log.Print("Connected to broker")
			break
		}
	}

	Connection.SetWriteDeadline(time.Now().Add(10 * time.Second))

	return connectionError
}

func WriteMessages(messages ...kafka.Message) (int, error) {
	if err := checkConnection(); err != nil {
		return 0, err
	}
	return Connection.WriteMessages(messages...)
}
