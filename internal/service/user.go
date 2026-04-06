package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/domain"
)

type DbInterface interface {
	Save(c context.Context, user domain.User) error
	Get(c context.Context, id string) (domain.User, error)
	GetAll(c context.Context) []domain.User
	Delete(c context.Context, id string) error
	Update(c context.Context, id string, user domain.User) error
	EmailExists(c context.Context, email string) bool
	UserExists(c context.Context, id string) bool
}

// Types
type userService struct {
	repo DbInterface
}

// Constructors and injections - startup
func NewUserService(repo DbInterface) *userService {
	return &userService{repo: repo}
}

// Implementations
func (s *userService) Create(ctx context.Context, u domain.User) (domain.User, error) {
	// Flytta logiken för ID-generering hit!
	if u.Id == "" {
		u.Id = uuid.New().String()
	}
	if s.repo.EmailExists(ctx, u.Email) {
		return domain.User{}, domain.ErrUserAlreadyExists
	}
	err := s.repo.Save(ctx, u)
	return u, err
}

func (s *userService) Get(ctx context.Context, id string) (domain.User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return domain.User{}, domain.ErrUserNotFound
	}
	return user, nil
}

func (s *userService) GetAll(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetAll(ctx), nil
}

func (s *userService) Update(ctx context.Context, id string, user domain.User) error {

	// Change repo to only accept user and change service to check if id matches
	if !s.repo.UserExists(ctx, user.Id) {
		return domain.ErrUserNotFound
	}
	return s.repo.Update(ctx, user.Id, user)
}

func (s *userService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
