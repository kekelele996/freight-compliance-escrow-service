package freight

func (s *PageSession) Resume(raw, tenant, filter, direction string) error {
	v, e := decodeScopedCursor(raw)
	if e != nil {
		return e
	}
	previous := s.Checkpoint
	s.stageCheckpoint(v)
	if e = authorizeScopedCursor(v, tenant, filter, direction); e != nil {
		s.rollbackCheckpoint(previous)
		return e
	}
	return nil
}
