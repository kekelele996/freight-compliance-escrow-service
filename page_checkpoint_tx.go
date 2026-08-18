package freight

func (s *PageSession) stageCheckpoint(v ScopedCursor) PageCheckpoint {
	return PageCheckpoint{v.Tenant, v.Filter, v.Direction, v.LastID, s.Checkpoint.Version + 1}
}
func (s *PageSession) commitCheckpoint(v PageCheckpoint)          { s.Checkpoint = v }
func (s *PageSession) rollbackCheckpoint(previous PageCheckpoint) { s.Checkpoint = previous }
