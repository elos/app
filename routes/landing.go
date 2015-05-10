package routes

import (
	"log"

	"github.com/elos/app/views"
	"github.com/elos/ehttp/serve"
)

func LandingGET(c *serve.Conn) {
	if err := views.Engine.Execute(c.ResponseWriter(), views.Index, nil); err != nil {
		log.Print(err)
	}
}
