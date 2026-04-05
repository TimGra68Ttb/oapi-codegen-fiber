package handlers

// I handlers-paketet (t.ex. i en ny fil user_handlers.go)
func (f *Factory) UserController() IUserController {
	// Här skickar vi med Repo från Application-structen
	return NewUserController(f.app.Repo)
}
