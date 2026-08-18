package freight

func buildQuoteCandidate(r QuoteRequest) *QuoteCandidate {
	c := &QuoteCandidate{Tenant: r.Tenant, ID: r.ID, Amount: r.Weight * r.BaseRate}
	if r.Surcharge < 0 {
		c.Err = ErrInvalidAmount
		return c
	}
	c.Amount += r.Surcharge
	c.Complete = true
	return c
}
