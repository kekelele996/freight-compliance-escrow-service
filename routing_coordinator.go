package freight

import "context"

func RunRouting(ctx context.Context, key string, cache *RoutingResultCache, work RoutingWork) (string, error) {
	runCtx := routingContext(ctx)
	value, err := executeRoutingWork(runCtx, work)
	if err != nil {
		return "", err
	}
	cache.put(runCtx, key, value)
	cache.mu.Lock()
	cache.commits++
	cache.mu.Unlock()
	return value, nil
}
