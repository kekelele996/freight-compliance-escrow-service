package freight

import "sync"

type Metrics struct {
	mu       sync.Mutex
	counters map[string]int64
}

func NewMetrics() *Metrics            { return &Metrics{counters: map[string]int64{}} }
func (m *Metrics) Inc(n string)       { m.mu.Lock(); defer m.mu.Unlock(); m.counters[n]++ }
func (m *Metrics) Get(n string) int64 { m.mu.Lock(); defer m.mu.Unlock(); return m.counters[n] }
func (m *Metrics) Snapshot() map[string]int64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	o := map[string]int64{}
	for k, v := range m.counters {
		o[k] = v
	}
	return o
}
