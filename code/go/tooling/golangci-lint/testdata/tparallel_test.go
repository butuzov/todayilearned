package testdata

import (
	"fmt"
	"testing"
)

// PROJECT URL: https://github.com/moricho/tparallel

func setup_tparallel(name string) func() {
	fmt.Printf("setup: %s\n", name)
	return func() {
		fmt.Printf("clean %s\n", name)
	}
}

func Test_FunctionMissingCallToParallel_Tparallel(t *testing.T) {
	teardown := setup_tparallel("Test_FunctionMissingCallToParallel_ParallleTest")
	defer teardown()
}

func Test_FunctionRangeMissingCallToParallel_Tparallel(t *testing.T) {
	t.Parallel()

	teardown := setup_tparallel("Test_FunctionRangeMissingCallToParallel_Tparallel")
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

func Test_FunctionRangeNotUsingRangeValueInTDotRun_Tparallel(t *testing.T) {
	t.Parallel()

	teardown := setup_tparallel("Test_FunctionRangeNotUsingRangeValueInTDotRun_Tparallel")
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

func Test_FunctionRangeNotReInitialisingVariable_Tparallel(t *testing.T) {
	t.Parallel()

	teardown := setup_tparallel("Test_FunctionRangeNotReInitialisingVariable_Tparallel")
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
