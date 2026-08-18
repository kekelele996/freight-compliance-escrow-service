package freight

import "context"

func (p *LifecyclePool) Run(ctx context.Context, jobs []LifecycleJob) []LifecycleResult {
	out := []LifecycleResult{}
	for _, job := range jobs {
		if !admitLifecycleJob(ctx, job) {
			if ctx != nil && ctx.Err() != nil {
				out = append(out, LifecycleResult{job.ID, ctx.Err(), false})
			}
			continue
		}
		runCtx := lifecycleJobContext(ctx)
		err := job.Run(runCtx)
		if err == nil {
			err = runCtx.Err()
		}
		committed := err == nil && p.commitLifecycle(runCtx, job.ID)
		out = append(out, LifecycleResult{job.ID, err, committed})
		if runCtx.Err() != nil {
			break
		}
	}
	return out
}
