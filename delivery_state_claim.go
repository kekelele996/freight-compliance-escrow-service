package freight

func (q *DeliveryQueue) claim(id string) (DeliveryRecord, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	v, ok := q.records[id]
	if !ok {
		return DeliveryRecord{}, ErrNotFound
	}
	if v.State != "pending" {
		return DeliveryRecord{}, ErrConflict
	}
	v.State = "inflight"
	v.Attempts++
	q.records[id] = v
	return v, nil
}
