package db

import "github.com/timgra68ttb/oapi-codegen-fiber/api"

type DbRepository interface {
	Save(user api.User) error
	Get(id string) (api.User, error)
	GetAll() []api.User
	Delete(id string) error
	Update(id string, user api.User) error
}
