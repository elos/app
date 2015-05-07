package views

import (
	"github.com/elos/ehttp/templates"
)

// our templates.Context implementation

type context struct {
	Data interface{}
}

func (c *context) WithData(d interface{}) templates.Context {
	return &context{
		Data: d,
	}
}
