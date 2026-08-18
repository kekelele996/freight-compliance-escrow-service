package freight

func leaseTenantScope(_ string) string          { return "shared" }
func leaseStorageKey(tenant, key string) string { return leaseTenantScope(tenant) + ":" + key }
