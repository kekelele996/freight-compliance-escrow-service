package freight

import (
	"context"
	"sync"
)

type RoutingWork func(context.Context) (string, error)
type RoutingResultCache struct {
	mu      sync.Mutex
	values  map[string]string
	commits int
}

func NewRoutingResultCache() *RoutingResultCache {
	return &RoutingResultCache{values: map[string]string{}}
}
func (c *RoutingResultCache) Count() int   { c.mu.Lock(); defer c.mu.Unlock(); return len(c.values) }
func (c *RoutingResultCache) Commits() int { c.mu.Lock(); defer c.mu.Unlock(); return c.commits }
