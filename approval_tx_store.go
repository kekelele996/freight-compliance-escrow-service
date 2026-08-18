package freight

func (r *ApprovalRepository) save(v ApprovalRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.record = v
	if r.Conflict {
		return ErrConflict
	}
	return nil
}
