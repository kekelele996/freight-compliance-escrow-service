package freight

import "time"

func (s ZonedSchedule) NextOpenAt(t time.Time) time.Time {
	if s.ContainsAt(t) {
		return t
	}
	local := scheduleLocal(s, t)
	for i := 0; i < 8; i++ {
		day := local.AddDate(0, 0, i)
		if !s.Weekdays[day.Weekday()] {
			continue
		}
		candidate := scheduleOpeningOn(s, day)
		if candidate.Before(t) {
			continue
		}
		return candidate
	}
	return time.Time{}
}
