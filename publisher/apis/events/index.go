package events

import "github.com/gofiber/fiber/v3"

func Initialize(app *fiber.App) {
	group := app.Group("/api/events")
	group.Post("/", eventsController)
}
