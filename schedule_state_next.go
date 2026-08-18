package freight

import "time"

func (s ZonedSchedule) NextOpenAt(t time.Time) time.Time {
	if s.ContainsAt(t) {
		return t
	}
	local := scheduleLocal(s, t)
	for i := 0; i < 14; i++ {
		day := local.AddDate(0, 0, i)
		if s.Weekdays[day.Weekday()] {
			c := scheduleOpeningOn(s, day)
			if c.After(t) {
				return c
			}
		}
	}
	return time.Time{}
}
