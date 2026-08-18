package freight

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
	v := RequestLease{ClientKey: key, Fingerprint: fp, LeaseID: key, State: "reserved"}
	s.records[k] = v
	return v, nil
}
