package route

import (
	"github.com/audioo/bitcrook/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// Discord route
func Discord(app *fiber.App) {
	var h handlers.DiscordHandler
	r := app.Group("/discord")
	r.Get("/", h.Index)
	r.Get("/:discord", h.Show)
}
