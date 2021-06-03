# `golangci-lint`


<!-- TOC depthfrom:2 depthto:3 -->

- [Exampless](#exampless)
    - [defaults](#defaults)
    - [gocritic](#gocritic)
- [Integrating golangci-lint with...](#integrating-golangci-lint-with)
    - [github actions](#github-actions)
    - [Visual Studio Code](#visual-studio-code)
    - [shell](#shell)
- [Testing](#testing)
    - [unused - alternatives](#unused---alternatives)
    - [TODO: Code Complexity - alternatives](#todo-code-complexity---alternatives)
- [Metalinters](#metalinters)
    - [gosec](#gosec)
    - [revive](#revive)
    - [staticcheck](#staticcheck)
    - [Commands](#commands)
    - [Update golangci-version.](#update-golangci-version)

<!-- /TOC -->

This is a set of demos/examples of different codes gathered into one place in order to:

1. perform visual (user) testing of `golangci-lint`
2. test some linters implementations.
3. define (overlapping) boundaries of `golangci-lint` and other linters


## Exampless

### `defaults`

Running `golangci-lint` with defaults (default `.goalngci.yaml`)

```shell
golangci-lint run -c .golangci.default.yml run testdata
```

### [`gocritic`](https://go-critic.github.io/)

`gocritic` updates can late a bit to be merged into `golangci-lint`.

```shell
# check if samples require updates.
./scripts/updates_for_gocritic

# run checks
golangci-lint run -c .golangci.gocritic.yml run testdata-gocritic

# check currently supported checks in golangci-lint
make compare_gocritic
```


## Integrating `golangci-lint` with...

### github actions

```yaml
# .github/workflows/main.yaml
name: Continuous Integration

on:
  push:
    tags:
      - v*
    branches: [ "main", "master", "develope" ]
  pull_request:

jobs:
  analysis:
    name: Static Code Analysis
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Golang Ci Linter
        uses: golangci/golangci-lint-action@v2.3.0
        with:
          version: latest
          args: -c .github/.golangci.yml
```

### Visual Studio Code

```json
{
  "go.lintTool" : "golangci-lint",
  "go.lintFlags" : [
        "-c",
        "~/.dotfiles/.golangci.yml",
        "--issues-exit-code=0"
  ]
}
```

### `shell`

```bash
golangci-lint -c .golangci.gocritic.yml run testdata-gocritic
```


## Testing

```shell
# show last 3 unimplemented linters
# run tests on not existing examples
# grep results
make update-linters | head -3 && echo "--------" &&  make test... && echo "-------" || make check | grep "funlen"
```

### unused - alternatives


checking           | unused | deadcode | structcheck  | varcheck  | unparam
-------------------|--------|----------|--------------|-----------|---------
variable           | x      | x        |              | x         |
func               | x      | x        |              |           |
method             | x      |          |              |           |
const              |        | x        |              | x         |
type               | x      | x*       |              |           |
struct fields      |        |          | x            |           |
interface          | x      | x        |              |           |
parameter          |        |          |              |           | x

- `x*` can't find if the type really unused or just have methods.


### TODO: Code Complexity - alternatives

- [ ] gocognit
- [ ] cyclop
- [ ] revive

## Metalinters

### `gosec`

`// research`
### `revive`

`// research`

### `staticcheck`

- [ ] S1001 : Replace for loop with call to copy
- [ ] S1002 : Omit comparison with boolean constant
- [ ] S1003 : Replace call to strings.Index with strings.Contains
- [ ] S1004 : Replace call to bytes.Compare with bytes.Equal
- [ ] S1005 : Drop unnecessary use of the blank identifier
- [ ] S1006 : Use for { ... } for infinite loops
- [ ] S1007 : Simplify regular expression by using raw string literal
- [ ] S1008 : Simplify returning boolean expression
- [ ] S1009 : Omit redundant nil check on slices
- [ ] S1010 : Omit default slice index
- [ ] S1011 : Use a single append to concatenate two slices
- [ ] S1012 : Replace time.Now().Sub(x) with time.Since(x)
- [ ] S1016 : Use a type conversion instead of manually copying struct fields
- [ ] S1017 : Replace manual trimming with strings.TrimPrefix
- [ ] S1018 : Use copy for sliding elements
- [ ] S1019 : Simplify make call by omitting redundant arguments
- [ ] S1020 : Omit redundant nil check in type assertion
- [ ] S1021 : Merge variable declaration and assignment
- [ ] S1023 : Omit redundant control flow
- [ ] S1024 : Replace x.Sub(time.Now()) with time.Until(x)
- [ ] S1025 : Don't use fmt.Sprintf("%s", x) unnecessarily
- [ ] S1028 : Simplify error construction with fmt.Errorf
- [ ] S1029 : Range over the string directly
- [ ] S1030 : Use bytes.Buffer.String or bytes.Buffer.Bytes
- [ ] S1031 : Omit redundant nil check around loop
- [ ] S1032 : Use sort.Ints(x), sort.Float64s(x), and sort.Strings(x)
- [ ] S1033 : Unnecessary guard around call to delete
- [ ] S1034 : Use result of type assertion to simplify cases
- [ ] S1035 : Redundant call to net/http.CanonicalHeaderKey in method call on net/http.Header
- [ ] S1036 : Unnecessary guard around map access
- [ ] S1037 : Elaborate way of sleeping
- [ ] S1038 : Unnecessarily complex way of printing formatted string
- [ ] S1039 : Unnecessary use of fmt.Sprint
- [ ] S1040 : Type assertion to current type
- [ ] SA1000 : Invalid regular expression
- [ ] SA1001 : Invalid template
- [ ] SA1002 : Invalid format in time.Parse
- [ ] SA1003 : Unsupported argument to functions in encoding/binary
- [ ] SA1004 : Suspiciously small untyped constant in time.Sleep
- [ ] SA1005 : Invalid first argument to exec.Command
- [ ] SA1006 : Printf with dynamic first argument and no further arguments
- [ ] SA1007 : Invalid URL in net/url.Parse
- [ ] SA1008 : Non-canonical key in http.Header map
- [ ] SA1010 : (*regexp.Regexp).FindAll called with n == 0, which will always return zero results
- [ ] SA1011 : Various methods in the strings package expect valid UTF-8, but invalid input is provided
- [ ] SA1012 : A nil context.Context is being passed to a function, consider using context.TODO instead
- [ ] SA1013 : io.Seeker.Seek is being called with the whence constant as the first argument, but it should be the second
- [ ] SA1014 : Non-pointer value passed to Unmarshal or Decode
- [ ] SA1015 : Using time.Tick in a way that will leak. Consider using time.NewTicker, and only use time.Tick in tests, commands and endless functions
- [ ] SA1016 : Trapping a signal that cannot be trapped
- [ ] SA1017 : Channels used with os/signal.Notify should be buffered
- [ ] SA1018 : strings.Replace called with n == 0, which does nothing
- [ ] SA1019 : Using a deprecated function, variable, constant or field
- [ ] SA1020 : Using an invalid host:port pair with a net.Listen-related function
- [ ] SA1021 : Using bytes.Equal to compare two net.IP
- [ ] SA1023 : Modifying the buffer in an io.Writer implementation
- [ ] SA1024 : A string cutset contains duplicate characters
- [ ] SA1025 : It is not possible to use (*time.Timer).Reset's return value correctly
- [ ] SA1026 : Cannot marshal channels or functions
- [ ] SA1027 : Atomic access to 64-bit variable must be 64-bit aligned
- [ ] SA1028 : sort.Slice can only be used on slices
- [ ] SA1029 : Inappropriate key in call to context.WithValue
- [ ] SA1030 : Invalid argument in call to a strconv function
- [ ] SA2000 : sync.WaitGroup.Add called inside the goroutine, leading to a race condition
- [ ] SA2001 : Empty critical section, did you mean to defer the unlock?
- [ ] SA2002 : Called testing.T.FailNow or SkipNow in a goroutine, which isn't allowed
- [ ] SA2003 : Deferred Lock right after locking, likely meant to defer Unlock instead
- [ ] SA3000 : TestMain doesn't call os.Exit, hiding test failures
- [ ] SA3001 : Assigning to b.N in benchmarks distorts the results
- [ ] SA4000 : Boolean expression has identical expressions on both sides
- [ ] SA4001 : &*x gets simplified to x, it does not copy x
- [ ] SA4003 : Comparing unsigned values against negative values is pointless
- [ ] SA4004 : The loop exits unconditionally after one iteration
- [ ] SA4006 : A value assigned to a variable is never read before being overwritten. Forgotten error check or dead code?
- [ ] SA4008 : The variable in the loop condition never changes, are you incrementing the wrong variable?
- [ ] SA4009 : A function argument is overwritten before its first use
- [ ] SA4010 : The result of append will never be observed anywhere
- [ ] SA4011 : Break statement with no effect. Did you mean to break out of an outer loop?
- [ ] SA4012 : Comparing a value against NaN even though no value is equal to NaN
- [ ] SA4013 : Negating a boolean twice (!!b) is the same as writing b. This is either redundant, or a typo.
- [ ] SA4014 : An if/else if chain has repeated conditions and no side-effects; if the condition didn't match the first time, it won't match the second time, either
- [ ] SA4015 : Calling functions like math.Ceil on floats converted from integers doesn't do anything useful
- [ ] SA4016 : Certain bitwise operations, such as x ^ 0, do not do anything useful
- [ ] SA4017 : A pure function's return value is discarded, making the call pointless
- [ ] SA4018 : Self-assignment of variables
- [ ] SA4019 : Multiple, identical build constraints in the same file
- [ ] SA4020 : Unreachable case clause in a type switch
- [ ] SA4021 : x = append(y) is equivalent to x = y
- [ ] SA4022 : Comparing the address of a variable against nil
- [ ] SA4023 : Impossible comparison of interface value with untyped nil
- [ ] SA4024 : Checking for impossible return value from a builtin function
- [ ] SA4025 : Integer division of literals that results in zero
- [ ] SA5000 : Assignment to nil map
- [ ] SA5001 : Defering Close before checking for a possible error
- [ ] SA5002 : The empty for loop (for {}) spins and can block the scheduler
- [ ] SA5003 : Defers in infinite loops will never execute
- [ ] SA5004 : for { select { ... with an empty default branch spins
- [ ] SA5005 : The finalizer references the finalized object, preventing garbage collection
- [ ] SA5007 : Infinite recursive call
- [ ] SA5008 : Invalid struct tag
- [ ] SA5009 : Invalid Printf call
- [ ] SA5010 : Impossible type assertion
- [ ] SA5011 : Possible nil pointer dereference
- [ ] SA5012 : Passing odd-sized slice to function expecting even size
- [ ] SA6000 : Using regexp.Match or related in a loop, should use regexp.Compile
- [ ] SA6001 : Missing an optimization opportunity when indexing maps by byte slices
- [ ] SA6002 : Storing non-pointer values in sync.Pool allocates memory
- [ ] SA6003 : Converting a string to a slice of runes before ranging over it
- [ ] SA6005 : Inefficient string comparison with strings.ToLower or strings.ToUpper
- [ ] SA9001 : Defers in range loops may not run when you expect them to
- [ ] SA9002 : Using a non-octal os.FileMode that looks like it was meant to be in octal.
- [ ] SA9003 : Empty body in an if or else branch
- [ ] SA9004 : Only the first constant has an explicit type
- [ ] SA9005 : Trying to marshal a struct with no public fields nor custom marshaling
- [ ] SA9006 : Dubious bit shifting of a fixed size integer value
- [ ] ST1000 : Incorrect or missing package comment
- [ ] ST1001 : Dot imports are discouraged
- [ ] ST1003 : Poorly chosen identifier
- [ ] ST1005 : Incorrectly formatted error string
- [ ] ST1006 : Poorly chosen receiver name
- [ ] ST1008 : A function's error value should be its last return value
- [ ] ST1011 : Poorly chosen name for variable of type time.Duration
- [ ] ST1012 : Poorly chosen name for error variable
- [ ] ST1013 : Should use constants for HTTP error codes, not magic numbers
- [ ] ST1015 : A switch's default case should be the first or last case
- [ ] ST1016 : Use consistent method receiver names
- [ ] ST1017 : Don't use Yoda conditions
- [ ] ST1018 : Avoid zero-width and control characters in string literals
- [ ] ST1019 : Importing the same package multiple times
- [ ] ST1020 : The documentation of an exported function should start with the function's name
- [ ] ST1021 : The documentation of an exported type should start with type's name
- [ ] ST1022 : The documentation of an exported variable or constant should start with variable's name


### Commands

golangci-lint cache clean

### Update golangci-version.

```shell
make config
# to check if we have missing examples?
make update-linters

# to generate logs to see if anything diferent in two version checks
VERSION=1.40.0 make check > test-1.40.0.log
VERSION=1.39.0 make check > test-1.39.0.log

# to run current tests.
make test...
```
