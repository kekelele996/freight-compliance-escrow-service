package freight

func (q *DeliveryQueue) finish(id string, sendErr error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	v := q.records[id]
	if sendErr != nil {
		v.LastError = sendErr.Error()
		v.State = "pending"
	} else {
		v.LastError = ""
		v.State = "published"
	}
	q.records[id] = v
}
