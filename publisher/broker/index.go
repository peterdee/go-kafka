package broker

import (
	"context"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func CreateWriter(address string) {
	Writer = &kafka.Writer{
		Addr:     kafka.TCP(address),
		Balancer: &kafka.LeastBytes{},
	}

	Writer.AllowAutoTopicCreation = true
}

func DestroyWriter() error {
	if closeError := Writer.Close(); closeError != nil {
		return closeError
	}
	return nil
}

func WriteMessages(ctx context.Context, messages ...kafka.Message) error {
	return Writer.WriteMessages(ctx, messages...)
}
