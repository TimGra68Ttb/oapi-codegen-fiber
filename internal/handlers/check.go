package handlers

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/api"

	"github.com/gofiber/fiber/v2"
)

// checkHandler handles the logic for checking the application's status.
type checkHandler struct {
	// Add any necessary initialization for use cases or other components here.
}

// CheckHandler defines the interface for the checkController.
type CheckHandler interface {
	Execute(c *fiber.Ctx) error
}

// NewCheckHandler creates a new instance of checkController.
func NewCheckHandler() CheckHandler {
	return &checkHandler{}
}

// Execute implements the ICheckController interface.
func (*checkHandler) Execute(c *fiber.Ctx) error {
	var apiRes api.GlobalResponses
	apiRes.ResponseCode = "200"
	apiRes.ResponseMessage = "Go is running!"
	return c.JSON(apiRes)
}
