
package app

import (
    "github.com/elos/ehttp/builtin"
    "github.com/elos/app/routes"
)

func router( agents services.Agents,  db services.DB,  sessions services.Sessions, ) serve.Router {
    router := builtin.NewRouter()
    userAuther := routes.NewUserAuthenticator(db, sessions)

    
            
            
            router.GET(routes.Landing, func(c *serve.Conn) {
                    routes.LandingGET(c)
                })
            
            
    
            
            
            router.GET(routes.Sessions, func(c *serve.Conn) {
                    routes.SessionsGET(c)
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
                    routes.UserReplGET(c, u, agents, db)
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
            
            
    

    
    router.ServeFiles("css/*filepath", http.Dir("assets/css"))
    
    router.ServeFiles("img/*filepath", http.Dir("assets/img"))
    
    router.ServeFiles("js/*filepath", http.Dir("assets/js"))
    

    return router
}
