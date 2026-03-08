# `*testing.B`

## Reading & Watching & Tooling

- [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [`benchstat`: tips or quickstart](https://github.com/golang/go/issues/23471)
- [`benchstat`](https://godoc.org/golang.org/x/perf/cmd/benchstat)
- [Go performance measurement, storage, and analysis tools](https://go.googlesource.com/perf/+/83061fdb061356c976cc90806fa391b09da42568#go-performance-measurement_storage_and-analysis-tools)
- https://github.com/golang/go/wiki/Performance
- [`benchstat` & `pprof` & `benchmarking`](https://www.cloudbees.com/blog/real-life-go-benchmarking)

## Arguments

- **-bench** `regexp` - run only benchmarks matching regexp
- **-benchmem** - print memory allocations for benchmarks
- **-benchtime** `d` - run each benchmark for duration d (default `1s`)
- **-blockprofile** `file` - write a goroutine blocking profile to file
- **-blockprofilerate** `rate` - set blocking profile rate (see `runtime.SetBlockProfileRate`) (default `1`)
- **-count** `n` - run tests and benchmarks n times (default 1)
- **-coverprofile** `file` - write a coverage profile to file
- **-cpu** `list` - comma-separated list of cpu counts to run each test with
- **-cpuprofile** `file` - write a cpu profile to file
- **-failfast** - do not start new tests after the first test failure
- **-list** `regexp` - list tests, examples, and benchmarks matching regexp then exit
- **-memprofile** `file` - write an allocation profile to file
- **-memprofilerate** `rate` - set memory allocation profiling rate (see runtime.MemProfileRate)
- **-mutexprofile** `string` - write a mutex contention profile to the named file after execution
- **-mutexprofilefraction** `int` - if >= 0, calls runtime.SetMutexProfileFraction() (default `1`)
- **-outputdir** `dir` - write profiles to dir
- **-parallel** `n` - run at most n tests in parallel (default `24`)
- **-run** `regexp` - run only tests and examples matching regexp
- **-short** - run smaller test suite to save time
- **-testlogfile** `file` - write test action log to file (for use only by `cmd/go`)
- **-timeout** `d` - panic test binary after duration d (default `0`, timeout disabled)
- **-trace** `file` - write an execution trace to file
- **-v** - verbose: print additional output

## Examples

> Before concluding I wanted to highlight that to be completely accurate, any benchmark should be careful to avoid compiler optimisations eliminating the function under test and artificially lowering the run time of the benchmark.m

```go
var result int

func BenchmarkFibComplete(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		r = Fib(10)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}


// Skipping (not mesuring) time of all not requied things (create reso etc)
func bench(b *testing.B, fnc func(...) ...) {
    // b.ResetTimer() // will reset time and memory allocation.
                      // doesn't work if timer is running.

	b.StopTimer()     // <- Stoping timing test
	heavyLifting()    // some setup heavylifting
	b.StartTimer()    // <- Starting timing test

	// now we bench

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r int
			for _, test := range TestCases {
				r = fnc(...)
			}
			result = r
		}
	})
}
```

### benchstat

```shell
go get -u golang.org/x/perf/cmd/benchstat
go test ./...  -bench=BenchmarkName -benchmem -run=^$ -count=10 -cpu=1,4  > results.txt
go test ./...  -bench='.*' -benchmem -run=^$ -count=10 -cpu=1,4  > results.txt

benchmark results.txt
```

### Other

#### `b.Run`

Sub-Benchmark

```golang
func Benchmark_...(b *testing.B) {
    ....
    // we can run subbenchmarks in same way as we running subtests...
    b.Run(Name, func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            // Our testing fucntionality...
        }
    })
}
```

#### Allocations

```golang
var result int
// Benchmark_...	 2000000	       771 ns/op	     208 B/op	       7 allocs/op => Go 1.4.2
// Benchmark_...  	 3000000	       404 ns/op	     160 B/op	       3 allocs/op => Go 1.5.0
func Benchmark_MoneyNewGetf(b *testing.B) {
    va r int
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c := money.New(money.WithPrecision(100)).Setf(-123456.789)
		r = c.Getf()
	}
    result = r
}
```
