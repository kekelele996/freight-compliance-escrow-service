package freight

import "time"

// operationalDayKey derives the per-site operational-day checkpoint key.
//
// The key is anchored to the *local* calendar date of `at` in `loc`, not the
// UTC date, so that a single operational day survives a DST transition even
// when its late-evening hours roll past midnight UTC. The location name is
// embedded to keep distinct sites (which may share a date) from colliding.
func operationalDayKey(loc *time.Location, at time.Time) string {
	if loc == nil {
		loc = time.UTC
	}
	return loc.String() + "/" + at.In(loc).Format("2006-01-02")
}
