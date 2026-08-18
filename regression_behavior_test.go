package freight

import (
	"sync"
	"testing"
)

func TestRequestLeaseLifecycleIsTenantScopedAndAtomic(t *testing.T) {
	s := NewRequestLeaseStore()
	start := make(chan struct{})
	type reserveResult struct {
		tenant string
		lease  RequestLease
		err    error
	}
	results := make(chan reserveResult, 2)
	var wg sync.WaitGroup
	for _, tc := range []struct{ tenant, fingerprint string }{{" North ", "fa"}, {"south", "fb"}} {
		tc := tc
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			lease, err := s.ReserveLease(tc.tenant, "same", tc.fingerprint)
			results <- reserveResult{tenant: CanonicalTenant(tc.tenant), lease: lease, err: err}
		}()
	}
	close(start)
	wg.Wait()
	close(results)

	leases := map[string]RequestLease{}
	for result := range results {
		if result.err != nil {
			t.Fatalf("concurrent reserve for %s failed: %v", result.tenant, result.err)
		}
		leases[result.tenant] = result.lease
	}
	north, northOK := leases["north"]
	south, southOK := leases["south"]
	if !northOK || !southOK || s.Count() != 2 || north.Scope == south.Scope {
		t.Fatalf("tenant reservations collided: north=%+v south=%+v count=%d", north, south, s.Count())
	}
	if north.LeaseID == north.ClientKey || south.LeaseID == south.ClientKey || north.LeaseID == south.LeaseID {
		t.Fatalf("lease identity is not owner-scoped: north=%+v south=%+v", north, south)
	}
	if err := s.CommitLease("north", "same", south.LeaseID, "wrong"); err == nil {
		t.Fatal("foreign lease committed")
	}
	current, _ := s.Lookup("north", "same")
	if current.State != "reserved" || current.Response != "" {
		t.Fatalf("failed commit mutated state: %+v", current)
	}
	if err := s.CommitLease("north", "same", north.LeaseID, "ok"); err != nil {
		t.Fatal(err)
	}
	if err := s.AbortLease("south", "same", south.LeaseID); err != nil {
		t.Fatal(err)
	}
	current, _ = s.Lookup("north", "same")
	if current.State != "committed" || current.Response != "ok" || s.Count() != 1 {
		t.Fatalf("commit/abort state machine mismatch: current=%+v count=%d", current, s.Count())
	}
}
