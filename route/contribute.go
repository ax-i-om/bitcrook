package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/omniscient/handlers"
)

// Contribute route
func Contribute(app *fiber.App) {
	var h handlers.ContributeHandler
	r := app.Group("/contribute")
	r.Get("/", h.Index)
}
