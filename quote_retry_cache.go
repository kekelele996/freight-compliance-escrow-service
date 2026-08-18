package freight

func cloneQuoteCandidate(v *QuoteCandidate) *QuoteCandidate {
	if v == nil {
		return nil
	}
	c := *v
	return &c
}
func (m *QuoteMemo) load(t, id string) (*QuoteCandidate, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.values[quoteMemoKey(t, id)]
	return cloneQuoteCandidate(v), ok
}
func (m *QuoteMemo) store(c *QuoteCandidate) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if c == nil || c.Err != nil || !c.Complete {
		return
	}
	m.values[quoteMemoKey(c.Tenant, c.ID)] = cloneQuoteCandidate(c)
}
func (m *QuoteMemo) rollback(t, id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.values, quoteMemoKey(t, id))
}
