package freight

import "sync"

type Outbox struct {
	mu     sync.Mutex
	events []OutboxEvent
}

func (o *Outbox) Enqueue(e OutboxEvent) {
	o.mu.Lock()
	defer o.mu.Unlock()
	e.Sequence = len(o.events) + 1
	o.events = append(o.events, e)
}
func (o *Outbox) Pending(tenant string) []OutboxEvent {
	o.mu.Lock()
	defer o.mu.Unlock()
	r := []OutboxEvent{}
	for _, e := range o.events {
		if e.TenantID == tenant && !e.Published {
			r = append(r, e)
		}
	}
	return r
}
func (o *Outbox) MarkPublished(id string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	for i := range o.events {
		if o.events[i].ID == id {
			o.events[i].Published = true
			return
		}
	}
}
func (o *Outbox) Publish(tenant string, send func(OutboxEvent) error) error {
	for _, e := range o.Pending(tenant) {
		if e := send(e); e != nil {
			return e
		}
		o.MarkPublished(e.ID)
	}
	return nil
}
