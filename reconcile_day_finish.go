package freight

func (b *ReconcileBook) finish(v ReconcileCheckpoint, runErr error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if runErr != nil {
		v.State = "pending"
		b.values[v.Key] = v
		return
	}
	v.State = "done"
	b.values[v.Key] = v
	b.Processed = append(b.Processed, v.Key)
}
