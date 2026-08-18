package freight

import "context"

func (c *RoutingResultCache) put(ctx context.Context, key, value string) bool {
	if ctx == nil || ctx.Err() != nil {
		return false
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if ctx.Err() != nil {
		return false
	}
	c.values[key] = value
	return true
}
