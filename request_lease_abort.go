package freight

func (s *RequestLeaseStore) AbortLease(tenant, key, leaseID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := leaseStorageKey(tenant, key)
	v, ok := s.records[k]
	if !ok {
		return ErrNotFound
	}
	if v.LeaseID != leaseID || v.State != "reserved" {
		return leaseConflict()
	}
	delete(s.records, k)
	return nil
}
