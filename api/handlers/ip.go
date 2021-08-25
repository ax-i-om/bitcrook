package handlers

import (
	"github.com/gofiber/fiber/v2"
	hits "github.com/maraudery/ra/pkg/noauth/ip"
)

// IpHandler is ...
type IpHandler struct {
}

// Index ...
func (h IpHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/ip.html")
}

// Show user information
func (h IpHandler) Show(ctx *fiber.Ctx) error {
	ip := ctx.Params("ip")

	ipinfo, err := hits.IPLookup(ip)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "failure"})
	}

	// return ctx.JSON(fiber.Map{"data": userinfo})
	return ctx.JSON(ipinfo)
}
