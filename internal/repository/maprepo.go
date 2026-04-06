package repository

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"sync"

	"github.com/timgra68ttb/oapi-codegen-fiber/internal/domain"
)

// Types
type MapRepository struct {
	data map[string]domain.User
	mu   sync.RWMutex
}

// Constructs and injections
func NewMapRepository() *MapRepository {
	repo := &MapRepository{
		data: make(map[string]domain.User),
	}
	for i := range 10 {
		id := strconv.Itoa(i)
		email := fmt.Sprintf("user%s@example.com", id)
		repo.data[id] = domain.User{
			Id:    id,
			Email: email,
			Name:  fmt.Sprintf("User%s", id),
		}
	}
	return repo
}

// Implementations
func (r *MapRepository) Save(c context.Context, user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[user.Id] = user
	return nil
}

func (r *MapRepository) Get(c context.Context, id string) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.data[id]
	if !ok {
		return domain.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *MapRepository) GetAll(c context.Context) []domain.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	users := []domain.User{}
	for _, user := range r.data {
		users = append(users, user)
	}
	slices.SortFunc(users, func(a, b domain.User) int {
		if a.Id < b.Id {
			return -1
		}
		if a.Id > b.Id {
			return 1
		}
		return 0
	})
	return users
}

func (r *MapRepository) Delete(c context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
	return nil
}

func (r *MapRepository) Update(c context.Context, id string, user domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if id != user.Id {
		return fmt.Errorf("Data doesn't match the given key")
	}
	r.data[id] = user
	return nil
}

func (r *MapRepository) EmailExists(c context.Context, email string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, val := range r.data {
		if email == string(val.Email) {
			return true
		}
	}
	return false
}

func (r *MapRepository) UserExists(c context.Context, id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; ok {
		return true
	}
	return false
}
