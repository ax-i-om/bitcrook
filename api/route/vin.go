package route

import (
	"github.com/audioo/cycull/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// Vin route
func Vin(app *fiber.App) {
	var h handlers.VinHandler
	r := app.Group("/vin")
	r.Get("/", h.Index)
	r.Get("/:vin", h.Show)
}
