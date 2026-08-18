package freight

import "sync"

type DeliveryRecord struct {
	ID, State string
	Attempts  int
	LastError string
}
type DeliveryQueue struct {
	mu      sync.Mutex
	records map[string]DeliveryRecord
}

func NewDeliveryQueue() *DeliveryQueue { return &DeliveryQueue{records: map[string]DeliveryRecord{}} }
func (q *DeliveryQueue) Add(id string) { q.records[id] = DeliveryRecord{ID: id, State: "pending"} }
func (q *DeliveryQueue) Get(id string) DeliveryRecord {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.records[id]
}
