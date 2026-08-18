package freight

import "time"

func (b *ReconcileBook) RunDay(loc *time.Location, at time.Time, work func(string) error) error {
	key := operationalDayKey(loc, at)
	v := b.claim(key)
	err := work(key)
	b.finish(v, err)
	return err
}
