package handlers

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/bootstrap"
	"github.com/timgra68ttb/oapi-codegen-fiber/controller"
	"github.com/timgra68ttb/oapi-codegen-fiber/db"

	"github.com/gofiber/fiber/v2"
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
	UserHandler  controller.IUserController
	Repo         db.DbInterface
}

func NewServerInterfaceWrapper(app bootstrap.Application) *ServerInterfaceWrapper {
	serve := NewServiceInitial(app)

	return &ServerInterfaceWrapper{
		CheckHandler: serve.CheckHandler(),
		UserHandler:  serve.UserController(),
		Repo:         app.Repo,
	}
}

// Check implements api.ServerInterface.
func (h *ServerInterfaceWrapper) Check(c *fiber.Ctx) error {
	return h.CheckHandler.Execute(c)
}

// POST /api/v1.0/users
func (s *ServerInterfaceWrapper) CreateUser(c *fiber.Ctx) error {
	return s.UserHandler.Create(c)
}

// GET /api/v1.0/users/{id}
func (s *ServerInterfaceWrapper) GetUserById(c *fiber.Ctx, id string) error {
	return s.UserHandler.GetById(c, id)
}

// GET /api/v1.0/users
func (s *ServerInterfaceWrapper) GetUsers(c *fiber.Ctx) error {
	return s.UserHandler.GetAll(c)
}

// PUT /api/v1.0/users/{id}
func (s *ServerInterfaceWrapper) UpdateUser(c *fiber.Ctx, id string) error {
	return s.UserHandler.Update(c, id)
}

// DELETE /api/v1.0/users/{id}
func (s *ServerInterfaceWrapper) DeleteUser(c *fiber.Ctx, id string) error {
	return s.UserHandler.Delete(c, id)
}
