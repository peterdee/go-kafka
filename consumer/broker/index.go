package broker

import (
	"context"
	"fmt"
	"go-kafka-consumer/constants"

	"github.com/segmentio/kafka-go"
)

var Reader *kafka.Reader

func DestroyReader() error {
	if closeError := Reader.Close(); closeError != nil {
		return closeError
	}
	return nil
}

func CreateReader(address string) {
	Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{address},
		Topic:     constants.DEFAULT_TOPIC_NAME,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	_ = Reader.Close()
	m, er := Reader.ReadMessage(context.Background())
	e := Reader.Close()
	// s := Reader.Stats()
	fmt.Println(m.Value, er, e)
}

// func ReadMessages() error {
// 	if checkError := checkConnection(); checkError != nil {
// 		return checkError
// 	}
// 	batch := Connection.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

// 	buffer := make([]byte, 10e4) // 100KB max per message
// 	for i := 0; ; i += 1 {
// 		bytes, readError := batch.Read(buffer)
// 		if readError != nil {
// 			break
// 		}
// 		fmt.Println(i, string(buffer[:bytes]))
// 	}

// 	if closeBatchError := batch.Close(); closeBatchError != nil {
// 		return closeBatchError
// 	}
// 	return nil
// }
