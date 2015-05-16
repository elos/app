package views

import (
	"text/template"

	"github.com/elos/ehttp/templates"
)

var (
	funcs = template.FuncMap{
		"has":  templates.Has,
		"dict": templates.Dict,
		"css":  templates.CSS,
		"js":   templates.JS,
	}
)
