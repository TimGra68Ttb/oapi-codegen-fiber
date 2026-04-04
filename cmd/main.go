package main

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/api"
	"github.com/timgra68ttb/oapi-codegen-fiber/bootstrap"
	"github.com/timgra68ttb/oapi-codegen-fiber/cmd/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := bootstrap.NewInitializeBootstrap()
	f := fiber.New()

	f.Static("/swagger", "./swagger")
	f.Group("/api/v1.0")

	wrapper := handlers.NewServer(app)

	api.RegisterHandlers(f, wrapper)
	f.Listen(":3000")
}
