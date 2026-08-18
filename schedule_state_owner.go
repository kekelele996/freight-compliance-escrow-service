package freight

import "time"

func scheduleOwnerDay(s ZonedSchedule, t time.Time) time.Weekday {
	local := scheduleLocal(s, t)
	m := scheduleMinute(local)
	if s.OpenMinute >= s.CloseMinute && m < s.CloseMinute {
		owner := local.Weekday() - 1
		if owner < 0 {
			owner = time.Saturday
		}
		return owner
	}
	return local.Weekday()
}
