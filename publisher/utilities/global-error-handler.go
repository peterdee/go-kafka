package utilities

import (
	"errors"

	"github.com/gofiber/fiber/v3"

	"go-kafka-publisher/constants"
)

func GlobalErrorHandler(context fiber.Ctx, err error) error {
	info := constants.RESPONSE_INFO.InternalServerError
	status := fiber.StatusInternalServerError

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		info = fiberError.Message
		if info == "Internal Server Error" {
			info = constants.RESPONSE_INFO.InternalServerError
		}
		status = fiberError.Code
	}

	return Response(ResponseOptions{
		Context: context,
		Info:    info,
		Status:  status,
	})
}
