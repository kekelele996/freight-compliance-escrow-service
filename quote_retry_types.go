package freight

import "sync"

type QuoteRequest struct {
	Tenant, ID                  string
	Weight, BaseRate, Surcharge int64
}
type QuoteCandidate struct {
	Tenant, ID string
	Amount     int64
	Complete   bool
	Err        error
}
type QuoteMemo struct {
	mu     sync.Mutex
	values map[string]*QuoteCandidate
}

func NewQuoteMemo() *QuoteMemo  { return &QuoteMemo{values: map[string]*QuoteCandidate{}} }
func (m *QuoteMemo) Count() int { m.mu.Lock(); defer m.mu.Unlock(); return len(m.values) }
