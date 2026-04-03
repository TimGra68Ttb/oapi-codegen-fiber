package bootstrap

import "github.com/gofiber/fiber/v2/log"

type Application struct {
	Logger *log.Logger
}

func NewInitializeBootstrap() Application {
	InitDb()
	app := Application{}
	return app
}
