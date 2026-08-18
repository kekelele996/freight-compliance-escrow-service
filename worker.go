package freight

import (
	"context"
	"sync"
)

type Job struct {
	ID, TenantID string
	Run          func(context.Context) error
}
type JobResult struct {
	ID  string
	Err error
}
type WorkerPool struct{ Parallelism int }

func (p WorkerPool) Run(ctx context.Context, jobs []Job) []JobResult {
	n := p.Parallelism
	if n < 1 {
		n = 1
	}
	in := make(chan Job)
	out := make(chan JobResult, len(jobs))
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range in {
				if ctx.Err() != nil {
					out <- JobResult{job.ID, ctx.Err()}
					continue
				}
				out <- JobResult{job.ID, job.Run(ctx)}
			}
		}()
	}
	go func() {
		defer close(in)
		for _, j := range jobs {
			select {
			case <-ctx.Done():
				return
			case in <- j:
			}
		}
	}()
	go func() { wg.Wait(); close(out) }()
	results := make([]JobResult, 0, len(jobs))
	for r := range out {
		results = append(results, r)
	}
	return results
}
