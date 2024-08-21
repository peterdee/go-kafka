package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"

	"go-kafka-consumer/constants"
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

func CreateConnection(address string) (*kafka.Conn, error) {
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

	Connection.SetReadDeadline(time.Now().Add(10 * time.Second))

	return Connection, connectionError
}

func ReadMessages() error {
	if checkError := checkConnection(); checkError != nil {
		return checkError
	}
	batch := Connection.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	buffer := make([]byte, 10e4) // 100KB max per message
	for i := 0; ; i += 1 {
		bytes, readError := batch.Read(buffer)
		if readError != nil {
			break
		}
		fmt.Println(i, string(buffer[:bytes]))
	}

	if closeBatchError := batch.Close(); closeBatchError != nil {
		return closeBatchError
	}
	return nil
}
