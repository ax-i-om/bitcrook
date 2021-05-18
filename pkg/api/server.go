package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/omniscient/pkg/api/route"
)

// Launch starts the API
func Launch() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Omniscient API",
		})
	})
	route.Username(app)
	route.Discord(app)
	app.Listen(":1337")
}
