package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/omniscient/handlers"
)

// Ip route
func Ip(app *fiber.App) {
	var h handlers.IpHandler
	r := app.Group("/ip")
	r.Get("/", h.Index)
	r.Get("/:ip", h.Show)
}
