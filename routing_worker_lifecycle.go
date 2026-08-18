package freight

import "context"

func executeRoutingWork(ctx context.Context, work RoutingWork) (string, error) { return work(ctx) }
