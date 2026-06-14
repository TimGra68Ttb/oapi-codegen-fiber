package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/api"
	app "github.com/timgra68ttb/oapi-codegen-fiber/internal/application"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/debug"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/handlers"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := app.NewInitializeApp()

	debug.StartDebugServer(app.Repo)
	
	wrapper := handlers.NewServer(app)
	f := setupFiber(wrapper)

	port := fmt.Sprintf(":%s", util.GetEnvVariable("PORT", "3000"))
	f.Listen(port)
}

func setupFiber(wrapper api.ServerInterface) *fiber.App {
	f := fiber.New()

	f.Static("/swagger", "./swagger")
	f.Static("/swagger/openapi.yaml", "./internal/api/oapi_codegen.yml")
	apiGroup := f.Group("/api/v1")

	api.RegisterHandlers(apiGroup, wrapper)
	return f
}

func init() {
	godotenv.Load()
}
