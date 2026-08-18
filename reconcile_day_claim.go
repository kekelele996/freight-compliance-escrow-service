package freight

func (b *ReconcileBook) claim(key string) (ReconcileCheckpoint, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	v := b.values[key]
	if v.State == "done" || v.State == "running" {
		return v, ErrConflict
	}
	v.Key = key
	v.State = "running"
	v.Attempts++
	b.values[key] = v
	return v, nil
}
