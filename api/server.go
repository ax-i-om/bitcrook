package main

import (
	"log"

	"github.com/bitcrook/cybull/api/handlers"
	"github.com/bitcrook/cybull/api/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

// Starting the Fiber web app
func main() {
	// Create new app
	app := fiber.New(fiber.Config{
		// Pass view engine
		Views: html.New("./views", ".html"),
		// Pass global error handler
		ErrorHandler: handlers.Errors("./public/500.html"),
	})

	// Render index template with IP input value
	app.Get("/", handlers.Render())

	// Serve static assets
	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		CacheDuration: -1,
		MaxAge:        0, // this is redundant as 0 is the default
	})

	// Page routing
	route.Username(app)
	route.Ip(app)
	route.Vin(app)
	route.Tools(app)

	// Handle 404 errors
	app.Use(handlers.NotFound("./public/404.html"))

	// Set port to 6174
	port := "6174"

	// Start server on port 6174
	log.Fatal(app.Listen(":" + port))
}
