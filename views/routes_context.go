package views

import (
	"fmt"

	"github.com/elos/ehttp/templates"
)

type RoutesContext struct {
}

func (r *RoutesContext) Landing() string {
	return fmt.Sprintf("/")
}

func (r *RoutesContext) Sessions() string {
	return fmt.Sprintf("/sessions")
}

func (r *RoutesContext) SessionsRegister() string {
	return fmt.Sprintf("/sessions/register")
}

func (r *RoutesContext) SessionsSignIn() string {
	return fmt.Sprintf("/sessions/sign_in")
}

func (r *RoutesContext) User() string {
	return fmt.Sprintf("/user")
}

func (r *RoutesContext) UserCalendar() string {
	return fmt.Sprintf("/user/calendar")
}

func (r *RoutesContext) UserEvents() string {
	return fmt.Sprintf("/user/events")
}

func (r *RoutesContext) UserInteractive() string {
	return fmt.Sprintf("/user/interactive")
}

func (r *RoutesContext) UserRoutines() string {
	return fmt.Sprintf("/user/routines")
}

func (r *RoutesContext) UserSchedules() string {
	return fmt.Sprintf("/user/schedules")
}

func (r *RoutesContext) UserSchedulesBase() string {
	return fmt.Sprintf("/user/schedules/base")
}

func (r *RoutesContext) UserSchedulesBaseFixtures(fixture_id string) string {
	return fmt.Sprintf("/user/schedules/base/fixtures/%s", fixture_id)
}

func (r *RoutesContext) UserSchedulesBaseFixturesCreate(fixture_id string) string {
	return fmt.Sprintf("/user/schedules/base/fixtures/%s/create", fixture_id)
}

func (r *RoutesContext) UserSchedulesBaseFixturesDelete(fixture_id string) string {
	return fmt.Sprintf("/user/schedules/base/fixtures/%s/delete", fixture_id)
}

func (r *RoutesContext) UserSchedulesBaseFixturesEdit(fixture_id string) string {
	return fmt.Sprintf("/user/schedules/base/fixtures/%s/edit", fixture_id)
}

func (r *RoutesContext) UserSchedulesWeekly() string {
	return fmt.Sprintf("/user/schedules/weekly")
}

func (r *RoutesContext) UserSchedulesWeeklyWeekday(weekday string) string {
	return fmt.Sprintf("/user/schedules/weekly/%s", weekday)
}

func (r *RoutesContext) UserSchedulesYearly() string {
	return fmt.Sprintf("/user/schedules/yearly")
}

func (r *RoutesContext) UserSchedulesYearlyYearday(yearday string) string {
	return fmt.Sprintf("/user/schedules/yearly/%s", yearday)
}

func (r *RoutesContext) UserTasks() string {
	return fmt.Sprintf("/user/tasks")
}

var routesContext = &RoutesContext{}

type context struct {
	Routes *RoutesContext
	Data   interface{}
}

func (c *context) WithData(d interface{}) templates.Context {
	return &context{
		Routes: c.Routes,
		Data:   d,
	}
}

var globalContext = &context{
	Routes: routesContext,
}
