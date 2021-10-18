package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/goseek/handlers"
)

// Usage route
func Usage(app *fiber.App) {
	var h handlers.UsageHandler
	r := app.Group("/usage")
	r.Get("/", h.Index)
}
