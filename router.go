package app

import (
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/elos/app/routes"
	"github.com/elos/ehttp/builtin"
	"github.com/elos/ehttp/serve"
)

var root string

func init() {
	_, filename, _, _ := runtime.Caller(1)
	root = filepath.Dir(filename)
}

func router(m *Middleware, s *Services) serve.Router {
	router := builtin.NewRouter()

	router.GET(routes.Landing, func(c *serve.Conn) {

		routes.LandingGET(c)

	})

	router.GET(routes.Sessions, func(c *serve.Conn) {

		routes.SessionsGET(c)

	})

	router.GET(routes.SessionsRegister, func(c *serve.Conn) {

		routes.SessionsRegisterGET(c, s.DB)

	})

	router.GET(routes.SessionsSignIn, func(c *serve.Conn) {

		routes.SessionsSignInGET(c, s.DB, s.Sessions)

	})

	router.GET(routes.User, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserGET(c)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserCalendar, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserCalendarGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserEvents, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserEventsGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserInteractive, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserInteractiveGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserRepl, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserReplGET(c, s.DB, s.Agents)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserRoutines, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserRoutinesGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedules, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedulesBase, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesBaseGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedulesBaseFixtures, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesBaseFixturesGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.UserSchedulesBaseFixtures, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesBaseFixturesDELETE(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedulesBaseFixturesCreate, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesBaseFixturesCreateGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedulesBaseFixturesDelete, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesBaseFixturesDeleteGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedulesBaseFixturesEdit, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesBaseFixturesEditGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedulesWeekly, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesWeeklyGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedulesWeeklyWeekday, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesWeeklyWeekdayGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedulesYearly, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesYearlyGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserSchedulesYearlyYearday, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserSchedulesYearlyYeardayGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.UserTasks, func(c *serve.Conn) {

		if ok := m.UserAuth.Inbound(c); !ok {
			return
		}

		routes.UserTasksGET(c, s.DB)

		if ok := m.UserAuth.Outbound(c); !ok {
			return
		}

	})

	router.ServeFiles("/css/*filepath", http.Dir(filepath.Join(root, "/assets/css/")))

	router.ServeFiles("/img/*filepath", http.Dir(filepath.Join(root, "/assets/img/")))

	router.ServeFiles("/js/*filepath", http.Dir(filepath.Join(root, "/assets/js/")))

	return router
}
