package freight

func (q *DeliveryQueue) PublishOne(id string, send func(string) error) error {
	v, err := q.claim(id)
	if err != nil {
		return err
	}
	if err = invokeDeliverySend(v, send); err != nil {
		return err
	}
	q.finish(id, nil)
	return nil
}
func (q *DeliveryQueue) PendingIDs() []string {
	q.mu.Lock()
	defer q.mu.Unlock()
	out := []string{}
	for id, v := range q.records {
		if v.State == "pending" {
			out = append(out, id)
		}
	}
	return out
}
