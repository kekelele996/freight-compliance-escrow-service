package freight

import "context"

func RunRouting(ctx context.Context, key string, cache *RoutingResultCache, work RoutingWork) (string, error) {
	runCtx := routingContext(ctx)
	if err := runCtx.Err(); err != nil {
		return "", err
	}
	value, err := executeRoutingWork(runCtx, work)
	if err != nil {
		return "", err
	}
	if !cache.put(runCtx, key, value) {
		return "", runCtx.Err()
	}
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if err = runCtx.Err(); err != nil {
		delete(cache.values, key)
		return "", err
	}
	cache.commits++
	return value, nil
}
