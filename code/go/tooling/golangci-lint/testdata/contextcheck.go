package testdata

import "context"

func funcWithCtx(ctx context.Context) {}

func funcWithoutCtx() {
	funcWithCtx(context.TODO())
}

func contextcheckCase1(ctx context.Context) {
	funcWithoutCtx() // want "Function `funcWithoutCtx` should pass the context parameter"
}
