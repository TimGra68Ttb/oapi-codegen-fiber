package handlers

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/app"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/service"

	"github.com/gofiber/fiber/v2"
)

func (f *Factory) CheckHandler() CheckHandler {
	checkController := NewCheckHandler()
	return checkController
}
func (f *Factory) UserHandler() UserHandler {
	s := service.NewUserService(f.app.Repo)
	return NewUserHandler(s)
}

// Factory represents the handler struct.
type Factory struct {
	app app.Application
}

// NewFactory creates a new instance of MyHandler.
func NewFactory(app app.Application) Factory {
	return Factory{
		app: app,
	}
}

// Server wraps controller interfaces.
type Server struct {
	checkHandler CheckHandler
	userHandler  UserHandler
}

func NewServer(app app.Application) *Server {
	factory := NewFactory(app)

	return &Server{
		checkHandler: factory.CheckHandler(),
		userHandler:  factory.UserHandler(),
	}
}

// Check implements api.ServerInterface.
func (s *Server) Check(c *fiber.Ctx) error {
	return s.checkHandler.Execute(c)
}

// POST /api/v1.0/users
func (s *Server) CreateUser(c *fiber.Ctx) error {
	return s.userHandler.Create(c)
}

// GET /api/v1.0/users/{id}
func (s *Server) GetUserById(c *fiber.Ctx, id string) error {
	return s.userHandler.GetById(c, id)
}

// GET /api/v1.0/users
func (s *Server) GetUsers(c *fiber.Ctx) error {
	return s.userHandler.GetAll(c)
}

// PUT /api/v1.0/users/{id}
func (s *Server) UpdateUser(c *fiber.Ctx, id string) error {
	return s.userHandler.Update(c, id)
}

// DELETE /api/v1.0/users/{id}
func (s *Server) DeleteUser(c *fiber.Ctx, id string) error {
	return s.userHandler.Delete(c, id)
}
