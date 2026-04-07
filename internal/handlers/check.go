package handlers

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/api"

	"github.com/gofiber/fiber/v2"
)

// Types
type checkHandler struct {
}

// Constructors and injections - Startup
func NewCheckHandler() CheckHandler {
	return &checkHandler{}
}

// Implementations - of CheckHandler-interface.
func (*checkHandler) Execute(c *fiber.Ctx) error {
	var apiRes api.GlobalResponses
	apiRes.ResponseCode = "200"
	apiRes.ResponseMessage = "Go is running!"
	return c.JSON(apiRes)
}
