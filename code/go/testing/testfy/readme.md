# github.com/stretchr/testify

```go
$go get github.com/butuzov/errorf@v0.1.0

> go: go.mod file not found in current directory or any parent directory.
> 'go get' is no longer supported outside a module.
> To build and install a command, use 'go install' with a version,
> like 'go install example.com/cmd@latest'
> For more information, see https://golang.org/doc/go-get-install-deprecation
> or run 'go help get' or 'go help install'.
```

```go
import "github.com/stretchr/testify/assert"
import "testing"
// import "github.com/butuzov/errorf"
```

```go
// testing.T replacement
t := &testing.T{}
```

# `assert` of the `testify`

## Formatted Printing

It doesn't matter which you gonna use, but with formated you can have nice messafe if you need to.

## Helpers

* `FailNow` - fail
* `Fail` - fail
* `CallerInfo` - stack trace

```go
assert.Fail(t, "Failing Fast and Furious")
```

```go
assert.FailNow(t, "Failing Fast and Furious")

> Failing Fast and Furious
```

```go
import "fmt"

fmt.Sprintf("%#v", assert.CallerInfo())

result >>> []string{"value.go:476", "value.go:337", "value.go:127", "call.go:453", "callnret1.go:246", "literal.go:296", "call_variadic.go:423", "util.go:999", "repl.go:132", "kernel.go:587", "kernel.go:492", "kernel.go:346", "kernel.go:226", "main.go:26", "proc.go:225", "asm_amd64.s:1371"}
```

```go
t := &errorf.T{}
var i int

assert.Equal(t, map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 1, 3: 3, 2: 2})

result >>> true
```

## available `assert`'s 

### Bools

```go
_ = assert.True(t, false)

> Should be true
```

```go
_ = assert.Falsef(t, true, "wrong")

> Should be false
> Messages:   	wrong
```

### Zeros and Empty

```go
var value int
_ = assert.NotZero(t, value)

> Should not be zero, but was 0
```

```go
var value = 1
_ = assert.Zero(t, value)

> Should be zero, but was 1
```

```go
// emptines (zero value or len=0) of value
var value bool
_ = assert.NotEmpty(t, value)

> Should NOT be empty, but was false
```

```go
var value = []int{1}
_ = assert.Empty(t, value)

> Should be empty, but was [1]
```

### `nil` values

```go
_ = assert.Nil(t, 1)

> Expected nil, but got: 1
```

```go
_ = assert.NotNil(t, nil)

> Expected value not to be nil.
```

### Comparing Values

```go
_ = assert.Greater(t, 1, 2)

> "1" is not greater than "2"
> Messages:   	[]
```

```go
_ = assert.GreaterOrEqual(t, 1, 2)

> "1" is not greater than or equal to "2"
> Messages:   	[]
```

```go
_ = assert.LessOrEqual(t, 2, 1)

> "2" is not less than or equal to "1"
> Messages:   	[]
```

```go
_ = assert.Less(t, 2, 1)

> "2" is not less than "1"
> Messages:   	[]
```

```go
// type matters =)
_ = assert.EqualValues(t, int32(100), int64(100))
```

```go
// not as same as equal
_ = assert.Exactly(t, int32(123), int64(123))

> Types expected to match exactly
> int32 != int64
```

```go
_ = assert.Equal(t, int32(100), int64(100))

> Not equal:
> expected: int32(100)
> actual  : int64(100)
```

```go
// can be checked with containers
assert.Equal(t, []int{1, 2, 4}, []int{1, 2, 3})

> Not equal:
> expected: []int{1, 2, 4}
> actual  : []int{1, 2, 3}
> 
> Diff:
> --- Expected
> +++ Actual
> @@ -3,3 +3,3 @@
> (int) 2,
> - (int) 4
> + (int) 3
> }

result >>> false
```

```go
_ = assert.Equal(t, map[int]int{1: 1, 2: 2, 3: 1}, map[int]int{1: 1, 3: 3, 2: 2})

> Not equal:
> expected: map[int]int{1:1, 2:2, 3:1}
> actual  : map[int]int{1:1, 2:2, 3:3}
> 
> Diff:
> --- Expected
> +++ Actual
> @@ -3,3 +3,3 @@
> (int) 2: (int) 2,
> - (int) 3: (int) 1
> + (int) 3: (int) 3
> }
```

```go
_ = assert.NotEqual(t, 100, 100)

> Should not be: 100
```

### Containers

```go
_ = assert.ElementsMatch(t, []int{1, 2}, []int{2, 1})
```

```go
_ = assert.ElementsMatch(t, []int{2, 2}, []int{1, 1})

> elements differ
> 
> extra elements in list A:
> ([]interface {}) (len=2) {
> (int) 2,
> (int) 2
> }
> 
> 
> extra elements in list B:
> ([]interface {}) (len=2) {
> (int) 1,
> (int) 1
> }
> 
> 
> listA:
> ([]int) (len=2) {
> (int) 2,
> (int) 2
> }
> 
> 
> listB:
> ([]int) (len=2) {
> (int) 1,
> (int) 1
> }
```

### Approximate values

```go
import "math"

// approximate values
var (
    valuesA = map[string]float64{"pi": math.Pi}
    valuesB = map[string]float64{"pi": 22 / 7.0}
)
```

```go
// approximal values with in delta
_ = assert.InDelta(t, valuesA["pi"], valuesB["pi"], 0.001)

> Max difference between 3.141592653589793 and 3.142857142857143 allowed is 0.001, but difference was -0.0012644892673496777
```

```go
_ = assert.InDeltaMapValues(t, valuesA, valuesB, 0.001)

> Max difference between 3.141592653589793 and 3.142857142857143 allowed is 0.001, but difference was -0.0012644892673496777
```

```go
_ = assert.InDeltaSlice(t, []float64{valuesA["pi"]}, []float64{valuesB["pi"]}, 0.001)

> Max difference between 3.142857142857143 and 3.141592653589793 allowed is 0.001, but difference was 0.0012644892673496777
```

### Centains

```go
assert.Contains(t, []string{1,2}, 2)

> []string{"\x01", "\x02"} does not contain 2

result >>> false
```

```go
assert.Contains(t, []string{"foobar", "foo"}, "foobar")

result >>> true
```

```go
_ = assert.Contains(t, []int{2, 1}, "foobar")

> []int{2, 1} does not contain "foobar"
```

```go
_ = assert.NotContains(t, []string{"foo", "bar"}, "bar")

> "[foo bar]" should not contain "bar"
```

```go
assert.Contains(t, []string{"Hello", "World"}, "World")

result >>> true
```

```go
// works well with containers
assert.Contains(t, "Hello World", "World")
assert.Contains(t, []string{"Hello", "World"}, "World")
assert.Contains(t, map[string]string{"Hello": "World"}, "Hello")

result >>> true
```

```go
//but, uses keys in map for checking
assert.NotContains(t, map[string]string{"Hello": "World"}, "World")

result >>> true
```

```go
_ = assert.Subset(t, []int{1, 2, 3}, []int{4, 5})

> "[%!s(int=1) %!s(int=2) %!s(int=3)]" does not contain "%!s(int=4)"
```

```go
_ = assert.Subset(t, []int{1, 2, 3}, []int{1, 2})
```

```go
_ = assert.NotSubset(t, []int{1, 2, 3}, []int{2})

> ['\x02'] is a subset of ['\x01' '\x02' '\x03']
```

```go
// length of the container
_ = assert.Len(t, []int{1, 2, 3}, 4)

> "[%!s(int=1) %!s(int=2) %!s(int=3)]" should have 4 item(s), but has 3
```

### FileSystem

```go
$ls

> advanced-testing
> benchmarking.ipynb
> benchmarking.md
> containers
> docs.ipynb
> docs.md
> fuzzing
> httptesting
> readme.ipynb
> readme.md
> simplest-test
> testify.ipynb
> testify.md
```

```go
_ = assert.DirExists(t, "containers-go")

> unable to find file "containers-go"
```

```go
_ = assert.FileExists(t, "testify.rtf")

> unable to find file "testify.rtf"
```

```go
_ = assert.NoFileExists(t, "testify.ipynb")

> file "testify.ipynb" exists
```

### `JSON`

```go
_ = assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
```

```go
_ = assert.JSONEq(t, `{"hello": "world", "foo": "baz"}`, `{"foo": "bar", "hello": "world"}`)

> Not equal:
> expected: map[string]interface {}{"foo":"baz", "hello":"world"}
> actual  : map[string]interface {}{"foo":"bar", "hello":"world"}
> 
> Diff:
> --- Expected
> +++ Actual
> @@ -1,3 +1,3 @@
> (map[string]interface {}) (len=2) {
> - (string) (len=3) "foo": (string) (len=3) "baz",
> + (string) (len=3) "foo": (string) (len=3) "bar",
> (string) (len=5) "hello": (string) (len=5) "world"
```

### `YAML`

```go
_ = assert.YAMLEq(t, "[1,2,3]", "[3,2,3]")

> Not equal:
> expected: []interface {}{1, 2, 3}
> actual  : []interface {}{3, 2, 3}
> 
> Diff:
> --- Expected
> +++ Actual
> @@ -1,3 +1,3 @@
> ([]interface {}) (len=3) {
> - (int) 1,
> + (int) 3,
> (int) 2,
```

### objects

```go
type Var struct {
    num int
}
```

```go
assert.ObjectsAreEqualValues(Var{10}, Var{11})

result >>> false
```

```go
_ = assert.True(t, assert.ObjectsAreEqualValues(Var{10}, Var{11}))

> Should be true
```

```go
_ = assert.Same(t, &Var{1}, &Var{1})

> Not same:
> expected: 0xc000671418 &struct { 𒀸num int }{𒀸num:1}
> actual  : 0xc000671420 &struct { 𒀸num int }{𒀸num:1}
```

```go
_ = assert.NotSame(t, &Var{1}, &Var{1})
```

```go
_ = assert.IsType(t, int32(10), int64(10))

> Object expected to be of type int32, but was int64
```

### Interfaces

```go
type Var struct {
    value int
}

func (v Var) String() string {
    return string(v.value)
}
```

```go
// jupyter itself has issues with interface casting.
_ = assert.Implements(t, (*fmt.Stringer)(nil), Var{})

> struct { 𒀸value int } must implement fmt.Stringer
```

```go
_ = assert.Implements(t, (*assert.TestingT)(nil), errorf.T{})

> errorf.T must implement assert.TestingT
```

### `time`

```go
import "time"
```

```go
t1 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
t2 := time.Date(2000, 1, 1, 2, 0, 0, 0, time.UTC)
_ = assert.WithinDuration(t, t1, t2, time.Hour)

> Max difference between 2000-01-01 00:00:00 +0000 UTC and 2000-01-01 02:00:00 +0000 UTC allowed is 1h0m0s, but difference was -2h0m0s
```

### regular expressions

```go
import "regexp"
```

```go
_ = assert.Regexp(t, regexp.MustCompile("dark(wing)? duck"), "darkoduck")

> Expect "darkoduck" to match "dark(wing)? duck"
```

```go
_ = assert.NotRegexp(t, regexp.MustCompile("dark(wing)? duck"), "darkwing duck")

> Expect "darkwing duck" to NOT match "dark(wing)? duck"
```

### Errors & Panics

```go
import "errors"
```

```go
_ = assert.NoError(t, errors.New("we are all gonna die"))

> Received unexpected error:
> we are all gonna die
```

```go
_ = assert.Error(t, nil)

> An error is expected but got nil.
```

```go
// Todo
// import "context"
// 
// var ErrUno = errors.New("one")
// 
// _ = assert.ErrorAs(t, context.DeadlineExceeded, fmt.Errorf("%w: foobar", ErrUno))
```

```go
var ErrUno = errors.New("one")

_ = assert.ErrorIs(t, fmt.Errorf("%w: foobar", ErrUno), ErrUno)
```

```go
_ = assert.NotErrorIs(t, fmt.Errorf("%w: foobar", ErrUno), ErrUno)

> Target error should not be in err chain:
> found: "one"
> in chain: "one: foobar"
> "one"
```

```go
_ = assert.EqualError(t, errors.New("Dosn't work"), "Dosn't work")
```

```go
_ = assert.NotPanics(t, func() { panic("oops!") })

> func (assert.PanicTestFunc)(0x1404060) should not panic
> Panic value:	oops!
> Panic stack:	goroutine 1 [running]:
> runtime/debug.Stack(0xc0006cc648, 0x1645f60, 0xc000738da0)
> /usr/local/go/src/runtime/debug/stack.go:24 +0x9f
> github.com/stretchr/testify/assert.didPanic.func1.1(0xc0006cc9a8, 0xc0006cc997, 0xc0006cc998)
> /home/butuzov/go/pkg/mod/github.com/stretchr/testify@v1.7.0/assert/assertions.go:1013 +0x71
> panic(0x1645f60, 0xc000738da0)
> /usr/local/go/src/runtime/panic.go:965 +0x1b9
> github.com/cosmos72/gomacro/fast.(*Comp).call_builtin.func19(0xc0001d72c0)
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/fast/builtin.go:1034 +0x69
> github.com/cosmos72/gomacro/fast.funAsStmt.func1(0xc0001d72c0, 0xc0006cc888, 0x827a78)
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/fast/util.go:1284 +0x2f
> github.com/cosmos72/gomacro/fast.exec.func1(0xc0001d72c0)
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/fast/code.go:170 +0x15b
> github.com/cosmos72/gomacro/fast.(*Comp).func0ret0.func2.1()
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/fast/func0ret0.go:44 +0x66
> github.com/stretchr/testify/assert.didPanic.func1(0xc0006cc9a8, 0xc0006cc997, 0xc0006cc998, 0xc000eb8420)
> /home/butuzov/go/pkg/mod/github.com/stretchr/testify@v1.7.0/assert/assertions.go:1018 +0x6e
> github.com/stretchr/testify/assert.didPanic(0xc000eb8420, 0x7f85cc340348, 0x2143818, 0x0, 0x0, 0x1634f00)
> /home/butuzov/go/pkg/mod/github.com/stretchr/testify@v1.7.0/assert/assertions.go:1020 +0x65
> github.com/stretchr/testify/assert.NotPanics(0x7f85cc340348, 0x2143818, 0xc000eb8420, 0x2143818, 0x0, 0x0, 0x0)
> /home/butuzov/go/pkg/mod/github.com/stretchr/testify@v1.7.0/assert/assertions.go:1091 +0x77
> reflect.Value.call(0x7f85b7de5040, 0x7f85b7e29ed0, 0x13, 0x17fd3ba, 0x4, 0xc00054b9f0, 0x3, 0x2, 0x7f85b7ddb500, 0xc000eb8420, ...)
> /usr/local/go/src/reflect/value.go:476 +0x8e7
> reflect.Value.Call(0x7f85b7de5040, 0x7f85b7e29ed0, 0x13, 0xc000eb8540, 0x2, 0x2, 0x13, 0x13, 0xc00063a588)
> /usr/local/go/src/reflect/value.go:337 +0xb9
> github.com/cosmos72/gomacro/xreflect.Value.Call(0x7f85b7de5040, 0x7f85b7e29ed0, 0x13, 0xc0006cce18, 0x2, 0x2, 0x10, 0x10, 0x16ea5a0)
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/xreflect/value.go:127 +0xcf
> github.com/cosmos72/gomacro/fast.call_variadic_ret1.func19(0xc0001d60a0, 0xc0006cced0)
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/fast/call_variadic.go:265 +0x165
> github.com/cosmos72/gomacro/fast.funAsStmt.func4(0xc0001d60a0, 0x1, 0x1)
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/fast/util.go:1302 +0x2f
> github.com/cosmos72/gomacro/fast.exec.func1(0xc0001d60a0)
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/fast/code.go:170 +0x15b
> github.com/cosmos72/gomacro/fast.funAsXV.func1(0xc0001d60a0, 0x1, 0xc000738ed0, 0xc0001d6000, 0xc00054b9a0, 0x823e3f, 0xc000020000)
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/fast/util.go:827 +0x2f
> github.com/cosmos72/gomacro/fast.(*Interp).RunExpr(0xc000062aa0, 0xc00054b9a0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
> /go/pkg/mod/github.com/cosmos72/gomacro@v0.0.0-20210624153544-b4935e406a41/fast/repl.go:132 +0x1a4
> main.doEval(0xc000062aa0, 0x19c0760, 0xc00063a438, 0x19c0760, 0xc00063a450, 0xc000829400, 0x32, 0x0, 0x0, 0x0, ...)
> /go/pkg/mod/github.com/gopherdata/gophernotes@v0.7.3/kernel.go:587 +0x250
> main.(*Kernel).handleExecuteRequest(0xc0006cd6c0, 0xc000f9a6e0, 0x20, 0xc000221b98, 0x8, 0xc000f9a700, 0x20, 0xc000221ba0, 0xf, 0xc000221bb0, ...)
> /go/pkg/mod/github.com/gopherdata/gophernotes@v0.7.3/kernel.go:492 +0x69f
> main.(*Kernel).handleShellMsg(0xc0006cd6c0, 0xc000f9a6e0, 0x20, 0xc000221b98, 0x8, 0xc000f9a700, 0x20, 0xc000221ba0, 0xf, 0xc000221bb0, ...)
> /go/pkg/mod/github.com/gopherdata/gophernotes@v0.7.3/kernel.go:346 +0x25e
> main.runKernel(0x7ffcd1cd0dd9, 0x5b)
> /go/pkg/mod/github.com/gopherdata/gophernotes@v0.7.3/kernel.go:226 +0xa25
> main.main()
> /go/pkg/mod/github.com/gopherdata/gophernotes@v0.7.3/main.go:26 +0xb7
```

```go
_ = assert.Panics(t, func() { })

> func (assert.PanicTestFunc)(0x1321e40) should panic
> Panic value:	<nil>
```

```go
_ = assert.PanicsWithError(t, "damn!", func(){ })

> func (assert.PanicTestFunc)(0x1321e40) should panic
> Panic value:	<nil>
```

```go
_ = assert.PanicsWithValue(t, "damn!", func() { })

> func (assert.PanicTestFunc)(0x1321e40) should panic
> Panic value:	<nil>
```

### Sequences

```go
_ = assert.IsDecreasing(t, []int{1, 2, 0})

> "1" is not greater than "2"
> Messages:   	[]
```

```go
_ = assert.IsIncreasing(t, []int{1, 2, 0})

> "2" is not less than "0"
> Messages:   	[]
```

```go
_ = assert.IsNonDecreasing(t, []int{3, 2, 1})

> "3" is not less than or equal to "2"
> Messages:   	[]
```

```go
_ = assert.IsNonIncreasing(t, []int{1, 2, 1})

> "1" is not greater than or equal to "2"
> Messages:   	[]
```

### Conditions

```go
_ = assert.Condition(t, func() bool { return false })

> Condition failed!
```

```go
var n = 10
_ = assert.Eventually(t, func() bool { return false }, time.Second, 20 * time.Millisecond)

> Condition never satisfied
```

```go
import "math/rand" 

_ = assert.Never(t, func() bool { return rand.Int31n(10) == 4 }, time.Second, 20 * time.Millisecond)

> Condition satisfied
```

### Positivity/Negativity

```go
_ = assert.Negative(t, 1, 1)

> "1" is not negative%!(EXTRA int=0)
> Messages:   	[1]
```

```go
_ = assert.Positive(t, -1)

> "-1" is not positive%!(EXTRA int=0)
> Messages:   	[]
```

### `net/http`

```go
import "net/http"
```

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("ok"))
})
```

```go
assert.HTTPBody(handler, "GET", "/", nil)

result >>> ok
```

```go
assert.HTTPBodyContains(t, handler, "GET", "/", nil, "okey")

> Expected response body for "/?" to contain "okey" but found "ok"

result >>> false
```

```go
_ = assert.HTTPBodyNotContains(t, handler, "GET", "/", nil, "ok")

> Expected response body for "/?" to NOT contain "ok" but found "ok"
```

```go
assert.HTTPSuccess(t,`b` handler, "GET", "/", nil, "ok")
```

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
})
_ = assert.HTTPError(t, handler, "GET", "/", nil, "okey")

> Expected HTTP error status code for "/?" but received 200
```

```go
_ = assert.HTTPStatusCode(t, handler, "GET", "/", nil, 301)

> Expected HTTP status code 301 for "/?" but received 200
```

```go
_ = assert.HTTPRedirect(t, handler, "GET", "/", nil)

> Expected HTTP redirect status code for "/?" but received 200
```