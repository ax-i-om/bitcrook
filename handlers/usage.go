package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// UsageHandler is ...
type UsageHandler struct {
}

// Index ...
func (UsageHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/usage.html")
}
