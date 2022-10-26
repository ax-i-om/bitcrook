package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// ToolsHandler is ...
type ToolsHandler struct {
}

// Index ...
func (ToolsHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/tools.html")
}
