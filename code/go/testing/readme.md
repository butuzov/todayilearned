<!-- menu: Testing -->

# Testing with `go test`

## `func TestGo(*testing.T)`

### Commands



```shell
# Coverage for local packages except osme of them...

export PKGS=$(go list ./... | grep -vE "(gotests/gotests|.*data|templates)" | tr -s '\n' ',' | sed 's/.\{1\}$//')
go test -v -covermode=count -coverpkg=$PKGS -coverprofile=coverage.cov

# Find what tests were skipped
go test -v . | grep SKIP

# run tests 10 times + verbose output (also cleans cache)
go test -v -test.count 10 .

# clean cache
go clean -testcache

# run on 2 cores
go test -v -test.count 10 -test.cpu 2 .

# run tests (filter by name)
go test -v -run S .

# run 4 runners tests
go test -v -parallel 4 .

# run 4 runners tests
go test -race .

# json
go test -v --json .

# just compiling test code
go test --exec=/bin/true ./...
go test -c pkg

# Using jq to filter output of json based export.
go test -json | jq -s 'map(select(.Test != null)) | sort_by(.Elapsed)'

# benchmarks (+ memory)
go test -json -benchmem -run=^$ -bench .
```

### Examples

#### T(able)DD

Add Table Tests or (map) for easy tests examples and adding new code.

#### T(ests)DD

TEsts Driven Developments - `first tests then code`

### Linters (`golangci-lint`)

TODO:

- 
-
-
- 

## `func ExampleGo()`

```go
All examples suppose to be located in `_test.go` files. For easier navigation, make it `file_examples_tests.go` for `file.go`
```

```go
//mygopcg.go
package mygopcg

type Spanish struct{}

// Hello in spanish
func (s Spanish) Hello() string {
    return "Hola"
}

// PoliteHello tells polite hello
func (s Spanish) PoliteHello() string {
    return "Buenos Dias"
}

// Greetings return all greetings builtin for Spanish
func (s Spanish) Greetings() []string {
    var greetings []string
    greetings = append(greetings, s.PoliteHello())
    greetings = append(greetings, s.Hello())
    return greetings
}

// Hello function does hello!
func Hello() string {
    return "Hello"
}
```
```go
//mygopcg_exaamples_test.go
package mygopcg

import "fmt"

// Examples can be provided without output, just as example. Such tests wouldn't run.
func ExampleHello() {
    Hello()
}
```

The naming convention to declare examples for the package, a function `F`, a type `T` and method `M` on type `T` are:

- `func Example() { ... }`
- `func ExampleF() { ... }`
- `func ExampleT() { ... }`
- `func ExampleT_M() { ... }`

Multiple example functions for a package/type/function/method may be provided by appending a distinct suffix to the name. The suffix must start with a lower-case letter.

- `func Example_suffix() { ... }`
- `func ExampleF_suffix() { ... }`
- `func ExampleT_suffix() { ... }`
- `func ExampleT_M_suffix() { ... }`

```go
// mygopcg_texamples_test.go
package mygopcg

// In this example you can add `Type` at the end of the example + underscore and method name
// to show how your method works for this type.
func ExampleSpanish_Hello() {
    fmt.Println(Spanish{}.Hello())
    // Output: Hola
}

// By adding custom postfix (starts from underscore and lowercase letter)
// you can specify some costom value in braces
func ExampleSpanish_Hello_exclamations() {
    fmt.Printf("¡%s!\n", Spanish{}.Hello())
    // Output: ¡Hola!
}

// By adding custom postfix (starts from underscore and lowercase letter)
// you can specify some costom value in braces
func ExampleSpanish_PoliteHello_polite() {
    fmt.Println(Spanish{}.PoliteHello())
    // Output: Buenos Dias
}

// Sometimes if we use maps, we wouldn't get same order. So we shuld use
// <code>Unordered output</code> instead Output
func ExampleSpanish_Greetings() {
    for _, s := range (Spanish{}).Greetings() {
    	fmt.Println(s)
    }
    // Unordered output: Halo
    // Buenos Dias
}

```

## `func BenchmarkGo(*testing.B)`

### Reading & Watching & Tooling

* [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
* [`benchstat`: tips or quickstart](https://github.com/golang/go/issues/23471)
* [`benchstat`](https://godoc.org/golang.org/x/perf/cmd/benchstat)
* [Go performance measurement, storage, and analysis tools](https://go.googlesource.com/perf/+/83061fdb061356c976cc90806fa391b09da42568#go-performance-measurement_storage_and-analysis-tools)
* https://github.com/golang/go/wiki/Performance
* [`benchstat` & `pprof` & `benchmarking`](https://www.cloudbees.com/blog/real-life-go-benchmarking)

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

### Examples

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

```shell
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


## `func FuzzGo(*testing.F)`

### Reading

- https://go.dev/doc/tutorial/fuzz
- https://go.dev/doc/security/fuzz/