package bootstrap

import (
	"log/slog"
	"os"

	"github.com/timgra68ttb/oapi-codegen-fiber/db"
)

type Application struct {
	Logger *slog.Logger
	Repo   db.DbInterface
}

func NewInitializeBootstrap() Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	repo := db.NewMapRepository()

	app := Application{
		Logger: logger,
		Repo:   repo,
	}
	return app
}
