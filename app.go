package app

import (
	"net/http"

	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/builtin"
	"github.com/elos/ehttp/serve"
)

type App struct {
	router   serve.Router
	db       data.DB
	sessions auth.Sessions
	agents   autonomous.Manager
}

func New(db data.DB, man autonomous.Manager) *App {
	sessions := builtin.NewSessions()
	router := router(db, sessions, man)

	return &App{
		router:   router,
		db:       db,
		sessions: sessions,
		agents:   man,
	}
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}
