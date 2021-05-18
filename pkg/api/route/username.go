package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/omniscient/pkg/api/handler"
)

// Username route
func Username(app *fiber.App) {
	var h handler.UsernameHandler
	r := app.Group("/username")
	r.Get("/", h.Index)
	r.Get("/:username", h.Show)
}
