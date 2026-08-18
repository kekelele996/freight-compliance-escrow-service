package freight

import "sync"

type ShipmentStore struct {
	mu     sync.RWMutex
	values map[string]Shipment
}

func NewShipmentStore() *ShipmentStore { return &ShipmentStore{values: map[string]Shipment{}} }
func (s *ShipmentStore) Put(v Shipment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[v.ID] = cloneShipment(v)
}
func (s *ShipmentStore) Get(id string) (Shipment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.values[id]
	if !ok {
		return Shipment{}, ErrNotFound
	}
	return cloneShipment(v), nil
}
func (s *ShipmentStore) List() []Shipment {
	s.mu.RLock()
	defer s.mu.RUnlock()
	o := make([]Shipment, 0, len(s.values))
	for _, v := range s.values {
		o = append(o, cloneShipment(v))
	}
	return o
}
func cloneShipment(v Shipment) Shipment {
	v.Segments = append([]Segment(nil), v.Segments...)
	tags := map[string]string{}
	for k, x := range v.Tags {
		tags[k] = x
	}
	v.Tags = tags
	return v
}
