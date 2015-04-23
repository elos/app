package apptemplates

import (
	"log"

	"github.com/elos/ehttp/templates"
	"github.com/elos/transfer"
	"github.com/julienschmidt/httprouter"
)

var context *templates.Context

var Context = context

func init() {
	// templateSets defined in conf.go
	context = templates.NewContext(TemplatesDir, &templateSets)
	if err := context.ParseHTMLTemplates(); err != nil {
		log.Fatal(err)
	}
}

func Render(c *transfer.HTTPConnection, name templates.Name, data interface{}) error {
	return context.Render(c, name, data)
}

func Show(name templates.Name) httprouter.Handle {
	return context.Show(name)
}
