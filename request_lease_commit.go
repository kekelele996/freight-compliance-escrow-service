package freight

func (s *RequestLeaseStore) CommitLease(tenant, key, leaseID, response string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, v := range s.records {
		if v.ClientKey == key {
			v.State = "committed"
			v.Response = response
			s.records[k] = v
			return nil
		}
	}
	return ErrNotFound
}
