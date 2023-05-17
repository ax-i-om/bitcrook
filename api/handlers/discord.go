package handlers

import (
	hits "github.com/ax-i-om/bitcrook/pkg/noauth/discord"
	"github.com/gofiber/fiber/v2"
)

// DiscordHandler is ...
type DiscordHandler struct {
}

// Index ...
func (DiscordHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/discord.html")
}

// Show user information
func (DiscordHandler) Show(ctx *fiber.Ctx) error {
	discord := ctx.Params("discord")

	discordinfo, err := hits.TokenLookup(discord)
	if err != nil {
		return ctx.JSON(fiber.Map{"message": "failure"})
	}

	// return ctx.JSON(fiber.Map{"data": userinfo})
	return ctx.JSON(discordinfo)
}
