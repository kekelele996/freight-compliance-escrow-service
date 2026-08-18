package freight

import "time"

func (s ZonedSchedule) ContainsAt(t time.Time) bool {
	local := scheduleLocal(s, t)
	m := scheduleMinute(local)
	owner := scheduleOwnerDay(s, t)
	if !s.Weekdays[owner] {
		return false
	}
	if s.OpenMinute == s.CloseMinute {
		return true
	}
	if s.OpenMinute < s.CloseMinute {
		return m >= s.OpenMinute && m < s.CloseMinute
	}
	return m >= s.OpenMinute || m < s.CloseMinute
}
