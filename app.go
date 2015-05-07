package app

import (
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/builtin"
	"github.com/elos/ehttp/serve"
)

type App struct {
	router          serve.Router
	db              data.DB
	sessions        auth.Sessions
	*autonomous.Hub // inherits life from the hub
}

func New(db data.DB) *App {
	sessions := builtin.NewSessions()
	hub := autonomous.NewHub()
	router := router(db, sessions, hub)

	return &App{
		router:   router,
		db:       db,
		sessions: sessions,
		Hub:      hub,
	}
}
