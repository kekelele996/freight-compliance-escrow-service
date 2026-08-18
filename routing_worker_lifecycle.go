package freight

import "context"

func executeRoutingWork(ctx context.Context, work RoutingWork) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	value, err := work(ctx)
	if err != nil {
		return "", err
	}
	if err = ctx.Err(); err != nil {
		return "", err
	}
	return value, nil
}
