package handlers

import (
	hits "github.com/bitcrook/cycull/pkg/noauth/ip"
	"github.com/gofiber/fiber/v2"
)

// IpHandler is ...
type IpHandler struct {
}

// Index ...
func (IpHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/ip.html")
}

// Show user information
func (IpHandler) Show(ctx *fiber.Ctx) error {
	ip := ctx.Params("ip")

	ipinfo, err := hits.IPLookup(ip)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "failure"})
	}

	// return ctx.JSON(fiber.Map{"data": userinfo})
	return ctx.JSON(ipinfo)
}
