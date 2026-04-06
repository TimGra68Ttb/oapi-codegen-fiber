package app

import (
	"log/slog"
	"os"

	"github.com/timgra68ttb/oapi-codegen-fiber/internal/repository"
	"github.com/timgra68ttb/oapi-codegen-fiber/internal/service"
)

type Application struct {
	Logger *slog.Logger
	Repo   service.DbInterface
}

func NewInitializeApp() Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	repo := repository.NewMapRepository()

	app := Application{
		Logger: logger,
		Repo:   repo,
	}
	return app
}
