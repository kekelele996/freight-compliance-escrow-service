package freight

import "time"

func operationalDayKey(loc *time.Location, at time.Time) string {
	if loc == nil {
		loc = time.UTC
	}
	local := at.In(loc)
	return loc.String() + "/" + local.Format("2006-01-02")
}
