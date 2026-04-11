package main

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/api"
	app "github.com/timgra68ttb/oapi-codegen-fiber/internal/application"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := app.NewInitializeApp()
	f := fiber.New()

	f.Static("/swagger", "./swagger")
	f.Group("/api/v1.0")

	wrapper := handlers.NewServer(app)

	api.RegisterHandlers(f, wrapper)
	f.Listen(":3000")
}
