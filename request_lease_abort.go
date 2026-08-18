package freight

func (s *RequestLeaseStore) AbortLease(tenant, key, leaseID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.records, key)
	return nil
}
