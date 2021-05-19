package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/omniscient/api/handlers"
)

// Username route
func Username(app *fiber.App) {
	var h handlers.UsernameHandler
	r := app.Group("/username")
	r.Get("/", h.Index)
	r.Get("/:username", h.Show)
}
