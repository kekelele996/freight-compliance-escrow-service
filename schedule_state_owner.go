package freight

import "time"

func scheduleOwnerDay(_ ZonedSchedule, t time.Time) time.Weekday { return t.Weekday() }
