package handlers

import (
	hits "github.com/audioo/bitcrook/pkg/noauth/vin"
	"github.com/gofiber/fiber/v2"
)

// VinHandler is ...
type VinHandler struct {
}

// Index ...
func (VinHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/vin.html")
}

// Show VIN information
func (VinHandler) Show(ctx *fiber.Ctx) error {
	vin := ctx.Params("vin")

	vininfo, err := hits.VinLookup(vin)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "failure"})
	}

	// return ctx.JSON(fiber.Map{"data": userinfo})
	return ctx.JSON(vininfo)
}
