package observability

import "context"

type key struct{}

func WithValue(parent context.Context, c Context) context.Context {
	return context.WithValue(parent, key{}, c)
}

func FromContext(ctx context.Context) (Context, bool) {
	c, ok := ctx.Value(key{}).(Context)
	return c, ok
}
