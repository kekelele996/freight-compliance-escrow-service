package freight

import "context"

func (p *LifecyclePool) Run(ctx context.Context, jobs []LifecycleJob) []LifecycleResult {
	out := []LifecycleResult{}
	for _, job := range jobs {
		if !admitLifecycleJob(ctx, job) {
			continue
		}
		runCtx := lifecycleJobContext(ctx)
		err := job.Run(runCtx)
		committed := p.commitLifecycle(runCtx, job.ID)
		out = append(out, LifecycleResult{job.ID, err, committed})
	}
	return out
}
