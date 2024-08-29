package index

import (
	"github.com/gofiber/fiber/v3"

	"go-kafka-publisher/utilities"
)

func indexController(context fiber.Ctx) error {
	return utilities.Response(utilities.ResponseOptions{Context: context})
}
