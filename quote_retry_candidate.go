package freight

func buildQuoteCandidate(r QuoteRequest) *QuoteCandidate {
	c := &QuoteCandidate{Tenant: CanonicalTenant(r.Tenant), ID: r.ID}
	if r.Weight < 0 || r.BaseRate < 0 || r.Surcharge < 0 {
		c.Err = ErrInvalidAmount
		return c
	}
	c.Amount = r.Weight*r.BaseRate + r.Surcharge
	c.Complete = true
	return c
}
