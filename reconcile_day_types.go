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
