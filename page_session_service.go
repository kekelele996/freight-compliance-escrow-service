package freight

func (s *PageSession) Resume(raw, tenant, filter, direction string) error {
	v, e := decodeScopedCursor(raw)
	if e != nil {
		return e
	}
	if e = authorizeScopedCursor(v, tenant, filter, direction); e != nil {
		return e
	}
	previous := s.Checkpoint
	staged := s.stageCheckpoint(v)
	if staged.Version != previous.Version+1 {
		s.rollbackCheckpoint(previous)
		return ErrConflict
	}
	s.commitCheckpoint(staged)
	return nil
}
