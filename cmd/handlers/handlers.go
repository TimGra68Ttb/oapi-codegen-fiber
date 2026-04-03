package handlers

import (
	"slices"

	"github.com/timgra68ttb/oapi-codegen-fiber/api"
	"github.com/timgra68ttb/oapi-codegen-fiber/bootstrap"
	"github.com/timgra68ttb/oapi-codegen-fiber/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// MyHandler represents the handler struct.
type MyHandler struct {
	Application bootstrap.Application
	Users       map[string]api.User
}

// NewServiceInitial creates a new instance of MyHandler.
func NewServiceInitial(app bootstrap.Application) MyHandler {
	return MyHandler{
		Application: app,
	}
}

// ServerInterfaceWrapper wraps controller interfaces.
type ServerInterfaceWrapper struct {
	CheckHandler controller.ICheckController
}

// Check implements api.ServerInterface.
func (h *ServerInterfaceWrapper) Check(c *fiber.Ctx) error {
	return h.CheckHandler.Execute(c)
}

// POST /api/v1.0/users
func (s *ServerInterfaceWrapper) CreateUser(c *fiber.Ctx) error {
	var newUser api.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	// Generera ID om det saknas
	if newUser.Id == "" {
		newUser.Id = uuid.New().String()
	}

	bootstrap.Users[newUser.Id] = newUser
	return c.Status(201).JSON(newUser)
}

// GET /api/v1.0/users/{id}
func (s *ServerInterfaceWrapper) GetUserById(c *fiber.Ctx, id string) error {
	user, Galler := bootstrap.Users[id]
	if !Galler {
		return c.Status(404).SendString("User not found")
	}
	return c.Status(200).JSON(user)
}

// GET /api/v1.0/users
func (s *ServerInterfaceWrapper) GetUsers(c *fiber.Ctx) error {
	users := []api.User{}
	for _, user := range bootstrap.Users {
		users = append(users, user)
	}
	slices.SortFunc(users, func(a, b api.User) int {
		if a.Id < b.Id {
			return -1
		}
		if a.Id > b.Id {
			return 1
		}
		return 0
	})

	return c.Status(200).JSON(users)
}

// PUT /api/v1.0/users/{id}
func (s *ServerInterfaceWrapper) UpdateUser(c *fiber.Ctx, id string) error {
	var updatedUser api.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	if _, exists := bootstrap.Users[id]; !exists {
		return c.Status(404).SendString("User not found")
	}

	updatedUser.Id = id // Säkerställ att ID matchar pathen
	bootstrap.Users[id] = updatedUser
	return c.Status(200).JSON(updatedUser)
}

// DELETE /api/v1.0/users/{id}
func (s *ServerInterfaceWrapper) DeleteUser(c *fiber.Ctx, id string) error {
	delete(bootstrap.Users, id)
	return c.SendStatus(204)
}