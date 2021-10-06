package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/goseek/handlers"
)

// Tools route
func Tools(app *fiber.App) {
	var h handlers.IpHandler
	r := app.Group("/tools")
	r.Get("/", h.Index)
	r.Get("/:tools", h.Show)
}
