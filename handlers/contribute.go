package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// ContributeHandler is ...
type ContributeHandler struct {
}

// Index ...
func (h ContributeHandler) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).SendFile("./views/pages/contribute.html")
}
