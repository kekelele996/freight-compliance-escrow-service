package freight

// stageCheckpoint computes the next checkpoint without mutating session state,
// so a cursor that fails authorization can be discarded without rollback.
func (s *PageSession) stageCheckpoint(v ScopedCursor) PageCheckpoint {
	return PageCheckpoint{
		Tenant:    CanonicalTenant(v.Tenant),
		Filter:    v.Filter,
		Direction: v.Direction,
		LastID:    v.LastID,
		Version:   s.Checkpoint.Version + 1,
	}
}

// commitCheckpoint atomically applies a staged checkpoint to the session.
func (s *PageSession) commitCheckpoint(c PageCheckpoint) { s.Checkpoint = c }

func (s *PageSession) rollbackCheckpoint(_ PageCheckpoint) {}
