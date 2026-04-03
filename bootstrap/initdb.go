package bootstrap

import (
	"fmt"
	"strconv"

	"github.com/oapi-codegen/runtime/types"
	"github.com/timgra68ttb/oapi-codegen-fiber/api"
)

var Users map[string]api.User = make(map[string]api.User)

func InitDb() {
	for i := range 10 {
		id := strconv.Itoa(i)
		email := fmt.Sprintf("user%s@example.com", id)
		name := fmt.Sprintf("User%s", id)
		Users[id] = api.User{
			Id:    id,
			Email: types.Email(email),
			Name:  name,
		}
	}
}
