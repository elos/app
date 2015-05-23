package services

import (
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
)

type Agents interface {
	StartAgent(a autonomous.Agent)
}

type DB interface {
	data.DB
}

type Sessions interface {
	auth.Sessions
}
