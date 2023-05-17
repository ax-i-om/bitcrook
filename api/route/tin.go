package route

import (
	"github.com/ax-i-om/bitcrook/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// Tin route
func Tin(app *fiber.App) {
	var h handlers.TinHandler
	r := app.Group("/tin")
	r.Get("/", h.Index)
	r.Get("/:tin", h.Show)
}
