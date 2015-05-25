package middleware

import (
	"log"
	"net/http"

	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const (
	ElosSessionAuth = "elos-sessions-auth"
	TokenContext    = "token"
)

type SessionAuth struct {
	db       data.DB
	sessions auth.Sessions
	redirect string
}

func NewSessionAuth(db data.DB, sessions auth.Sessions, redirect string) *SessionAuth {
	return &SessionAuth{
		db:       db,
		sessions: sessions,
		redirect: redirect,
	}
}

func (sa *SessionAuth) Inbound(c *serve.Conn) bool {
	log.Print("SessionAuth: Inbound")
	tokenCookie, err := c.Request().Cookie(ElosSessionAuth)
	log.Printf("%+v", c.Request().Cookies())
	if err != nil {
		log.Print("Could not find token cookie")
		return sa.Redirect(c)
	}

	log.Printf("%+v", tokenCookie)

	q := sa.db.NewQuery(models.SessionKind)
	log.Print(tokenCookie.Value)
	q.Select(data.AttrMap{"token": tokenCookie.Value})
	i, err := q.Execute()
	if err != nil {
		log.Print("error while executing query")
		return sa.Redirect(c)
	}

	session := models.Session{}
	if i.Next(&session) {
		log.Print("retrieved session objecT")

		if !session.Valid() {
			log.Print("session invalid")
			return sa.Redirect(c)
		}

		u, err := session.User(sa.db)
		if err != nil {
			log.Print("failed to retrieve user")
			return sa.Redirect(c)
		}

		c.AddContext(UserArtifact, u)
		log.Print("SUCCESS")
		return true
	} else {
		log.Print("session query yielded no results")
		return sa.Redirect(c)
	}
}

func (sa *SessionAuth) Redirect(c *serve.Conn) bool {
	http.Redirect(c, c.Request(), sa.redirect, http.StatusUnauthorized)
	return false
}

func (sa *SessionAuth) Outbound(c *serve.Conn) bool {
	// always good
	return true
}
