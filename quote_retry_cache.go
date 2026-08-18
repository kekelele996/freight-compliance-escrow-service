package freight

func (m *QuoteMemo) load(t, id string) (*QuoteCandidate, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.values[quoteMemoKey(t, id)]
	if !ok {
		return nil, false
	}
	// Return a copy so callers cannot mutate the cached value.
	cp := *v
	return &cp, true
}
func (m *QuoteMemo) store(c *QuoteCandidate) {
	if c.Err != nil {
		// Never cache a failed candidate; otherwise a later retry with corrected
		// parameters would read back the stale failure.
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	cp := *c
	m.values[quoteMemoKey(c.Tenant, c.ID)] = &cp
}
func (m *QuoteMemo) rollback(t, id string) {}
