package utilities

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/julyskies/gohelpers"

	"go-kafka-publisher/constants"
)

func Response(options ResponseOptions) error {
	info := options.Info
	if info == "" {
		info = constants.RESPONSE_INFO.Ok
	}

	status := options.Status
	if status == 0 {
		status = fiber.StatusOK
	}

	responseObject := fiber.Map{
		"datetime": gohelpers.MakeTimestampSeconds(),
		"info":     info,
		"request": fmt.Sprintf(
			"%s [%s]",
			options.Context.OriginalURL(),
			options.Context.Method(),
		),
		"status": status,
	}

	if options.Data != nil {
		responseObject["data"] = options.Data
	}

	return options.Context.Status(status).JSON(responseObject)
}
