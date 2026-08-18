package freight

func (r *ApprovalRepository) save(v ApprovalRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.Conflict {
		return ErrConflict
	}
	if v.Version != r.record.Version+1 {
		return ErrConflict
	}
	r.record = v
	return nil
}
