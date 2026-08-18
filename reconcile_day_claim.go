package freight

func (b *ReconcileBook) claim(key string) ReconcileCheckpoint {
	b.mu.Lock()
	defer b.mu.Unlock()
	v := b.values[key]
	v.Key = key
	v.State = "done"
	v.Attempts++
	b.values[key] = v
	return v
}
