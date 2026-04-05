package repository

import (
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/domain"
)

type DbInterface interface {
	Save(user domain.User) error
	Get(id string) (domain.User, error)
	GetAll() []domain.User
	Delete(id string) error
	Update(id string, user domain.User) error
	EmailExists(email string) bool
}
