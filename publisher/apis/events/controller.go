package events

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/segmentio/kafka-go"

	"go-kafka-publisher/broker"
	"go-kafka-publisher/constants"
	"go-kafka-publisher/utilities"
)

func eventsController(context fiber.Ctx) error {
	event := new(clientEvent)
	if bindError := context.Bind().Body(event); bindError != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			constants.RESPONSE_INFO.BadRequest,
		)
	}
	if event.EventTarget == "" || event.EventType == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			constants.RESPONSE_INFO.MissingData,
		)
	}

	bytes, jsonError := json.Marshal(event)
	if jsonError != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			constants.RESPONSE_INFO.InternalServerError,
		)
	}
	writeError := broker.WriteMessages(
		context.Context(),
		kafka.Message{
			Topic: constants.DEFAULT_TOPIC_NAME,
			Value: bytes,
		},
	)
	if writeError != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			constants.RESPONSE_INFO.InternalServerError,
		)
	}

	return utilities.Response(utilities.ResponseOptions{Context: context})
}
