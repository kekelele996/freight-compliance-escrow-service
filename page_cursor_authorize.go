package freight

import "fmt"

func authorizeScopedCursor(v ScopedCursor, tenant, filter, direction string) error {
	if v.Tenant != CanonicalTenant(tenant) || v.Filter != filter || v.Direction != direction || v.LastID == "" {
		return fmt.Errorf("%w: cursor scope", ErrValidation)
	}
	return nil
}
