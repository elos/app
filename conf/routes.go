package conf

import (
	"path/filepath"
	"strings"
)

var join = filepath.Join

var (
	Index = "/"

	SessionBase           = join(Index, "session")
	SessionSignIn         = join(SessionBase, "sign-in")
	SessionRegister       = join(SessionBase, "register")
	SessionAccountCreated = join(SessionBase, "account-created")

	UserBase        = join(Index, "user")
	UserInteractive = join(UserBase, "interactive")
	UserREPL        = join(UserBase, "repl")
	UserCalendar    = join(UserBase, "calendar")
	UserEvents      = join(UserBase, "events")
	UserTasks       = join(UserBase, "tasks")
	UserRoutines    = join(UserBase, "routines")

	UserSchedules        = join(UserBase, "schedules")
	UserSchedulesBase    = join(UserSchedules, "base")
	UserSchedulesWeekly  = join(UserSchedules, "weekly")
	UserSchedulesYearly  = join(UserSchedules, "yearly")
	UserSchedulesWeekday = join(UserSchedulesWeekly, ":weekday")
	UserSchedulesYearday = join(UserSchedulesYearly, ":yearday")

	UserSchedulesBaseFixtures       = UserSchedulesBase + "/fixtures/:fixture_id"
	UserSchedulesBaseFixturesEdit   = UserSchedulesBaseFixtures + "/edit"
	UserSchedulesBaseFixturesDelete = UserSchedulesBaseFixtures + "/delete"
	UserSchedulesBaseAddFixture     = UserSchedulesBase + "/add_fixture"
)

type RouteInterpolator func(string) string

type RoutesDefinition struct {
	Index,

	SessionBase,
	SessionSignIn,
	SessionRegister,
	SessionAccountCreated,

	UserBase,
	UserInteractive,
	UserREPL,
	UserCalendar,
	UserEvents,
	UserTasks,
	UserRoutines,

	UserSchedules,
	UserSchedulesBase,
	UserSchedulesWeekly,
	UserSchedulesYearly,
	UserSchedulesWeekday,
	UserSchedulesYearday,

	UserSchedulesBaseFixturesEdit,
	UserSchedulesBaseFixturesDelete,
	UserSchedulesBaseAddFixture string

	UserSchedulesBaseFixtures RouteInterpolator
}

var Routes = RoutesDefinition{
	Index: Index,

	SessionBase:           SessionBase,
	SessionSignIn:         SessionSignIn,
	SessionRegister:       SessionRegister,
	SessionAccountCreated: SessionAccountCreated,

	UserBase:             UserBase,
	UserInteractive:      UserInteractive,
	UserREPL:             UserREPL,
	UserSchedulesBase:    UserSchedulesBase,
	UserSchedulesWeekly:  UserSchedulesWeekly,
	UserSchedulesYearly:  UserSchedulesYearly,
	UserSchedulesWeekday: UserSchedulesWeekday,
	UserSchedulesYearday: UserSchedulesYearday,

	UserSchedulesBaseFixturesEdit:   UserSchedulesBaseFixturesEdit,
	UserSchedulesBaseFixturesDelete: UserSchedulesBaseFixturesDelete,
	UserSchedulesBaseAddFixture:     UserSchedulesBaseAddFixture,

	UserSchedulesBaseFixtures: func(id string) string { return strings.Replace(UserSchedulesBaseFixtures, ":fixture_id", id, 1) },
}
