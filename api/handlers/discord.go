package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/maraudery/ra/pkg/noauth/discord"
)

// DiscordHandler is ...
type DiscordHandler struct {
}

// Index ...
func (h DiscordHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/discord.html")
}

// Show token information
func (h DiscordHandler) Show(ctx *fiber.Ctx) error {
	token := ctx.Params("token")

	tokeninfo, err := discord.TokenLookup(token)
	if err != nil {
		return errors.New("error")
	}
	// return ctx.JSON(fiber.Map{"data": tokeninfo})
	return ctx.JSON(tokeninfo)
}
