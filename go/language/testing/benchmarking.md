# Benchmarking (Go)

### Reading list

* https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

### Arguments

* __-bench__ `regexp` - run only benchmarks matching regexp
* __-benchmem__ - print memory allocations for benchmarks
* __-benchtime__ `d` - run each benchmark for duration d (default `1s`)
* __-blockprofile__ `file` - write a goroutine blocking profile to file
* __-blockprofilerate__ `rate` - set blocking profile rate (see `runtime.SetBlockProfileRate`) (default `1`)
* __-count__ `n` - run tests and benchmarks n times (default 1)
* __-coverprofile__ `file` - write a coverage profile to file
* __-cpu__ `list` - comma-separated list of cpu counts to run each test with
* __-cpuprofile__ `file` - write a cpu profile to file
* __-failfast__ - do not start new tests after the first test failure
* __-list__ `regexp` - list tests, examples, and benchmarks matching regexp then exit
* __-memprofile__ `file` - write an allocation profile to file
* __-memprofilerate__ `rate` - set memory allocation profiling rate (see runtime.MemProfileRate)
* __-mutexprofile__ `string` - write a mutex contention profile to the named file after execution
* __-mutexprofilefraction__ `int` - if >= 0, calls runtime.SetMutexProfileFraction() (default `1`)
* __-outputdir__ `dir` - write profiles to dir
* __-parallel__ `n` - run at most n tests in parallel (default `24`)
* __-run__ `regexp` - run only tests and examples matching regexp
* __-short__ - 	run smaller test suite to save time
* __-testlogfile__ `file` - 	write test action log to file (for use only by `cmd/go`)
* __-timeout__ `d` - panic test binary after duration d (default `0`, timeout disabled)
* __-trace__ `file` - write an execution trace to file
* __-v__ - verbose: print additional output

#### A note on compiler optimisations

Before concluding I wanted to highlight that to be completely accurate, any benchmark should be careful to avoid compiler optimisations eliminating the function under test and artificially lowering the run time of the benchmark.m

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
```

#### Generating Stat for benchmarking
 

* https://github.com/golang/go/issues/23471
* https://godoc.org/golang.org/x/perf/cmd/benchstat
* [Go performance measurement, storage, and analysis tools](https://go.googlesource.com/perf/+/83061fdb061356c976cc90806fa391b09da42568#go-performance-measurement_storage_and-analysis-tools)
* https://github.com/golang/go/wiki/Performance
* https://blog.codeship.com/real-life-go-benchmarking/

```bash
# install benchstat                                                           
go get -u golang.org/x/perf/cmd/benchstat
# running go test with next arguments will run BenchmarkName 10 
# times in verbose mode. output will be redirected to results.txt
go test ./...  -bench=BenchmarkName -benchmem -run=^$ -count=10 -cpu=1,4  > results.txt

# alternativly we can pass regular expression to run all benchmarks 
go test ./...  -bench='.*' -benchmem -run=^$ -count=10 -cpu=1,4  > results.txt

# analyze output
benchmark results.txt

# more deailed benchmarking with 
# https://github.com/alecthomas/go_serialization_benchmarks/blob/master/stats.sh
```

### Timers

* https://medium.com/justforfunc/analyzing-the-performance-of-go-functions-with-benchmarks-60b8162e61c6

```go
func bench(b *testing.B, fnc func(...) ...) {
    
    // b.ResetTimer() // will reset time and memory allocation.
                      // doesn't work if timer is running.
    
	b.StopTimer()     // <- Stoping timing test
    heavyLifting()    // some setup heavylifting  
	b.StartTimer()    // <- Starting timing test
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

### `b.Run`

Sub-Benchmark

```go
// we can run subbenchmarks in same way as we running subtests...
b.Run(Name, func(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Our testing fucntionality...
    }
})
```

### Useful to have

* `b.ReportAllocs()` - malocs reports