package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/timgra68ttb/oapi-codegen-fiber/api"
	"github.com/timgra68ttb/oapi-codegen-fiber/db"
)

type IUserController interface {
	Create(c *fiber.Ctx) error
	GetById(c *fiber.Ctx, id string) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx, id string) error
	Delete(c *fiber.Ctx, id string) error
}

type userController struct {
	repo db.DbInterface
}

// Vi skickar in repo här så controllern kan prata med databasen
func NewUserController(repo db.DbInterface) IUserController {
	return &userController{repo: repo}
}

func (u *userController) Create(c *fiber.Ctx) error {
	var newUser api.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	if newUser.Id == "" {
		newUser.Id = uuid.New().String()
	}

	if err := u.repo.Save(newUser); err != nil {
		return c.Status(500).JSON(newUser)
	}
	return c.Status(201).JSON(newUser)
}

func (u *userController) GetById(c *fiber.Ctx, id string) error {
	user, err := u.repo.Get(id)

	if err != nil {
		return c.Status(204).SendString("User not found")
	}
	return c.Status(200).JSON(user)
}

func (u *userController) GetAll(c *fiber.Ctx) error {
	users := u.repo.GetAll()
	if len(users) == 0 {
		return c.Status(204).JSON(users)
	}
	return c.Status(200).JSON(users)
}

func (u *userController) Update(c *fiber.Ctx, id string) error {
	var updatedUser api.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if id != updatedUser.Id {
		return c.Status(400).SendString(fmt.Sprintf("Id: %s, data: %v", id, updatedUser))
	}

	if _, err := u.repo.Get(id); err != nil {
		return c.Status(204).SendString(err.Error())
	}
	if err := u.repo.Update(id, updatedUser); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON(updatedUser)
}

func (u *userController) Delete(c *fiber.Ctx, id string) error {
	if err := u.repo.Delete(id); err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}
