package conf

import "path/filepath"

var join = filepath.Join

var (
	Index = "/"

	SessionBase           = join(Index, "session")
	SessionSignIn         = join(SessionBase, "sign-in")
	SessionRegister       = join(SessionBase, "register")
	SessionAccountCreated = join(SessionBase, "account-created")

	UserBase     = join(Index, "user")
	UserCalendar = join(UserBase, "calendar")
	UserEvents   = join(UserBase, "events")
	UserTasks    = join(UserBase, "tasks")
	UserRoutines = join(UserBase, "routines")

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
