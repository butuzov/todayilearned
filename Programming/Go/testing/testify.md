
# Testing Using `testify`

* [`github.com/stretchr/testify/assert`](https://github.com/stretchr/testify)
* https://godoc.org/github.com/stretchr/testify/assert

**Using `assert.Func` vs using `assert.Funcf`**

It doesn't matter, usually Func implements message and interface values.

```go
import "github.com/stretchra/testify"
```

## Helpers

* `FailNow` - fail
* `Fail` - fail
* `CallerInfo` - stack trace

## Asserts

### Bools

```go
func TestFunc(t *testing.T) {
    assert.True(t, false)
    assert.True(t, false, "Isn't this should be true?")
    assert.True(t, false, "Isn't this should be %s?", "true")

    assert.False(t, true)
    assert.False(t, true, "Isn't this should be false?")
    assert.False(t, true, "Isn't this should be %s?", "false")
}
```

### Zeros and Empty

```go
func TestFunc(t *testing.T) {
    var i int

    assert.NotZero(t, i)
    assert.NotZero(t, i, "Suppose to be non zero")
    assert.NotZerof(t, i, "(%+v) Suppose to be non zero", i)

    i = 1

    assert.Zero(t, i)
    assert.Zero(t, i, "Should be zero value")
    assert.Zero(t, i, "(%+v) Should be zero value", i)

    // emptines (zero value or len=0) of value
    var b bool
    assert.Empty(t, b)
}
```

### `nil`

```go
func TestFunc(t *testing.T) {
    assert.Nil(t, nil)
    assert.NotNil(t, 1) 
}
```

### Comparing Values

```go
// Should be same type
assert.Greater(t, 100, 99)
assert.GreaterOrEqual(t, 100, 99)
assert.GreaterOrEqual(t, 100, 100)

assert.Less(t, 99, 100)
assert.LessOrEqual(t, 100, 100)
assert.LessOrEqual(t, 99, 100)

// Equality
assert.Equal(t, 100, 100)
assert.NotEqual(t, 99, 100)

// Type doesn't matter!
assert.EqualValues(t, int32(100), int64(100))
assert.EqualValues(t, int32(100), byte(100))

// checking and value and type
assert.Exactly(t, int32(123), int64(123))
```

We can use `assert.Equal` to check containers aswell

```go
assert.Equal(t, []int{1, 2, 3}, []int{1, 2, 3})
assert.Equal(t, map[int]int{1: 1, 2: 2, 3: 3}, map[int]int{1: 1, 3: 3, 2: 2})
```

### Slice Containers

```go
// Elements existance check
assert.ElementsMatch(t, []int{1, 2}, []int{2, 1})
assert.ElementsMatch(t, []int{2, 2}, []int{1, 1})
// Once again we checking for equal containers
assert.Equal(t, []int{1, 2, 3}, []int{1, 2, 3})
```

### Approximate values

```go
// approximate values
var (
    valuesA = map[string]float64{"pi": math.Pi}
    valuesB = map[string]float64{"pi": 22 / 7.0}
)

// approximal values with in delta
assert.InDelta(t, valuesA["pi"], valuesB["pi"], 0.01)
assert.InDeltaMapValues(t, valuesA, valuesB, 0.01)
assert.InDeltaSlice(t, []float64{valuesA["pi"]}, []float64{valuesB["pi"]}, 0.01)

// error less then epsilon
assert.InEpsilon(t, valuesA["pi"], valuesB["pi"], 0.01)
assert.InEpsilonSlice(t, []float64{valuesA["pi"]}, []float64{valuesB["pi"]}, 0.01)
```

#### Containing element in sequance, container

```go
assert.Contains(t, []int{2, 1}, 1)
assert.Contains(t, "lorem ipsum", "lorem")

assert.NotContains(t, []int{2, 1}, 10)
assert.NotContains(t, "lorem ipsum", "dolor")

// containers
assert.Contains(t, "Hello World", "World")
assert.Contains(t, []string{"Hello", "World"}, "World")
assert.Contains(t, map[string]string{"Hello": "World"}, "Hello")
// but
assert.NotContains(t, map[string]string{"Hello": "World"}, "World")

// sets 
assert.Subset(t, []int{1, 2, 3}, []int{1, 2})
assert.Subset(t, []int{1, 2, 3}, []int{2, 2})
assert.NotSubset(t, []int{1, 2, 3}, []int{4, 2})

// length of the container
assert.Len(t, []int{1, 2, 3}, 3)
```

###  File System

* `assert.DirExists(f)`
* `assert.FileExists(f)`

```go
assert.FileExists(t, "testing.ipynb")
assert.DirExists(t, "../Go")
```

### JSON

```go
assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
```

### Objects

```go
//  does not performe assertion
assert.True(t, assert.ObjectsAreEqualValues(Var{10}, Var{10}))
assert.True(t, assert.ObjectsAreEqual(Var{10}, Var{10}))

// Sampe object (via pointer)
assert.Same(t, &Var{1}, &Var{1})

// Type
assert.IsType(t, int32(10), int64(10)) 
```

### Interface Contracts

```go
type Var struct {
    value int
}

func (v Var) String() string {
    return string(v.value)
}

func TestFunc(t *testing.T) {
    assert.Implements(t, (*fmt.Stringer)(nil), Var{})
}
```

### time.Duration

```go
assert.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)

t1 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
t2 := time.Date(2000, 1, 1, 2, 0, 0, 0, time.UTC)
assert.WithinDuration(t, t1, t2, 2*time.Hour)
```

### Matching regular expression

```go
assert.Regexp(t, regexp.MustCompile("start"), "it's starting")
assert.Regexp(t, "start...$", "it's not starting")
```

### `Todo` HTTP Testsing

* `assert.HTTPBody`
* `assert.HTTPBodyContains`
* `assert.HTTPBodyNotContains`
* `assert.HTTPError`
* `assert.HTTPRedirect`
* `assert.HTTPSuccess`

### Errors & Panics

```go
// ----- Error(f) -------------------------------------------------------
// pass
assert.Error(t, errors.New("Dosn't work"))
// fail
assert.Error(t, nil)

// ----- NoError(f) -----------------------------------------------------
// pass
assert.NoError(t, nil)
// fail
assert.NoError(t, errors.New("Dosn't work"))

// ----- EqualError(f) --------------------------------------------------
// pass
assert.EqualError(t, errors.New("Dosn't work"), "Dosn't work")
// fail
assert.EqualError(t, errors.New("Dosn't work"), "I am just an error!")

// ----- Panics(f) ------------------------------------------------------
// pass
assert.Panics(t, func() { panic("oops!") })
// fail
assert.Panics(t, func() {})

// ----- NotPanics(f) ---------------------------------------------------
// pass
assert.NotPanics(t, func() {})
// fail
assert.NotPanics(t, func() { panic("oops!") })

// ----- PanicsWithValue(f) ---------------------------------------------
// pass
assert.PanicsWithValue(t, "oops!", func() { panic("oops!") })
// fail
assert.PanicsWithValue(t, "unexcaptable!", func() { panic("oops!") })
```
