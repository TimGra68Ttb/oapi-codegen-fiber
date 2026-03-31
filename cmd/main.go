package main

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/api"
	"github.com/timgra68ttb/oapi-codegen-fiber/bootstrap"
	"github.com/timgra68ttb/oapi-codegen-fiber/cmd/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the application's bootstrap components.
	app := bootstrap.NewInitializeBootstrap()

	// Create a new Fiber instance.
	f := fiber.New()

	// Serve the Swagger UI at the "/swagger" endpoint.
	f.Static("/swagger", "./swagger")

	// Define the API versioned group.
	f.Group("/api/v1.0")

	// Initialize the application-specific handlers.
	serve := handlers.NewServiceInitial(app)
	checkController := serve.CheckHandler()

	// Create a wrapper for the controller interfaces.
	wrapper := &handlers.ServerInterfaceWrapper{
		CheckHandler: checkController,
	}

	// Register API handlers.
	api.RegisterHandlers(f, wrapper)

	// Start the server on port 3000.
	f.Listen(":3000")
}
