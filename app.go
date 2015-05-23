package app

import (
	"net/http"

	"github.com/elos/app/services"
	"github.com/elos/ehttp/serve"
	"github.com/gorilla/context"
)

type Middleware struct {
	UserAuth serve.Middleware
}

type Services struct {
	services.Agents

	services.DB

	services.Sessions
}

type App struct {
	router serve.Router
	*Middleware
	*Services
}

func New(m *Middleware, s *Services) *App {
	router := router(m, s)

	return &App{
		router:     router,
		Middleware: m,
		Services:   s,
	}
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context.ClearHandler(http.HandlerFunc(app.router.ServeHTTP)).ServeHTTP(w, r)
}
