package freight

import (
	"errors"
	"sync"
	"testing"
)

func TestApprovalConflictLeavesNoPartialState(t *testing.T) {
	r := NewApprovalRepository()
	r.Conflict = true
	start := make(chan struct{})
	effects := []*ApprovalEffects{{}, {}}
	errs := make(chan error, len(effects))
	var wg sync.WaitGroup
	run := func(effect *ApprovalEffects) {
		defer wg.Done()
		<-start
		errs <- ApproveSettlement(r, effect)
	}
	wg.Add(2)
	go run(effects[0])
	go run(effects[1])
	close(start)
	wg.Wait()
	close(errs)

	for err := range errs {
		if !errors.Is(err, ErrConflict) {
			t.Fatalf("err=%v", err)
		}
	}
	stored := r.Get()
	if stored.State != "pending" || stored.Version != 1 {
		t.Fatalf("repository mutated on conflict: %+v", stored)
	}
	for i, effect := range effects {
		if effect.Checkpoint != 0 || effect.Observed.ID != "" {
			t.Fatalf("checkpoint %d advanced on conflict: %+v", i, effect)
		}
		if len(effect.Events) != 0 {
			t.Fatalf("event %d published on conflict: %v", i, effect.Events)
		}
	}
}
