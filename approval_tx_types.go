package freight

import "sync"

type ApprovalRecord struct {
	ID, State string
	Version   int
}
type ApprovalRepository struct {
	mu       sync.Mutex
	record   ApprovalRecord
	Conflict bool
}
type ApprovalEffects struct {
	Checkpoint int
	Events     []string
	Observed   ApprovalRecord
}

func NewApprovalRepository() *ApprovalRepository {
	return &ApprovalRepository{record: ApprovalRecord{ID: "s1", State: "pending", Version: 1}}
}
func (r *ApprovalRepository) Get() ApprovalRecord { r.mu.Lock(); defer r.mu.Unlock(); return r.record }
