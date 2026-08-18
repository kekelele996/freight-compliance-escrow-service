package freight

import "time"

func scheduleOpeningOn(s ZonedSchedule, day time.Time) time.Time {
	return time.Date(day.Year(), day.Month(), day.Day(), s.OpenMinute/60, s.OpenMinute%60, 0, 0, time.UTC)
}
