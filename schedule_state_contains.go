package freight

import "time"

func (s ZonedSchedule) ContainsAt(t time.Time) bool {
	local := scheduleLocal(s, t)
	m := scheduleMinute(local)
	day := local.Weekday()
	overnight := s.OpenMinute >= s.CloseMinute
	if s.Weekdays[day] {
		if overnight {
			if m >= s.OpenMinute {
				return true
			}
		} else if m >= s.OpenMinute && m < s.CloseMinute {
			return true
		}
	}
	if overnight {
		yesterday := day - 1
		if yesterday < 0 {
			yesterday = time.Saturday
		}
		if s.Weekdays[yesterday] && m < s.CloseMinute {
			return true
		}
	}
	return false
}
