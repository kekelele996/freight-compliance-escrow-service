package freight

func (s *RequestLeaseStore) CommitLease(tenant, key, leaseID, response string) error {
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
	v.State = "committed"
	v.Response = response
	s.records[k] = v
	return nil
}
