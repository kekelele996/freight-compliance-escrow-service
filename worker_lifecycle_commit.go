package freight

import "context"

func (p *LifecyclePool) commitLifecycle(ctx context.Context, id string) bool {
	if ctx == nil || ctx.Err() != nil {
		return false
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if ctx.Err() != nil {
		return false
	}
	p.commits = append(p.commits, id)
	return true
}
