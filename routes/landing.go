package routes

import (
	"github.com/elos/app/views"
	"github.com/elos/ehttp/serve"
	"github.com/elos/ehttp/templates"
)

func LandingGET(c *serve.Conn) {
	templates.CatchError(c, views.Engine.Execute(c, views.Index, nil))
}
