package freight

func QuoteWithRetry(m *QuoteMemo, r QuoteRequest) (int64, error) {
	if v, ok := m.load(r.Tenant, r.ID); ok && v.Complete && v.Err == nil {
		return v.Amount, nil
	}
	c := buildQuoteCandidate(r)
	if c.Err != nil {
		m.rollback(r.Tenant, r.ID)
		return 0, c.Err
	}
	m.store(c)
	return c.Amount, nil
}
