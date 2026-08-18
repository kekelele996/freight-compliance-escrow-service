package freight

import (
	"fmt"
	"sync"
)

type RequestLease struct{ Scope, ClientKey, Fingerprint, LeaseID, Response, State string }
type RequestLeaseStore struct {
	mu       sync.Mutex
	records  map[string]RequestLease
	sequence int
}

func NewRequestLeaseStore() *RequestLeaseStore {
	return &RequestLeaseStore{records: map[string]RequestLease{}}
}
func (s *RequestLeaseStore) Count() int { s.mu.Lock(); defer s.mu.Unlock(); return len(s.records) }
func (s *RequestLeaseStore) Lookup(tenant, key string) (RequestLease, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.records[leaseStorageKey(tenant, key)]
	return v, ok
}
func leaseConflict() error { return fmt.Errorf("%w: lease conflict", ErrDuplicateRequest) }
