/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/audioo/bitcrook/api/handlers"
	"github.com/audioo/bitcrook/api/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/spf13/cobra"
)

// runServerCmd represents the runServer command
var runServerCmd = &cobra.Command{
	Use:   "runServer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create new app
		app := fiber.New(fiber.Config{
			// Pass view engine
			Views: html.New("./api/views", ".html"),
			// Pass global error handler
			ErrorHandler: handlers.Errors("./api/public/500.html"),
		})

		// Render index template with IP input value
		app.Get("/", handlers.Render())

		// Serve static assets
		app.Static("/", "./api/public", fiber.Static{
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
		app.Use(handlers.NotFound("./api/public/404.html"))

		// Set port to 6174
		port := "6174"

		// Start server on port 6174
		log.Fatal(app.Listen(":" + port))
	},
}

func init() {
	rootCmd.AddCommand(runServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
