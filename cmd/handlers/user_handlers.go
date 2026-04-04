package handlers

import "github.com/timgra68ttb/oapi-codegen-fiber/controller"

// I handlers-paketet (t.ex. i en ny fil user_handlers.go)
func (f *Factory) UserController() controller.IUserController {
	// Här skickar vi med Repo från Application-structen
	return controller.NewUserController(f.app.Repo)
}
