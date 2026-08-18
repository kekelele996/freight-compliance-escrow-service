package freight

import "time"

func (s ZonedSchedule) ContainsAt(t time.Time) bool {
	local := scheduleLocal(s, t)
	m := scheduleMinute(local)
	return s.Weekdays[local.Weekday()] && m >= s.OpenMinute && m < s.CloseMinute
}
