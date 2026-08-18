package freight

import (
	"context"
	"sync"
)

type LifecycleJob struct {
	ID  string
	Run func(context.Context) error
}
type LifecycleResult struct {
	ID        string
	Err       error
	Committed bool
}
type LifecyclePool struct {
	mu      sync.Mutex
	commits []string
}

func (p *LifecyclePool) CommitCount() int { p.mu.Lock(); defer p.mu.Unlock(); return len(p.commits) }
