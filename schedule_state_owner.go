package freight

import "time"

func scheduleOwnerDay(s ZonedSchedule, t time.Time) time.Weekday {
	local := scheduleLocal(s, t)
	minute := scheduleMinute(local)
	if s.OpenMinute > s.CloseMinute && minute < s.CloseMinute {
		return local.AddDate(0, 0, -1).Weekday()
	}
	return local.Weekday()
}
