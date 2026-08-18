package freight

import "context"

func admitLifecycleJob(ctx context.Context, job LifecycleJob) bool {
	return ctx != nil && ctx.Err() == nil && job.ID != "" && job.Run != nil
}
