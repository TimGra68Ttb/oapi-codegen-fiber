package handlers

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/bootstrap"
	"github.com/timgra68ttb/oapi-codegen-fiber/controller"

	"github.com/gofiber/fiber/v2"
)

// Factory represents the handler struct.
type Factory struct {
	app bootstrap.Application
}

// NewFactory creates a new instance of MyHandler.
func NewFactory(app bootstrap.Application) Factory {
	return Factory{
		app: app,
	}
}

// Server wraps controller interfaces.
type Server struct {
	checkController controller.ICheckController
	userController  controller.IUserController
}

func NewServer(app bootstrap.Application) *Server {
	factory := NewFactory(app)

	return &Server{
		checkController: factory.CheckController(),
		userController:  factory.UserController(),
	}
}

// Check implements api.ServerInterface.
func (s *Server) Check(c *fiber.Ctx) error {
	return s.checkController.Execute(c)
}

// POST /api/v1.0/users
func (s *Server) CreateUser(c *fiber.Ctx) error {
	return s.userController.Create(c)
}

// GET /api/v1.0/users/{id}
func (s *Server) GetUserById(c *fiber.Ctx, id string) error {
	return s.userController.GetById(c, id)
}

// GET /api/v1.0/users
func (s *Server) GetUsers(c *fiber.Ctx) error {
	return s.userController.GetAll(c)
}

// PUT /api/v1.0/users/{id}
func (s *Server) UpdateUser(c *fiber.Ctx, id string) error {
	return s.userController.Update(c, id)
}

// DELETE /api/v1.0/users/{id}
func (s *Server) DeleteUser(c *fiber.Ctx, id string) error {
	return s.userController.Delete(c, id)
}
