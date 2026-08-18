package freight

import "context"

func (p *LifecyclePool) commitLifecycle(_ context.Context, id string) bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.commits = append(p.commits, id)
	return true
}
