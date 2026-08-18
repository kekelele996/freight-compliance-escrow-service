package freight

func quoteMemoKey(tenant, id string) string { return CanonicalTenant(tenant) + ":" + id }
