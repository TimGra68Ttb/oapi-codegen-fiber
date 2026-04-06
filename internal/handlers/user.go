package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oapi-codegen/runtime/types"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/api"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/domain"
)

// Service-interfaces - the requirements
type UserService interface {
	Create(ctx context.Context, u domain.User) (domain.User, error)
	Get(ctx context.Context, id string) (domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, id string, user domain.User) error
	Delete(ctx context.Context, id string) error
}

// Types
type userHandler struct {
	service UserService
}


// Constructors and injections - Startup
func NewUserHandler(s UserService) UserHandler {
	return &userHandler{service: s}
}

// Implementations - of userHandler-interface.
func (u *userHandler) Create(c *fiber.Ctx) error {
	var newUser api.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if newUser.Email == "" {
		return c.Status(400).SendString("Email missing")
	}
	createdUser, err := u.service.Create(c.Context(), toDomain(newUser))
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return c.Status(409).SendString("Användaren finns redan")
		}
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(fromDomain(createdUser))
}

func (u *userHandler) GetById(c *fiber.Ctx, id string) error {
	user, err := u.service.Get(c.Context(), id)

	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.Status(409).SendString(err.Error())
		}
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(user)
}

func (u *userHandler) GetAll(c *fiber.Ctx) error {
	users, err := u.service.GetAll(c.Context())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if len(users) == 0 {
		return c.Status(204).JSON(users)
	}
	return c.Status(200).JSON(users)
}

func (u *userHandler) Update(c *fiber.Ctx, id string) error {
	var updatedUser api.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if id != updatedUser.Id {
		return c.Status(400).SendString(fmt.Sprintf("Id: %s, data: %v", id, updatedUser))
	}
	dom := toDomain(updatedUser)

	err := u.service.Update(c.Context(), id, dom)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.Status(400).SendString(err.Error())
		}
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(fromDomain(dom))
}

func (u *userHandler) Delete(c *fiber.Ctx, id string) error {
	if err := u.service.Delete(c.Context(), id); err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

// Mappers
func toDomain(api api.User) domain.User {
	return domain.User{
		Id:    api.Id,
		Email: string(api.Email),
		Name:  api.Name,
	}
}

func fromDomain(domain domain.User) api.User {
	return api.User{
		Id:    domain.Id,
		Email: types.Email(domain.Email),
		Name:  domain.Name,
	}
}
