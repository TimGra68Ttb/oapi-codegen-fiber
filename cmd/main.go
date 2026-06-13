package main

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/api"
	app "github.com/timgra68ttb/oapi-codegen-fiber/internal/application"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := app.NewInitializeApp()

	wrapper := handlers.NewServer(app)
	f := setupFiber(wrapper)

	f.Listen(":3000")
}

func setupFiber(wrapper api.ServerInterface) *fiber.App {
	f := fiber.New()

	f.Static("/swagger", "./swagger")
	f.Static("/swagger/openapi.yaml", "./internal/api/oapi_codegen.yml")
	apiGroup := f.Group("/api/v1")

	api.RegisterHandlers(apiGroup, wrapper)
	return f
}
