package freight

func (m *QuoteMemo) load(t, id string) (*QuoteCandidate, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.values[quoteMemoKey(t, id)]
	return v, ok
}
func (m *QuoteMemo) store(c *QuoteCandidate) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.values[quoteMemoKey(c.Tenant, c.ID)] = c
}
func (m *QuoteMemo) rollback(t, id string) {}
