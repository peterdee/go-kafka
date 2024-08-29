package index

import "github.com/gofiber/fiber/v3"

func Initialize(app *fiber.App) {
	group := app.Group("/")
	group.Get("/", indexController)
	group.Get("/api", indexController)
}
