package routes

import (
	"net/http"

	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
	"github.com/elos/models/user"
)

const (
	ElosAuth = "elos-auth"
	userID   = "id"
	userKey  = "key"
)

type (
	UserAuthenticator func(*http.Request) (*models.User, error)
	UserRoute         func(c *serve.Conn, u *models.User)
)

func UserAuth(route UserRoute, auther UserAuthenticator) serve.Route {
	return func(c *serve.Conn) {
		user, err := auther(c.Request())

		if err != nil {
			http.Redirect(c.ResponseWriter(), c.Request(), SessionsSignIn, 301)
		} else {
			route(c, user)
		}
	}
}

func NewUserAuthenticator(db data.DB, sessions auth.Sessions) UserAuthenticator {
	return func(r *http.Request) (*models.User, error) {
		session, err := sessions.Get(r, ElosAuth)
		if err != nil {
			return nil, err
		}

		id := session.Value(userID)
		key := session.Value(userKey)

		u, ok, err := user.Authenticate(db, id, key)
		if ok {
			return u, nil
		} else {
			return nil, err
		}
	}
}
