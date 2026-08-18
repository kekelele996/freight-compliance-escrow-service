package freight

func QuoteWithRetry(m *QuoteMemo, r QuoteRequest) (int64, error) {
	if v, ok := m.load(r.Tenant, r.ID); ok {
		return v.Amount, v.Err
	}
	c := buildQuoteCandidate(r)
	m.store(c)
	if c.Err != nil {
		return 0, c.Err
	}
	return c.Amount, nil
}
