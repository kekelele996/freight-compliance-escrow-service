package freight

import "fmt"

func (s *RequestLeaseStore) ReserveLease(tenant, key, fp string) (RequestLease, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := leaseStorageKey(tenant, key)
	if v, ok := s.records[k]; ok {
		if v.Fingerprint != fp {
			return RequestLease{}, leaseConflict()
		}
		return v, nil
	}
	s.sequence++
	scope := CanonicalTenant(tenant)
	v := RequestLease{
		Scope:       scope,
		ClientKey:   key,
		Fingerprint: fp,
		LeaseID:     fmt.Sprintf("%s#%d", scope, s.sequence),
		State:       "reserved",
	}
	s.records[k] = v
	return v, nil
}
