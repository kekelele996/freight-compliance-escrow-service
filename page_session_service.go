package freight

func (s *PageSession) Resume(raw, tenant, filter, direction string) error {
	v, e := decodeScopedCursor(raw)
	if e != nil {
		return e
	}
	if e = authorizeScopedCursor(v, tenant, filter, direction); e != nil {
		return e
	}
	s.commitCheckpoint(s.stageCheckpoint(v))
	return nil
}
