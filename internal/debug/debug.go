//go:build debug

package debug

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/service"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/util"
)

func StartDebugServer(repo service.DbInterface) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/debug", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"database": fiber.Map{
				"user": fiber.Map{
					"data": repo.GetAll(context.Background()),
				},
			},
		})
	})

	debugPort := util.GetEnvVariable("DEBUGPORT", "8081")
	fmt.Printf("[DEBUG] Server running on http://localhost:%s/debug/\n", debugPort)

	go func() {
		_ = app.Listen(":" + debugPort)
	}()
}
