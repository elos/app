package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elos/app/middleware"
	"github.com/elos/app/views"
	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/serve"
	"github.com/elos/ehttp/templates"
	"github.com/elos/models"
)

func SessionsGET(c *serve.Conn) {
	http.Redirect(c.ResponseWriter(), c.Request(), SessionsSignIn, 301)
}

func SessionsSignInGET(c *serve.Conn, db data.DB, sessions auth.Sessions) {
	templates.CatchError(c, views.Engine.Execute(c, views.SessionsSignIn, nil))
}

var userAuth = auth.Auth(auth.FormCredentialer)

func SessionsSignInPOST(c *serve.Conn, db data.DB, sessions auth.Sessions) {
	u, ok, err := userAuth(db, c.Request())

	if !ok {
		views.RenderSignIn(c, &views.Flash{"error", fmt.Sprintf("auth error: %s", err)})
		return
	}

	session := models.NewSessionForUser(u)
	session.SetID(db.NewID())
	if err := db.Save(session); err != nil {
		views.RenderSignIn(c, &views.Flash{"error", fmt.Sprintf("error saving session: %s", err)})
	}

	log.Printf("%+v", session)
	cookie := &http.Cookie{Name: middleware.ElosSessionAuth, Value: session.Token, Path: "/"}
	http.SetCookie(c, cookie)
	http.Redirect(c.ResponseWriter(), c.Request(), User, http.StatusFound)
}

func SessionsRegisterGET(c *serve.Conn, db data.DB) {
	templates.CatchError(c, views.Engine.Execute(c, views.SessionsRegister, nil))
}

func SessionsRegisterPOST(c *serve.Conn, db data.DB) {
	c.Write([]byte("not implemented"))
}
