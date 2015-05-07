package views

import (
	"log"

	"github.com/elos/ehttp/templates"
)

var engine *templates.Engine

var Engine = engine

func init() {
	// templateSets defined in conf.go
	engine = templates.
		NewEngine(TemplatesDir, &templateSets).
		WithEveryLoad().
		WithContext(&context{}).
		WithFuncMap(funcs)

	if err := engine.Parse(); err != nil {
		log.Fatal(err)
	}
}
