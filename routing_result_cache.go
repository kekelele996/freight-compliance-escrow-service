package freight

import "context"

func (c *RoutingResultCache) put(_ context.Context, key, value string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
	return true
}
