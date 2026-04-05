package repository

import (
	"fmt"
	"slices"
	"strconv"
	"sync"

	"github.com/oapi-codegen/runtime/types"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/api"
)

type MapRepository struct {
	data map[string]api.User
	mu   sync.RWMutex
}

func NewMapRepository() *MapRepository {
	repo := &MapRepository{
		data: make(map[string]api.User),
	}
	for i := range 10 {
		id := strconv.Itoa(i)
		repo.data[id] = api.User{
			Id:    id,
			Email: types.Email(fmt.Sprintf("user%s@example.com", id)),
			Name:  fmt.Sprintf("User%s", id),
		}
	}
	return repo
}

func (r *MapRepository) Save(user api.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[user.Id] = user
	return nil
}

func (r *MapRepository) Get(id string) (api.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.data[id]
	if !ok {
		return api.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *MapRepository) GetAll() []api.User {
	r.mu.Lock()
	defer r.mu.Unlock()
	users := []api.User{}
	for _, user := range r.data {
		users = append(users, user)
	}
	slices.SortFunc(users, func(a, b api.User) int {
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

func (r *MapRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
	return nil
}

func (r *MapRepository) Update(id string, user api.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if id != user.Id {
		return fmt.Errorf("Data doesn't match the given key")
	}
	r.data[id] = user
	return nil
}
