package route

import (
	"github.com/audioo/bitcrook/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// Domain route
func Domain(app *fiber.App) {
	var h handlers.DomainHandler
	r := app.Group("/domain")
	r.Get("/", h.Index)
	r.Get("/:domain", h.Show)
}
