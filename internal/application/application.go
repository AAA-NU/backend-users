package application

import (
	"log/slog"

	"github.com/aaanu/backendusers/internal/server"
	"github.com/aaanu/backendusers/internal/service"
	"github.com/aaanu/backendusers/internal/storage"
)

type Application struct {
	server *server.Server
}

func New(log *slog.Logger) *Application {
	storage := storage.New()

	service := service.New(log, storage)

	server := server.New(service)

	return &Application{
		server: server,
	}
}

func (a *Application) Start() {
	a.server.Start()
}

func (a *Application) GracefulStop() {
	a.server.GracefulStop()
}
