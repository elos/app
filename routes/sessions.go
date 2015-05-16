package routes

import (
	"log"
	"net/http"

	"github.com/elos/app/views"
	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/serve"
	"github.com/elos/ehttp/templates"
)

func SessionsGET(c *serve.Conn) {
	http.Redirect(c.ResponseWriter(), c.Request(), SessionsSignIn, 301)
}

func SessionsSignInGET(c *serve.Conn, db data.DB, sessions auth.Sessions) {
	templates.CatchError(c, views.Engine.Execute(c, views.SessionsSignIn, nil))
}

func SessionsSignInPOST(c *serve.Conn, db data.DB, sessions auth.Sessions) {
	sesh, err := sessions.Get(c.Request(), ElosAuth)
	if err != nil {
		if sesh, err = sessions.New(c.Request(), ElosAuth); err != nil {
			log.Print("Error getting session: ", err)
			return
		}
	}
	auther := auth.Auth(auth.FormCredentialer)
	u, ok, err := auther(db, c.Request())
	if !ok {
		views.RenderSignIn(c, &views.Flash{"error", err.Error()})
		return
	}

	sesh.SetValue(userID, u.ID().String())
	sesh.SetValue(userKey, u.Key)
	if err := sesh.Save(c.Request(), c); err != nil {
		views.RenderSignIn(c, &views.Flash{"error", err.Error()})
	} else {
		http.Redirect(c.ResponseWriter(), c.Request(), User, http.StatusFound)
	}
}

func SessionsRegisterGET(c *serve.Conn, db data.DB) {
	templates.CatchError(c, views.Engine.Execute(c, views.SessionsRegister, nil))
}

func SessionsRegisterPOST(c *serve.Conn, db data.DB) {
}
