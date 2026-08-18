package freight

func leaseTenantScope(tenant string) string          { return CanonicalTenant(tenant) }
func leaseStorageKey(tenant, key string) string { return leaseTenantScope(tenant) + ":" + key }
