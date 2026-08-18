package freight

import (
	"testing"
	"time"
)

func TestScheduleStateUsesOwnerDayAndFutureLocalCandidate(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	overnight := ZonedSchedule{Location: loc, OpenMinute: 22 * 60, CloseMinute: 3 * 60, Weekdays: map[time.Weekday]bool{time.Monday: true}}
	mon2315 := time.Date(2026, 8, 17, 23, 15, 0, 0, loc)
	tue0200 := time.Date(2026, 8, 18, 2, 0, 0, 0, loc)
	tue0300 := time.Date(2026, 8, 18, 3, 0, 0, 0, loc)
	if !overnight.ContainsAt(mon2315) || !overnight.ContainsAt(tue0200) || overnight.ContainsAt(tue0300) {
		t.Fatalf("contains=%v/%v/%v", overnight.ContainsAt(mon2315), overnight.ContainsAt(tue0200), overnight.ContainsAt(tue0300))
	}
	if owner := scheduleOwnerDay(overnight, tue0200); owner != time.Monday {
		t.Fatalf("owner=%v", owner)
	}
	wantOvernight := time.Date(2026, 8, 24, 22, 0, 0, 0, loc)
	if next := overnight.NextOpenAt(tue0300); !next.Equal(wantOvernight) || next.Location() != loc {
		t.Fatalf("overnight next=%s", next)
	}

	daytime := ZonedSchedule{Location: loc, OpenMinute: 9 * 60, CloseMinute: 17 * 60, Weekdays: map[time.Weekday]bool{time.Monday: true}}
	afterClose := time.Date(2026, 8, 17, 17, 30, 0, 0, loc)
	wantDaytime := time.Date(2026, 8, 24, 9, 0, 0, 0, loc)
	if next := daytime.NextOpenAt(afterClose); !next.Equal(wantDaytime) || !next.After(afterClose) {
		t.Fatalf("daytime next=%s", next)
	}
}
