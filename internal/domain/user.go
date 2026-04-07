package domain

import "errors"

type User struct {
	Id    string
	Name  string
	Email string
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)
