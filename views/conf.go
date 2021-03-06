package views

import (
	"path/filepath"

	"github.com/elos/ehttp/templates"
)

var (
	appPath      = "github.com/elos/app"
	AssetsDir    = filepath.Join(templates.PackagePath(appPath), "assets")
	TemplatesDir = filepath.Join(AssetsDir, "templates")
	ImgDir       = filepath.Join(AssetsDir, "img")
	CSSDir       = filepath.Join(AssetsDir, "css")
	JSDir        = filepath.Join(AssetsDir, "js")
)

const (
	Index templates.Name = iota
	SessionsSignIn
	SessionsRegister
	SessionsAccountCreated

	User
	UserInteractive
	UserCalendar
	UserEvents
	UserTasks
	UserRoutines
	UserSchedules
	UserSchedulesBase
	UserSchedulesWeekly
	UserSchedulesYearly
	UserSchedulesWeekday
	UserSchedulesYearday
	UserSchedulesBaseAddFixture
)

var (
	layoutTemplate          string = "layout.tmpl"
	sessionsLayoutTemplate  string = "sessions/layout.tmpl"
	schedulesLayoutTemplate string = "user/schedules/layout.tmpl"
)

// Layout prepends variadic arguments with the layoutTemplate
func Layout(v ...string) []string {
	return templates.Prepend(layoutTemplate, v...)
}

// Sessions prepends variadic arguments with the layout and sessions templates
func Sessions(v ...string) []string {
	return Layout(templates.Prepend(sessionsLayoutTemplate, v...)...)
}

// Schedules prepends variadic arguments with the layout and schedules templates
func Schedules(v ...string) []string {
	return Layout(templates.Prepend(schedulesLayoutTemplate, templates.Prepend("user/schedules/common.tmpl", v...)...)...)
}

// Definition of the available templateSets for elos
// used in initialization of the templates, see: init.go
var templateSets = templates.TemplateSet{
	Index: Layout("index.tmpl"),

	SessionsSignIn:         Sessions("sessions/sign-in.tmpl"),
	SessionsRegister:       Sessions("sessions/register.tmpl"),
	SessionsAccountCreated: Sessions("sessions/account-created.tmpl"),

	User:            Layout("user/base.tmpl"),
	UserInteractive: Layout("user/interactive.tmpl"),
	UserCalendar:    Layout("user/schedules/common.tmpl", "user/calendar.tmpl"),
	UserEvents:      Layout("user/events.tmpl"),
	UserTasks:       Layout("user/tasks.tmpl"),
	UserRoutines:    Layout("user/routines.tmpl"),
	UserSchedules:   Layout("user/schedules.tmpl"),

	UserSchedulesBase:           Schedules("user/schedules/base.tmpl"),
	UserSchedulesBaseAddFixture: Schedules("user/schedules/base-add.tmpl"),

	UserSchedulesWeekly:  Layout("user/schedules/weekly.tmpl"),
	UserSchedulesYearly:  Layout("user/schedules/yearly.tmpl"),
	UserSchedulesWeekday: Layout("user/schedules/weekday.tmpl"),
	UserSchedulesYearday: Layout("user/schedules/yearday.tmpl"),
}
