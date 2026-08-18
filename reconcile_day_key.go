package freight

import "time"

func operationalDayKey(_ *time.Location, at time.Time) string { return at.UTC().Format("2006-01-02") }
