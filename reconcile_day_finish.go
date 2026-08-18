package freight

func (b *ReconcileBook) finish(v ReconcileCheckpoint, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if err == nil {
		b.Processed = append(b.Processed, v.Key)
	}
	b.values[v.Key] = v
}
