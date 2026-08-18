package freight

import "time"

func scheduleOpeningOn(s ZonedSchedule, day time.Time) time.Time {
	local := scheduleLocal(s, day)
	return time.Date(local.Year(), local.Month(), local.Day(), s.OpenMinute/60, s.OpenMinute%60, 0, 0, s.Location)
}
