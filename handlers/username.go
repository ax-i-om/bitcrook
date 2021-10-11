package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/goseek/pkg/noauth/userlookup"
)

// UsernameHandler is ...
type UsernameHandler struct {
}

// Index ...
func (UsernameHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/username.html")
}

// Show user information
func (UsernameHandler) Show(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	userinfo := userlookup.UserLookup(username)

	// return ctx.JSON(fiber.Map{"data": userinfo})
	return ctx.JSON(userinfo)
}
