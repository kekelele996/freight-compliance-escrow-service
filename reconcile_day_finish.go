package freight

// finish records the outcome of a retry attempt against its checkpoint.
//
// Only a successful attempt finalizes the day ("done" and appended to
// Processed). A failed attempt leaves the checkpoint in-progress so the next
// retry of the same operational day reclaims it.
func (b *ReconcileBook) finish(v ReconcileCheckpoint, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if err == nil {
		v.State = "done"
		b.Processed = append(b.Processed, v.Key)
	}
	b.values[v.Key] = v
}
