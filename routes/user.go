package routes

import (
	"github.com/elos/app/views"
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/serve"
	"github.com/elos/ehttp/templates"
	"github.com/elos/models"
)

func UserGET(c *serve.Conn, u *models.User) {
	templates.CatchError(c, views.Engine.Execute(c, views.User, u))
}

func UserInteractiveGET(c *serve.Conn, u *models.User, db data.DB) {
	templates.CatchError(c, views.Engine.Execute(c, views.UserInteractive, u))
}

func UserReplGET(c *serve.Conn, u *models.User, db data.DB, agents autonomous.Manager) {
}

func UserCalendarGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserEventsGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserTasksGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserRoutinesGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesBaseGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesBaseFixturesGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesBaseFixturesPOST(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesBaseFixturesDELETE(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesBaseFixturesCreateGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesBaseFixturesEditGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesBaseFixturesDeleteGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesWeeklyGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesWeeklyWeekdayGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesYearlyGET(c *serve.Conn, u *models.User, db data.DB) {
}

func UserSchedulesYearlyYeardayGET(c *serve.Conn, u *models.User, db data.DB) {
}

/*

func UserSchedulesBase(w http.ResponseWriter, r *http.Request, p httprouter.Params, a data.Access) {
	c := transfer.NewHTTPConnection(w, r, a)
	templates.CatchError(c, userSchedulesBase(c, p))
}

func userSchedulesBase(c *transfer.HTTPConnection, p httprouter.Params) error {
	selectedFixtureID := c.Request().FormValue("selected_id")

	log.Print(selectedFixtureID)

	id, err := c.Access.ParseID(selectedFixtureID)
	if err != nil {
		log.Print(err)
		return views.RenderUserSchedulesBase(c, nil)
	}

	// we don't care  about this error
	// because the fixture will be nil if
	// the error isn't and then will just be ignored
	// which is what we want anyway
	f, _ := fixture.Find(c.Access, id)

	return views.RenderUserSchedulesBase(c, f)
}

func UserSchedulesBaseFixtures(w http.ResponseWriter, r *http.Request, p httprouter.Params, a data.Access) {
	c := transfer.NewHTTPConnection(w, r, a)
	templates.CatchError(c, userSchedulesBaseFixtures(c, p))
}

func userSchedulesBaseFixtures(c *transfer.HTTPConnection, p httprouter.Params) error {
	selectedFixtureID := p.ByName("fixture_id")

	if selectedFixtureID == "" {
		return userSchedulesBase(c, p)
	}

	id, err := c.Access.ParseID(selectedFixtureID)
	if err != nil {
		log.Print(err)
		views.RenderUserSchedulesBase(c, nil)
	}

	// err != nil => f == nil, which is acceptable
	f, _ := fixture.Find(c.Access, id)

	return views.RenderUserSchedulesBase(c, f)
}

func UserSchedulesWeekday(w http.ResponseWriter, r *http.Request, p httprouter.Params, a data.Access) {
	c := transfer.NewHTTPConnection(w, r, a)
	handles.CatchError(c, userSchedulesWeekday(c, p))
}

func userSchedulesWeekday(c *transfer.HTTPConnection, p httprouter.Params) error {
	weekday, err := strconv.Atoi(p.ByName("weekday"))
	if err != nil {
		return ehttp.NewMissingParamError("weekday")
	}

	if weekday < 0 || weekday > 6 {
		return ehttp.NewBadParamError("weekday", "must be in range 0-6 inclusive")
	}

	templates.CatchError(c, views.RenderUserSchedulesWeekday(c, weekday))
	return nil
}

func UserSchedulesYearday(w http.ResponseWriter, r *http.Request, p httprouter.Params, a data.Access) {
	c := transfer.NewHTTPConnection(w, r, a)
	handles.CatchError(c, userSchedulesYearday(c, p))
}

func userSchedulesYearday(c *transfer.HTTPConnection, p httprouter.Params) error {
	yearday, err := strconv.Atoi(p.ByName("yearday"))
	if err != nil {
		return ehttp.NewMissingParamError("yearday")
	}

	if yearday < 0 || yearday > 1231 {
		return ehttp.NewBadParamError("yearday", "must at least be in range 0-1231 inclusive to be potentially valid")
	}

	templates.CatchError(c, views.RenderUserSchedulesYearday(c, yearday))
	return nil
}

func UserSchedulesBaseAddFixture(w http.ResponseWriter, r *http.Request, p httprouter.Params, a data.Access) {
	c := transfer.NewHTTPConnection(w, r, a)
	handles.CatchError(c, userSchedulesBaseAddFixture(c, a))
}

var formTimeLayout = "15:04"

func userSchedulesBaseAddFixture(c *transfer.HTTPConnection, a data.Access) error {
	r := c.Request()

	params, err := getVals(r, "name", "start_time", "end_time")
	if err != nil {
		return err
	}

	*
		label, err := strconv.ParseBool(params["label"])
		if err != nil {
			return NewBadParamError("label", err.Error())
		}
	*

	start_time, err := time.Parse(formTimeLayout, params["start_time"])
	if err != nil {
		return ehttp.NewBadParamError("start_time", err.Error())
	}
	end_time, err := time.Parse(formTimeLayout, params["end_time"])
	if err != nil {
		return ehttp.NewBadParamError("end_time", err.Error())
	}

	cal, err := c.Client().(models.User).Calendar(persistence.ModelsStore(a))
	if err != nil {
		return err
	}

	s, err := cal.BaseSchedule(persistence.ModelsStore(a))
	if err != nil {
		return err
	}

	f := fixture.New(persistence.ModelsStore(a))

	f.SetName(params["name"])
	f.SetStartTime(start_time)
	f.SetEndTime(end_time)
	// f.SetLabel(label)

	if err = a.Save(f); err != nil {
		return err
	}

	s.IncludeFixture(f)

	a.Save(f)
	a.Save(s)

	http.Redirect(c.ResponseWriter(), c.Request(), "/user/schedules/base", http.StatusFound)
	return nil
}

func getVals(r *http.Request, v ...string) (map[string]string, error) {
	params := make(map[string]string)

	for _, v := range v {
		s := r.FormValue(v)
		if s == "" {
			return nil, ehttp.NewMissingParamError(v)
		}
		params[v] = s
	}

	return params, nil
}
*/
