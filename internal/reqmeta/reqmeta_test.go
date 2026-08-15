
package reqmeta

import (
	"context"
	"testing"
)

func TestRoundTripValue(t *testing.T) {
	ctx := WithUser(context.Background(), "alice")
	// collide with foreign string key
	ctx = context.WithValue(ctx, "user", "bob")
	if User(ctx) != "alice" {
		t.Fatalf("got %q", User(ctx))
	}
}
