package freight

import (
	"context"
	"errors"
	"testing"
)

func TestWorkerCancellationStopsCommitAndLaterAdmission(t *testing.T) {
	canceled, stop := context.WithCancel(context.Background())
	stop()
	p := &LifecyclePool{}
	if p.commitLifecycle(canceled, "direct") || p.CommitCount() != 0 {
		t.Fatal("direct canceled commit was admitted")
	}

	ctx, cancel := context.WithCancel(context.Background())
	started := make(chan struct{})
	release := make(chan struct{})
	seenCancel := false
	secondRan := false
	done := make(chan []LifecycleResult, 1)
	go func() {
		done <- p.Run(ctx, []LifecycleJob{{ID: "first", Run: func(got context.Context) error {
			close(started)
			<-release
			seenCancel = errors.Is(got.Err(), context.Canceled)
			return nil
		}}, {ID: "second", Run: func(context.Context) error { secondRan = true; return nil }}})
	}()
	<-started
	cancel()
	close(release)
	results := <-done
	if !seenCancel {
		t.Fatal("running job lost caller cancellation")
	}
	if secondRan {
		t.Fatal("queued job started after cancellation")
	}
	if p.CommitCount() != 0 {
		t.Fatalf("canceled job committed: %d", p.CommitCount())
	}
	if len(results) == 0 || !errors.Is(results[0].Err, context.Canceled) || results[0].Committed {
		t.Fatalf("bad cancellation result: %+v", results)
	}
}
