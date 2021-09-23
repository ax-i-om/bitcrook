package handlers

import (
	"github.com/gofiber/fiber/v2"
	hits "github.com/maraudery/goseek/pkg/noauth/vin"
)

// VinHandler is ...
type VinHandler struct {
}

// Index ...
func (h VinHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/vin.html")
}

// Show VIN information
func (h VinHandler) Show(ctx *fiber.Ctx) error {
	vin := ctx.Params("vin")

	vininfo, err := hits.VinLookup(vin)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "failure"})
	}

	// return ctx.JSON(fiber.Map{"data": userinfo})
	return ctx.JSON(vininfo)
}
