package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/omniscient/handlers"
)

// Discord route
func Discord(app *fiber.App) {
	var h handlers.DiscordHandler
	r := app.Group("/discord")
	r.Get("/", h.Index)
	r.Get("/:token", h.Show)
}
