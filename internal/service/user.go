package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/domain"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/repository"
)

type userService struct {
	repo repository.DbInterface
}

func NewUserService(repo repository.DbInterface) *userService {
	return &userService{repo: repo}
}

func (s *userService) Create(ctx context.Context, u domain.User) (domain.User, error) {
	// Flytta logiken för ID-generering hit!
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	if s.repo.EmailExists(u.Email) {
		return domain.User{}, domain.ErrUserAlreadyExists
	}
	err := s.repo.Save(u)
	return u, err
}

func (s *userService) Get(ctx context.Context, id string) (domain.User, error) {
	if _, err := s.repo.Get(id); err != nil {
		return domain.User{}, domain.ErrUserNotFound
	}
	return s.repo.Get(id)
}

func (s *userService) GetAll(ctx context.Context) []domain.User {
	return s.repo.GetAll()
}
func (s *userService) Update(ctx context.Context, id string, user domain.User) error {

	// Change repo to only accept user and change service to check if id matches
	return s.repo.Update(user.Id, user)
}
func (s *userService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}