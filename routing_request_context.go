package freight

import "context"

func routingContext(parent context.Context) context.Context {
	if parent == nil {
		return context.Background()
	}
	return parent
}
