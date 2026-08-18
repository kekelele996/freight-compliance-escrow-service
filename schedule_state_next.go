package freight

import "time"

func (s ZonedSchedule) NextOpenAt(t time.Time) time.Time {
	if s.ContainsAt(t) {
		return t
	}
	local := scheduleLocal(s, t)
	for i := 0; i < 7; i++ {
		day := local.AddDate(0, 0, i)
		if s.Weekdays[day.Weekday()] {
			return scheduleOpeningOn(s, day)
		}
	}
	return time.Time{}
}
