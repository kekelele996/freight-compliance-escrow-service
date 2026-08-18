package freight

import (
	"sync"
	"time"
)

type IdempotencyRecord struct {
	Fingerprint, Response string
	ExpiresAt             time.Time
}
type IdempotencyStore struct {
	mu      sync.Mutex
	records map[string]IdempotencyRecord
	clock   Clock
}

func NewIdempotencyStore(c Clock) *IdempotencyStore {
	return &IdempotencyStore{records: map[string]IdempotencyRecord{}, clock: c}
}
func (s *IdempotencyStore) key(tenant, key string) string { return CanonicalTenant(tenant) + ":" + key }
func (s *IdempotencyStore) Reserve(tenant, key, fp string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := s.key(tenant, key)
	now := s.clock.Now()
	if r, ok := s.records[k]; ok && r.ExpiresAt.After(now) {
		if r.Fingerprint != fp {
			return "", ErrDuplicateRequest
		}
		return r.Response, nil
	}
	s.records[k] = IdempotencyRecord{Fingerprint: fp, ExpiresAt: now.Add(24 * time.Hour)}
	return "", nil
}
func (s *IdempotencyStore) Commit(tenant, key, response string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	k := s.key(tenant, key)
	r := s.records[k]
	r.Response = response
	s.records[k] = r
}
func (s *IdempotencyStore) Sweep() {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := s.clock.Now()
	for k, r := range s.records {
		if !r.ExpiresAt.After(now) {
			delete(s.records, k)
		}
	}
}
