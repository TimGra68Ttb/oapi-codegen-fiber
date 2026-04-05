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

type UserHandler interface {
	Create(c *fiber.Ctx) error
	GetById(c *fiber.Ctx, id string) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx, id string) error
	Delete(c *fiber.Ctx, id string) error
}

type UserService interface {
	Create(ctx context.Context, u domain.User) (domain.User, error)
	Get(ctx context.Context, id string) (domain.User, error)
	GetAll(ctx context.Context) []domain.User
	Update(ctx context.Context, id string, user domain.User) error
	Delete(ctx context.Context, id string) error
}

type userHandler struct {
	service UserService
}

// Vi skickar in repo här så controllern kan prata med databasen
func NewUserHandler(s UserService) UserHandler {
	return &userHandler{service: s}
}

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
		if errors.Is(err, domain.ErrUserNotFound){
			return c.Status(409).SendString(err.Error())
		}
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(user)
}

func (u *userHandler) GetAll(c *fiber.Ctx) error {
	users := u.service.GetAll(c.Context())
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
	/*
		if _, err := u.service.Get(c.Context(), id); err != nil {
			return c.Status(204).SendString(err.Error())
		}
		if err := u.service.Update(id, updatedUser); err != nil {
			return c.Status(500).SendString(err.Error())
		}*/

	if err := u.service.Update(c.Context(), id, dom); err != nil {
		return c.Status(204).SendString(err.Error())
	}
	return c.Status(200).JSON(fromDomain(dom))
}

func (u *userHandler) Delete(c *fiber.Ctx, id string) error {
	if err := u.service.Delete(c.Context(), id); err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

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
