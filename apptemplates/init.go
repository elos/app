package apptemplates

import (
	"log"

	"github.com/elos/app/conf"
	"github.com/elos/ehttp/templates"
	"github.com/elos/transfer"
	"github.com/julienschmidt/httprouter"
)

var engine *templates.Engine

var Engine = engine

func init() {
	// templateSets defined in conf.go
	engine = templates.
		NewEngine(TemplatesDir, &templateSets).
		WithEveryLoad().
		WithContext(&context{Routes: conf.Routes}).
		WithFuncMap(funcs)

	if err := engine.ParseHTMLTemplates(); err != nil {
		log.Fatal(err)
	}
}

func Render(c *transfer.HTTPConnection, name templates.Name, data interface{}) error {
	return engine.Render(c, name, data)
}

func Show(name templates.Name) httprouter.Handle {
	return engine.Show(name)
}
