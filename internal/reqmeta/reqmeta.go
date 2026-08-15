
package reqmeta

import "context"

type key int

const userKey key = 1

func WithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) string {
	v, _ := ctx.Value(userKey).(string)
	return v
}
