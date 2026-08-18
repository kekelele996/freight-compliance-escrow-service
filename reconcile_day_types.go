package freight

import "sync"

type ReconcileCheckpoint struct {
	Key, State string
	Attempts   int
}
type ReconcileBook struct {
	mu        sync.Mutex
	values    map[string]ReconcileCheckpoint
	Processed []string
}

func NewReconcileBook() *ReconcileBook {
	return &ReconcileBook{values: map[string]ReconcileCheckpoint{}}
}
func (b *ReconcileBook) Count() int { b.mu.Lock(); defer b.mu.Unlock(); return len(b.Processed) }

// isDone reports whether the operational-day checkpoint for key has already
// been finalized by a successful prior run.
func (b *ReconcileBook) isDone(key string) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.values[key].State == "done"
}
