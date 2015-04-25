package apptemplates

import (
	"github.com/elos/app/conf"
	"github.com/elos/ehttp/templates"
)

// our templates.Context implementation

type context struct {
	Routes conf.RoutesDefinition
	Data   interface{}
}

func (c *context) WithData(d interface{}) templates.Context {
	return &context{
		Routes: c.Routes,
		Data:   d,
	}
}
