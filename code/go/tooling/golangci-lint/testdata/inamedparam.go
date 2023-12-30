package testdata

import "context"

type NamedParam interface {
	Void()

	NoArgs() string

	WithName(ctx context.Context, number int, toggle bool, tStruct *struct{ b bool }) (bool, error)

	WithoutName(
		context.Context, // want "interface method WithoutName must have named param for type context.Context"
		int, // want "interface method WithoutName must have named param for type int"
		bool, // want "interface method WithoutName must have named param for type bool"
		struct{ b bool }, // want "interface method WithoutName must have all named params"
	)
}
