package freight

import (
	"context"
	"errors"
	"testing"
	"time"
)

type errOnCallContext struct{ after, calls int }

func (c *errOnCallContext) Deadline() (time.Time, bool) { return time.Time{}, false }
func (c *errOnCallContext) Done() <-chan struct{}       { return nil }
func (c *errOnCallContext) Err() error {
	c.calls++
	if c.calls >= c.after {
		return context.Canceled
	}
	return nil
}
func (c *errOnCallContext) Value(key any) any { return nil }

func TestRoutingCancellationBoundaries(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	cancel()
	if !errors.Is(routingContext(parent).Err(), context.Canceled) {
		t.Fatal("request cancellation was not preserved")
	}

	running, stop := context.WithCancel(context.Background())
	_, err := executeRoutingWork(running, func(context.Context) (string, error) { stop(); return "late", nil })
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("late worker result=%v", err)
	}

	cache := NewRoutingResultCache()
	stopped, end := context.WithCancel(context.Background())
	end()
	if cache.put(stopped, "direct", "stale") || cache.Count() != 0 {
		t.Fatal("stopped request populated the cache")
	}

	edge := &errOnCallContext{after: 6}
	value, err := RunRouting(edge, "edge", cache, func(context.Context) (string, error) { return "late", nil })
	if !errors.Is(err, context.Canceled) || value != "" || cache.Count() != 0 || cache.Commits() != 0 {
		t.Fatalf("commit boundary published value=%q err=%v cached=%d commits=%d", value, err, cache.Count(), cache.Commits())
	}
}
