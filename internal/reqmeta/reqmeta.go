
package reqmeta

import "context"

func WithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, "user", user) // BUG: string key
}

func User(ctx context.Context) string {
	v, _ := ctx.Value("user").(string)
	return v
}
