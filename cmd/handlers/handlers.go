package handlers

import (
	"fmt"

	"github.com/timgra68ttb/oapi-codegen-fiber/api"
	"github.com/timgra68ttb/oapi-codegen-fiber/bootstrap"
	"github.com/timgra68ttb/oapi-codegen-fiber/controller"
	"github.com/timgra68ttb/oapi-codegen-fiber/db"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// MyHandler represents the handler struct.
type MyHandler struct {
	Application bootstrap.Application
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
	Repo         db.DbRepository
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

	err := s.Repo.Save(newUser)
	if err != nil {
		return c.Status(500).JSON(newUser)
	}
	return c.Status(201).JSON(newUser)
}

// GET /api/v1.0/users/{id}
func (s *ServerInterfaceWrapper) GetUserById(c *fiber.Ctx, id string) error {
	user, err := s.Repo.Get(id)

	if err != nil {
		return c.Status(204).SendString("User not found")
	}
	return c.Status(200).JSON(user)
}

// GET /api/v1.0/users
func (s *ServerInterfaceWrapper) GetUsers(c *fiber.Ctx) error {
	users := s.Repo.GetAll()
	if len(users) == 0 {
		return c.Status(204).JSON(users)
	}
	return c.Status(200).JSON(users)
}

// PUT /api/v1.0/users/{id}
func (s *ServerInterfaceWrapper) UpdateUser(c *fiber.Ctx, id string) error {
	var updatedUser api.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if id != updatedUser.Id {
		return c.Status(400).SendString(fmt.Sprintf("Id: %s, data: %v", id, updatedUser))
	}

	if _, err := s.Repo.Get(id); err != nil {
		return c.Status(204).SendString(err.Error())
	}
	if err := s.Repo.Update(id, updatedUser); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(updatedUser)
}

// DELETE /api/v1.0/users/{id}
func (s *ServerInterfaceWrapper) DeleteUser(c *fiber.Ctx, id string) error {
	if err := s.Repo.Delete(id); err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}