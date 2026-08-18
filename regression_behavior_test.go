package freight

import "testing"

func TestFailedQuoteDoesNotPoisonRetryOrOtherTenant(t *testing.T) {
	m := NewQuoteMemo()
	if _, e := QuoteWithRetry(m, QuoteRequest{"north", "q1", 2, 10, -1}); e == nil {
		t.Fatal("invalid quote succeeded")
	}
	if m.Count() != 0 {
		t.Fatalf("failed candidate cached: %d", m.Count())
	}
	a, e := QuoteWithRetry(m, QuoteRequest{"north", "q1", 2, 10, 3})
	if e != nil || a != 23 {
		t.Fatalf("retry=%d err=%v", a, e)
	}
	b, e := QuoteWithRetry(m, QuoteRequest{"south", "q1", 3, 10, 4})
	if e != nil || b != 34 || m.Count() != 2 {
		t.Fatalf("tenant cache collision: %d %v count=%d", b, e, m.Count())
	}
	v, _ := m.load("north", "q1")
	v.Amount = 999
	again, _ := QuoteWithRetry(m, QuoteRequest{"north", "q1", 2, 10, 3})
	if again != 23 {
		t.Fatalf("cache alias mutated quote: %d", again)
	}
}
