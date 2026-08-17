package reqmeta

import "context"

// ctxKey is an unexported type so keys of this type cannot collide with
// keys defined in any other package, which is the standard way to avoid
// context value key collisions.
type ctxKey int

const userKey ctxKey = 0

func WithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) string {
	v, _ := ctx.Value(userKey).(string)
	return v
}
