package handlers

import (
	hits "github.com/audioo/bitcrook/pkg/noauth/tin"
	"github.com/gofiber/fiber/v2"
)

// TinHandler is ...
type TinHandler struct {
}

// Index ...
func (TinHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/tin.html")
}

// Show user information
func (TinHandler) Show(ctx *fiber.Ctx) error {
	tin := ctx.Params("tin")

	tininfo, err := hits.TINLookup(tin)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "failure"})
	}

	// return ctx.JSON(fiber.Map{"data": userinfo})
	return ctx.JSON(tininfo)
}
