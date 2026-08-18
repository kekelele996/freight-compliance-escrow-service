package freight

import (
	"errors"
	"testing"
	"time"
)

func TestOperationalDayRetriesWithoutDuplicateAcrossDST(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	b := NewReconcileBook()
	at := time.Date(2026, 3, 8, 0, 30, 0, 0, loc)
	calls := 0
	if e := b.RunDay(loc, at, func(key string) error { calls++; return errors.New("store unavailable") }); e == nil {
		t.Fatal("first run succeeded")
	}
	if b.Count() != 0 {
		t.Fatalf("failed day marked processed: %d", b.Count())
	}
	if e := b.RunDay(loc, at.Add(4*time.Hour), func(key string) error {
		calls++
		if key != "America/New_York/2026-03-08" {
			t.Fatalf("key=%s", key)
		}
		return nil
	}); e != nil {
		t.Fatal(e)
	}
	if b.Count() != 1 || calls != 2 {
		t.Fatalf("retry count processed=%d calls=%d", b.Count(), calls)
	}
	if e := b.RunDay(loc, at.Add(20*time.Hour), func(string) error { calls++; return nil }); e == nil {
		t.Fatal("same local day processed twice")
	}
	if b.Count() != 1 || calls != 2 {
		t.Fatalf("duplicate work ran: processed=%d calls=%d", b.Count(), calls)
	}
}
