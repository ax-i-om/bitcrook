package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/omniscient/pkg/noauth/userlookup"
)

// UsernameHandler is ...
type UsernameHandler struct {
}

// Index ...
func (h UsernameHandler) Index(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "404"})
}

// Show user information
func (h UsernameHandler) Show(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	userinfo := userlookup.UserLookup(username)

	// return ctx.JSON(fiber.Map{"data": userinfo})
	return ctx.JSON(userinfo)
}
