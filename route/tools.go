package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/goseek/handlers"
)

// Tools route
func Tools(app *fiber.App) {
	var h handlers.ToolsHandler
	r := app.Group("/tools")
	r.Get("/", h.Index)
}
