package testdata

import "context"

func fooctxval(ctx context.Context) context.Context {
	return context.WithValue(ctx, "foo", "foo")
}
