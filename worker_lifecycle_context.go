package freight

import "context"

func lifecycleJobContext(parent context.Context) context.Context {
	if parent == nil {
		return context.Background()
	}
	return parent
}
