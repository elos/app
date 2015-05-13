package app

import (
	"net/http"

	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/builtin"
	"github.com/elos/ehttp/serve"
	"github.com/gorilla/context"
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
	context.ClearHandler(http.HandlerFunc(app.router.ServeHTTP)).ServeHTTP(w, r)
}
