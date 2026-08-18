package freight

import "time"

// RunDay runs `work` exactly once per (site, operational day), keyed by the
// local calendar date of `at` in `loc`.
//
// A day already finalized ("done") is skipped: `work` is not invoked and the
// day is not re-appended to Processed. A day whose prior attempt failed is
// reclaimed and retried, because its checkpoint was left "in-progress" by
// finish.
func (b *ReconcileBook) RunDay(loc *time.Location, at time.Time, work func(string) error) error {
	key := operationalDayKey(loc, at)
	if b.isDone(key) {
		return ErrConflict
	}
	v := b.claim(key)
	err := work(key)
	b.finish(v, err)
	return err
}
