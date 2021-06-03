package testdata

import (
	"fmt"
	"testing"
)

// PROJECT URL:  https://github.com/kunwardeep/paralleltest

func setup_paralleltest(name string) func() {
	fmt.Printf("setup: %s\n", name)
	return func() {
		fmt.Printf("clean %s\n", name)
	}
}

func Test_FunctionMissingCallToParallel_ParallelTest(t *testing.T) {
	teardown := setup_paralleltest("Test_FunctionMissingCallToParallel_ParallleTest")
	defer teardown()
}

func Test_FunctionRangeMissingCallToParallel_ParallelTest(t *testing.T) {
	t.Parallel()

	teardown := setup_paralleltest("Test_FunctionRangeMissingCallToParallel_ParallelTest")
	defer teardown()

	testCases := []struct {
		name string
	}{{name: "foo"}}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
		})
	}
}

func Test_FunctionRangeNotUsingRangeValueInTDotRun_ParallelTest(t *testing.T) {
	t.Parallel()

	teardown := setup_paralleltest("Test_FunctionRangeNotUsingRangeValueInTDotRun_ParallelTest")
	defer teardown()

	testCases := []struct {
		name string
	}{{name: "foo"}}
	for _, tc := range testCases {
		t.Run("this is a test name", func(t *testing.T) {
			// ^ call to tc.name missing
			t.Parallel()
			fmt.Println(tc.name)
		})
	}
}

func Test_FunctionRangeNotReInitialisingVariable_ParallelTest(t *testing.T) {
	t.Parallel()

	teardown := setup_paralleltest("Test_FunctionRangeNotReInitialisingVariable_ParallelTest")
	defer teardown()

	testCases := []struct {
		name string
	}{{name: "foo"}}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fmt.Println(tc.name)
		})
	}
}
