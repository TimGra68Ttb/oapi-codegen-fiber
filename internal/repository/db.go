package repository

import (
	"context"

	"github.com/timgra68ttb/oapi-codegen-fiber/internal/domain"
)

type DbInterface interface {
	Save(c context.Context, user domain.User) error
	Get(c context.Context, id string) (domain.User, error)
	GetAll(c context.Context, ) []domain.User
	Delete(c context.Context, id string) error
	Update(c context.Context, id string, user domain.User) error
	EmailExists(c context.Context, email string) bool
	UserExists(c context.Context, id string) bool
}
