package app

import (
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp"
	"github.com/elos/ehttp/handles"
	"github.com/elos/transfer"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

type App struct {
	autonomous.Life
	autonomous.Stopper
	autonomous.Managed

	*ehttp.Server
	hub *autonomous.Hub

	store               data.Store
	router              *httprouter.Router
	sessions            sessions.Store
	cookieCredentialer  transfer.Credentialer
	cookieAuthenticator transfer.Authenticator
}

func New(host string, port int, store data.Store) *App {
	sessionsStore := sessions.NewCookieStore([]byte("something-very-secret"), securecookie.GenerateRandomKey(32))
	cookieCredentialer := transfer.NewCookieCredentialer(sessionsStore)
	cookieAuth := transfer.Auth(cookieCredentialer)
	router := httprouter.New()
	server := ehttp.NewServer(host, port, router, store)

	return &App{
		Life:    autonomous.NewLife(),
		Stopper: make(autonomous.Stopper),

		hub: autonomous.NewHub(),

		Server: server,
		router: router,

		store:               store,
		sessions:            sessionsStore,
		cookieCredentialer:  cookieCredentialer,
		cookieAuthenticator: cookieAuth,
	}
}

func (app *App) UserAuth(f handles.AccessHandle, s data.Store) httprouter.Handle {
	return handles.Auth(f, app.cookieAuthenticator, s)
}

func (app *App) UserTemplate(f handles.TemplateHandle, s data.Store) httprouter.Handle {
	return app.UserAuth(handles.Template(f), s)
}

func (app *App) Start() {
	app.ApplyRoutes()

	app.Life.Begin()

	go app.hub.Start()
	go app.Server.Start()

	serverstop := make(chan bool, 1)

	go func() {
		app.Server.WaitStop()
		serverstop <- true
	}()

	select {
	case <-serverstop:
	case <-app.Stopper:
		app.Server.Stop()
		app.Server.WaitStop()
	}

	app.Life.End()
}
