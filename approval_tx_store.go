package freight

func (r *ApprovalRepository) save(v ApprovalRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.Conflict {
		return ErrConflict
	}
	r.record = v
	return nil
}
