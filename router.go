package app

import (
	"net/http"

	"github.com/elos/app/routes"
	"github.com/elos/app/views"
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/auth"
	"github.com/elos/ehttp/builtin"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func router(db data.DB, sessions auth.Sessions, agents autonomous.Manager) serve.Router {
	router := builtin.NewRouter()
	userAuther := routes.NewUserAuthenticator(db, sessions)

	router.GET(routes.Landing, func(c *serve.Conn) {
		routes.LandingGET(c)
	})

	router.GET(routes.Sessions, func(c *serve.Conn) {
		routes.SessionsGET(c)
	})

	router.GET(routes.SessionsRegister, func(c *serve.Conn) {
		routes.SessionsRegisterGET(c, db)
	})

	router.POST(routes.SessionsRegister, func(c *serve.Conn) {
		routes.SessionsRegisterPOST(c, db)
	})

	router.GET(routes.SessionsSignIn, func(c *serve.Conn) {
		routes.SessionsSignInGET(c, db, sessions)
	})

	router.POST(routes.SessionsSignIn, func(c *serve.Conn) {
		routes.SessionsSignInPOST(c, db, sessions)
	})

	router.GET(routes.User, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserGET(c, u)
	}, userAuther))

	router.GET(routes.UserCalendar, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserCalendarGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserEvents, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserEventsGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserInteractive, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserInteractiveGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserRepl, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserReplGET(c, u, db, agents)
	}, userAuther))

	router.GET(routes.UserRoutines, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserRoutinesGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedules, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedulesBase, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesBaseGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedulesBaseFixtures, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesBaseFixturesGET(c, u, db)
	}, userAuther))

	router.POST(routes.UserSchedulesBaseFixtures, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesBaseFixturesPOST(c, u, db)
	}, userAuther))

	router.DELETE(routes.UserSchedulesBaseFixtures, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesBaseFixturesDELETE(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedulesBaseFixturesCreate, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesBaseFixturesCreateGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedulesBaseFixturesDelete, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesBaseFixturesDeleteGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedulesBaseFixturesEdit, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesBaseFixturesEditGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedulesWeekly, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesWeeklyGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedulesWeeklyWeekday, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesWeeklyWeekdayGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedulesYearly, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesYearlyGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserSchedulesYearlyYearday, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserSchedulesYearlyYeardayGET(c, u, db)
	}, userAuther))

	router.GET(routes.UserTasks, routes.UserAuth(func(c *serve.Conn, u *models.User) {
		routes.UserTasksGET(c, u, db)
	}, userAuther))

	router.ServeFiles("/img/*filepath", http.Dir(views.ImgDir))
	router.ServeFiles("/css/*filepath", http.Dir(views.CSSDir))
	router.ServeFiles("/js/*filepath", http.Dir(views.JSDir))

	return router
}
