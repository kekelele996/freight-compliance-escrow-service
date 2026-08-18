package freight

import "time"

type ServiceWindow struct {
	Location               *time.Location
	StartMinute, EndMinute int
	Weekdays               map[time.Weekday]bool
}

func (w ServiceWindow) Contains(t time.Time) bool {
	loc := w.Location
	if loc == nil {
		loc = time.UTC
	}
	v := t.In(loc)
	if len(w.Weekdays) > 0 && !w.Weekdays[v.Weekday()] {
		return false
	}
	m := v.Hour()*60 + v.Minute()
	if w.StartMinute == w.EndMinute {
		return true
	}
	if w.StartMinute < w.EndMinute {
		return m >= w.StartMinute && m < w.EndMinute
	}
	return m >= w.StartMinute || m < w.EndMinute
}
func (w ServiceWindow) NextOpen(t time.Time) time.Time {
	loc := w.Location
	if loc == nil {
		loc = time.UTC
	}
	p := t.In(loc)
	for i := 0; i < 8; i++ {
		if w.Weekdays == nil || w.Weekdays[p.Weekday()] {
			c := time.Date(p.Year(), p.Month(), p.Day(), w.StartMinute/60, w.StartMinute%60, 0, 0, loc)
			if w.Contains(p) {
				return p
			}
			if c.After(p) {
				return c
			}
		}
		p = time.Date(p.Year(), p.Month(), p.Day()+1, 0, 0, 0, 0, loc)
	}
	return t
}
