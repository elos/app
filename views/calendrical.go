package views

import (
	"fmt"
	"time"

	"github.com/elos/data"
	"github.com/elos/models"
)

type (
	CalendarWeek struct {
		Days []*CalendarDay
	}

	CalendarDay struct {
		Header   string
		Fixtures []*CalendarFixture
	}

	CalendarFixture struct {
		ID        string
		Name      string
		Time      string
		RelStart  float32
		RelHeight float32
	}

	Schedule struct {
		HasSelectedFixture bool
		SelectedFixture    *models.Fixture
		Fixtures           []*CalendarFixture
	}
)

const (
	TodayHeader    = "Today"
	TomorrowHeader = "Tomorrow"
)

func CalendarHeader(t time.Time) string {
	switch t.YearDay() - time.Now().YearDay() {
	case 0:
		return TodayHeader
	case 1:
		return TomorrowHeader
	default:
		return t.Weekday().String()
	}
}

func FormattedHour(t time.Time) string {
	suffix := "AM"
	if t.Hour() > 11 {
		suffix = "PM"
	}

	extraZero := ""
	if t.Minute()%60 < 10 {
		extraZero = "0"
	}

	hour := t.Hour() % 12
	if hour == 0 {
		hour = 12
	}

	return fmt.Sprintf("%d:%s%d %s", hour, extraZero, t.Minute()%60, suffix)
}

func FormattedFixture(f *models.Fixture) string {
	return fmt.Sprintf("%s - %s", FormattedHour(f.StartTime), FormattedHour(f.EndTime))
}

const (
	minutesPerDay = 1440
)

func AbsMinute(t time.Time) int {
	return t.Hour()*60 + t.Minute()
}

func RelativeStartPosition(f *models.Fixture) float32 {
	return float32(AbsMinute(f.StartTime)) / minutesPerDay
}

func RelativeHeight(f *models.Fixture) float32 {
	return float32(AbsMinute(f.EndTime)-AbsMinute(f.StartTime)) / minutesPerDay
}

func MakeCalendarFixture(f *models.Fixture) *CalendarFixture {
	return &CalendarFixture{
		ID:        f.ID().String(),
		Name:      f.Name,
		Time:      FormattedFixture(f),
		RelStart:  RelativeStartPosition(f) * 100,
		RelHeight: RelativeHeight(f) * 100,
	}
}

func MakeCalendarFixtures(fs []*models.Fixture) []*CalendarFixture {
	calfs := make([]*CalendarFixture, len(fs))

	for i := range fs {
		calfs[i] = MakeCalendarFixture(fs[i])
	}

	return calfs
}

func MakeCalendarDay(db data.DB, s *models.Schedule) (*CalendarDay, error) {
	cd := new(CalendarDay)
	cd.Header = CalendarHeader(s.StartTime)

	fs, err := s.OrderedFixtures(db)
	if err != nil {
		return nil, err
	}

	cd.Fixtures = make([]*CalendarFixture, len(fs))
	for i, f := range fs {
		cd.Fixtures[i] = MakeCalendarFixture(f)
	}

	return cd, nil
}

func MakeCalendarWeek(db data.DB, cal *models.Calendar) (*CalendarWeek, error) {
	cw := new(CalendarWeek)
	cw.Days = make([]*CalendarDay, 5)

	now := time.Now()
	for i := 0; i < 5; i++ {
		//sched, err := cal.IntegratedSchedule(db, now)
		sched, err := cal.BaseSchedule(db)
		if err != nil {
			return nil, err
		}

		cd, err := MakeCalendarDay(db, sched)
		if err != nil {
			return nil, err
		}

		cw.Days[i] = cd
		now.Add(24 * time.Hour)
	}

	return cw, nil
}

/*
this code was for creating the calendar if a user didn't already have one
	cal, err := u.Calendar(a)
	if err != nil {
		if err == models.ErrEmptyRelationship {
			cal, _ = calendar.Create(a)
			u.SetCalendar(cal)
			a.Save(u)
			a.Save(cal)
		} else {
			return nil, err
		}
	}
*/

func MakeSchedule(db data.DB, sched models.Schedule) (*Schedule, error) {
	fixtures, err := sched.Fixtures(db)
	if err != nil {
		return nil, err
	}

	return &Schedule{
		Fixtures: MakeCalendarFixtures(fixtures),
	}, nil
}
