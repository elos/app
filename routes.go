package app

import (
	"github.com/elos/app/apphandles"
	"github.com/elos/app/apptemplates"
	. "github.com/elos/app/conf"
	"github.com/elos/ehttp"
	"github.com/elos/ehttp/handles"
	"github.com/elos/transfer"
	"github.com/julienschmidt/httprouter"
)

const (
	GET  = ehttp.GET
	POST = ehttp.POST
)

func Route(a ehttp.Action, p string, h ehttp.Handle) *ehttp.Route {
	return &ehttp.Route{a, p, h}
}

// temporary
func FanOut(f httprouter.Handle) ehttp.Handle {
	return func(c *ehttp.Conn) {
		f(c.ResponseWriter(), c.Request(), *c.Params())
	}
}

func (app *App) Routes() []*ehttp.Route {
	return []*ehttp.Route{
		Route(GET, Index, FanOut(apptemplates.Show(apptemplates.SessionSignIn))),

		Route(GET, SessionSignIn, FanOut(apptemplates.Show(apptemplates.SessionSignIn))),
		Route(POST, SessionSignIn, FanOut(handles.Auth(apphandles.SignIn(app.sessions, UserCalendar), transfer.Auth(transfer.FormCredentialer), app.store))),

		Route(GET, SessionRegister, FanOut(apptemplates.Show(apptemplates.SessionRegister))),
		Route(POST, SessionRegister, FanOut(apphandles.RegisterHandle(app.store))),

		Route(GET, UserCalendar, FanOut(app.UserTemplate(apptemplates.RenderUserCalendar, app.store))),
		Route(GET, UserTasks, FanOut(app.UserTemplate(apptemplates.RenderUserEvents, app.store))),
		Route(GET, UserRoutines, FanOut(app.UserTemplate(apptemplates.RenderUserTasks, app.store))),
		Route(GET, UserSchedules, FanOut(app.UserTemplate(apptemplates.RenderUserSchedules, app.store))),
		Route(GET, UserSchedulesWeekly, FanOut(app.UserTemplate(apptemplates.RenderUserSchedulesWeekly, app.store))),
		Route(GET, UserSchedulesYearly, FanOut(app.UserTemplate(apptemplates.RenderUserSchedulesYearly, app.store))),

		Route(GET, UserSchedulesBase, FanOut(app.UserAuth(apphandles.UserSchedulesBase, app.store))),
		Route(GET, UserSchedulesBaseAddFixture, FanOut(app.UserTemplate(apptemplates.RenderUserSchedulesBaseAddFixture, app.store))),
		Route(POST, UserSchedulesBaseAddFixture, FanOut(app.UserAuth(apphandles.UserSchedulesBaseAddFixture, app.store))),

		//Route(GET, UserSchedulesBaseFixturesEdit, FanOut(app.UserAuth(apphandles.UserSchedulesBaseFixturesEdit, app.store))),
		//Route(POST, UserSchedulesBaseFixturesEdit, FanOut(app.UserAuth(apphandles.UserSchedulesBaseFixturesEdit, app.store))),

		Route(GET, UserSchedulesWeekday, FanOut(app.UserAuth(apphandles.UserSchedulesWeekday, app.store))),
		Route(GET, UserSchedulesYearday, FanOut(app.UserAuth(apphandles.UserSchedulesYearday, app.store))),
	}
}

func (app *App) ApplyRoutes() {
	routes := app.Routes()
	r := app.router

	for _, route := range routes {
		p := route.Path
		h := handles.FanIn(route.Handle)

		switch route.Action {
		case POST:
			r.POST(p, h)
		case GET:
			r.GET(p, h)
		}
	}
}
