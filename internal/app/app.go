package app

import (
	"context"
	"vertical-slice-arch/internal/login"
	"vertical-slice-arch/internal/pkg/postgres"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	router    chi.Router
	loginFlow *login.Flow
}

func NewApplication() *Application {
	return new(Application)
}

func (app *Application) Init(ctx context.Context) error {
	postgresClient := postgres.NewClient()
	app.loginFlow = login.NewFlow(postgresClient)

	app.router = app.initRouter()
	return nil
}

func (app *Application) Run(ctx context.Context) error {
	return nil
}

func (app *Application) initRouter() chi.Router {
	router := chi.NewRouter()
	router.Post("/sign-in", app.loginFlow.Login)
	return router
}
