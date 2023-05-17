package handlers

import (
	hits "github.com/ax-i-om/bitcrook/pkg/authfree/ip2location"
	"github.com/gofiber/fiber/v2"
)

// DomainHandler is ...
type DomainHandler struct {
}

// Index ...
func (DomainHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/domain.html")
}

// Show user information
func (DomainHandler) Show(ctx *fiber.Ctx) error {
	domain := ctx.Params("domain")

	domainInfo, err := hits.DomainLookup("API KEY GOES HERE, WILL IMPLEMENT KEYCONFIG FILE CONFIGURATION IN FUTURE", domain)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "failure"})
	}

	// return ctx.JSON(fiber.Map{"data": userinfo})
	return ctx.JSON(domainInfo)
}
