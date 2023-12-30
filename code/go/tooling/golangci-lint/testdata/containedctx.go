package testdata

import "context"

type ng struct {
	ctx context.Context // want "found a struct that contains a context.Context field"
}
