package views

import "github.com/elos/ehttp/serve"

type Flash struct {
	Class, Msg string
}

type SessionsView struct {
	*Flash
}

func RenderSignIn(c *serve.Conn, f *Flash) error {
	return engine.Execute(c, SessionsSignIn, &SessionsView{f})
}
