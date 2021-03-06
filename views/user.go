package views

/*

func RenderUserCalendar(c *transfer.HTTPConnection) error {
	u, ok := c.Client().(models.User)
	if !ok {
		return models.CastError(models.UserKind)
	}

	cal, err := u.Calendar(persistence.ModelsStore(c.Access))
	if err != nil {
		return err
	}

	cw, err := MakeCalendarWeek(c.Access, cal)
	if err != nil {
		return err
	}

	return engine.Render(c, UserCalendar, cw)
}

func RenderUserEvents(c *transfer.HTTPConnection) error {
	return engine.Render(c, UserEvents, c.Client().(models.User))
}

func RenderUserTasks(c *transfer.HTTPConnection) error {
	return engine.Render(c, UserTasks, c.Client().(models.User))
}

func RenderUserRoutines(c *transfer.HTTPConnection) error {
	return engine.Render(c, UserRoutines, c.Client().(models.User))
}

func RenderUserSchedules(c *transfer.HTTPConnection) error {
	return engine.Render(c, UserSchedules, c.Client().(models.User))
}

func RenderUserSchedulesBase(c *transfer.HTTPConnection, selectedFixture models.Fixture) error {
	sv, err := userSchedulesBaseView(c)
	if err != nil {
		return err
	}

	sv.SelectedFixture = selectedFixture
	if selectedFixture != nil {
		sv.HasSelectedFixture = true
	}

	log.Print("%+v", pathToTemplate)

	return engine.Render(c, UserSchedulesBase, sv)
}

func userSchedulesBaseView(c *transfer.HTTPConnection) (*Schedule, error) {
	u := c.Client().(models.User)
	a := c.Access

	cal, err := u.Calendar(persistence.ModelsStore(a))
	if err != nil {
		if err == models.ErrEmptyRelationship {
			if err = user.NewCalendar(persistence.ModelsStore(a), u); err != nil {
				return nil, templates.NewServerError(err)
			}
		} else {
			return nil, templates.NewServerError(err)
		}
	}

	sch, err := cal.BaseSchedule(persistence.ModelsStore(a))
	if err != nil {
		if err == models.ErrEmptyRelationship {
			if err = calendar.NewBaseSchedule(persistence.ModelsStore(a), cal); err != nil {
				return nil, templates.NewServerError(err)
			}
		} else {
			return nil, templates.NewServerError(err)
		}
	}

	return MakeSchedule(a, sch)
}

func RenderUserSchedulesBaseAddFixture(c *transfer.HTTPConnection) error {
	sv, err := userSchedulesBaseView(c)
	if err != nil {
		return err
	}

	return engine.Render(c, UserSchedulesBaseAddFixture, sv)
}

func RenderUserSchedulesWeekly(c *transfer.HTTPConnection) error {
	return engine.Render(c, UserSchedulesWeekly, c.Client().(models.User))
}

func RenderUserSchedulesYearly(c *transfer.HTTPConnection) error {
	return engine.Render(c, UserSchedulesYearly, c.Client().(models.User))
}

func RenderUserSchedulesWeekday(c *transfer.HTTPConnection, weekday int) error {
	return engine.Render(c, UserSchedulesWeekday, c.Client().(models.User))
}

func RenderUserSchedulesYearday(c *transfer.HTTPConnection, yearday int) error {
	return engine.Render(c, UserSchedulesYearday, c.Client().(models.User))
}

func RenderUserInteractive(c *transfer.HTTPConnection) error {
	return engine.Render(c, UserInteractive, c.Client().(models.User))
}
*/
