package apptemplates

import (
	"text/template"

	"github.com/elos/app/conf"
	"github.com/elos/ehttp/templates"
)

var (
	interpolate = func(r conf.RouteInterpolator, v string) string {
		return r(v)
	}

	funcs = template.FuncMap{
		"dict":        templates.Dict,
		"css":         templates.CSS,
		"js":          templates.JS,
		"interpolate": interpolate,
	}
)
