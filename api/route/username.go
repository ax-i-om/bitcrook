package route

import (
	"github.com/audioo/bitcrook/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// Username route
func Username(app *fiber.App) {
	var h handlers.UsernameHandler
	r := app.Group("/username")
	r.Get("/", h.Index)
	r.Get("/:username", h.Show)
}
