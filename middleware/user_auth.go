package middleware

import (
	"net/http"

	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
	"github.com/elos/models/user"
)

const (
	ElosAuth     = "elos-auth"
	userID       = "id"
	userKey      = "key"
	UserArtifact = "user"
)

type UserAuthenticator func(r *http.Request) (*models.User, error)

type UserAuth struct {
	auther   UserAuthenticator
	redirect string
}

func NewUserAuth(auther UserAuthenticator, redirect string) *UserAuth {
	return &UserAuth{auther, redirect}
}

func (ua *UserAuth) Inbound(c *serve.Conn) bool {
	user, err := ua.auther(c.Request())

	if err != nil {
		http.Redirect(c.ResponseWriter(), c.Request(), ua.redirect, 301)
		return false
	} else {
		c.AddContext(UserArtifact, user)
		return true
	}
}

func (ua *UserAuth) Outbound(c *serve.Conn) bool {
	return true
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
