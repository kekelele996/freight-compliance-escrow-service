package freight

// claim reserves the operational-day checkpoint for a retry attempt.
//
// The checkpoint is marked "in-progress", never "done", so a task that later
// fails remains unfinished and a subsequent retry of the same day picks it up
// rather than being treated as a fresh batch.
func (b *ReconcileBook) claim(key string) ReconcileCheckpoint {
	b.mu.Lock()
	defer b.mu.Unlock()
	v := b.values[key]
	v.Key = key
	v.State = "in-progress"
	v.Attempts++
	b.values[key] = v
	return v
}
