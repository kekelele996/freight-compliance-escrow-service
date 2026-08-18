package freight

func authorizeScopedCursor(v ScopedCursor, tenant, filter, direction string) error {
	if CanonicalTenant(v.Tenant) != CanonicalTenant(tenant) {
		return ErrValidation
	}
	if v.Filter != filter {
		return ErrValidation
	}
	if v.Direction != direction {
		return ErrValidation
	}
	return nil
}
