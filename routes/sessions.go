package routes

import (
	"net/http"

	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/serve"
)

func SessionsGET(c *serve.Conn) {
	http.Redirect(c.ResponseWriter(), c.Request(), SessionsSignIn, 301)
}

func SessionsSignInGET(c *serve.Conn, db data.DB, sessions auth.Sessions) {
	//views.Execute(c.ResponseWriter(), templates.SessionsSignIn, nil)
}

func SessionsSignInPOST(c *serve.Conn, db data.DB, sessions auth.Sessions) {
	/*
		sesh, err := sessions.Get(r, "elos-auth")
		if err != nil {
			return
		}
		err := session.Save(r, w)
		if err != nil {
			return
		} else {
			http.Redirect(w, r, User, http.StatusFound)
		}
	*/
}

func SessionsRegisterGET(c *serve.Conn, db data.DB) {
}

func SessionsRegisterPOST(c *serve.Conn, db data.DB) {
}
