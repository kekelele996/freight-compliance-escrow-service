package freight

import "time"

func (b *ReconcileBook) RunDay(loc *time.Location, at time.Time, work func(string) error) error {
	key := operationalDayKey(loc, at)
	v, err := b.claim(key)
	if err != nil {
		return err
	}
	runErr := work(key)
	b.finish(v, runErr)
	return runErr
}
