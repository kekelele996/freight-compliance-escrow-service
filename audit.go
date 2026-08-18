package freight

import "sort"

type AuditLog struct{ entries []AuditEntry }

func (a *AuditLog) Add(e AuditEntry) { a.entries = append(a.entries, e) }
func (a *AuditLog) ForTenant(tenant string) []AuditEntry {
	o := []AuditEntry{}
	for _, e := range a.entries {
		if CanonicalTenant(e.Tenant) == CanonicalTenant(tenant) {
			o = append(o, e)
		}
	}
	sort.SliceStable(o, func(i, j int) bool { return o[i].At.Before(o[j].At) })
	return o
}
