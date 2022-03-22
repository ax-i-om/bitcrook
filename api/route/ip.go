package route

import (
	"github.com/audioo/cycull/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// Ip route
func Ip(app *fiber.App) {
	var h handlers.IpHandler
	r := app.Group("/ip")
	r.Get("/", h.Index)
	r.Get("/:ip", h.Show)
}
