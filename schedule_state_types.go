package freight

import "time"

type ZonedSchedule struct {
	Location                *time.Location
	OpenMinute, CloseMinute int
	Weekdays                map[time.Weekday]bool
}

func scheduleLocal(s ZonedSchedule, t time.Time) time.Time {
	if s.Location == nil {
		return t.UTC()
	}
	return t.In(s.Location)
}
func scheduleMinute(t time.Time) int { return t.Hour()*60 + t.Minute() }
