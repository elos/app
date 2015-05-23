package app

import (
	"net/http"

	"github.com/elos/app/services"
	"github.com/elos/ehttp/serve"
	"github.com/gorilla/context"
)

type App struct {
	router serve.Router

	agents services.Agents

	db services.DB

	sessions services.Sessions
}

func New(agents services.Agents, db services.DB, sessions services.Sessions) *App {
	router := router(agents, db, sessions)

	return &App{
		router: router,

		agents: agents,

		db: db,

		sessions: sessions,
	}
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context.ClearHandler(http.HandlerFunc(app.router.ServeHTTP)).ServeHTTP(w, r)
}
