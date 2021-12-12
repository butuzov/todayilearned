# Go General Tips & Tricks

## Compiler

###  Inlining Loop Function and & `-m`

https://github.com/golang/go/wiki/CompilerOptimizations

_Code_ 

```go
// inline-tests.go
package inlinetests

func plus(a, b int) int {
    return a + b
}

func plus_plus(a, b, c int) int {
    return plus(plus(a, b), plus(b, c))
}

func plus_iter(l ...int) (res int) {
    for _, v := range l {
        res += v
    }
    return
}
```

_Running_

```bash
go build -gcflags=-m inline-tests.go
```

_Result_ 

```
go build -gcflags=-m inline-tests.go
# command-line-arguments
./inline-tests.go:4: can inline plus
./inline-tests.go:8: can inline plus_plus
./inline-tests.go:9: inlining call to plus
./inline-tests.go:9: inlining call to plus
./inline-tests.go:9: inlining call to plus
./inline-tests.go:12: plus_iter l does not escape
```

## Coding

### Conditional (ternary) operator in Go

* https://stackoverflow.com/questions/19979178/

```go
type If bool

func (cond If) ThenElse(expr1, expr2 interface{}) interface{} {
    if cond {
        return expr1
    } 
    return expr2
}
```

```go
import "fmt"

var a, b = 100, 101

max := If(a>b).ThenElse(a, b)
min := If(a<=b).ThenElse(a, b)

_,_ = fmt.Printf("Min: Value %[1]d / Type %[1]T\n", min)
_,_ = fmt.Printf("Max: Value %[1]d / Type %[1]T\n", max)

stdout >>> Min: Value 100 / Type int
stdout >>> Max: Value 101 / Type int
```

### Packing Structs to Bytes

the code below not works, but why? 

```go
import "unsafe"
import "reflect"

var in = []uint64{1,2,3,4,10,20}
var out []byte 


var (
    inHeader  = (*reflect.SliceHeader)(unsafe.Pointer(&in))
    outHeader = (*reflect.SliceHeader)(unsafe.Pointer(&out))
)

outHeader.Data = inHeader.Data
outHeader.Len = inHeader.Len * 8
outHeader.Cap = inHeader.Cap * 8
```

## Profiling

## Tips & Tricks

### Mutable & Immutable

Implement for a same data Mutable for unsorted (so it can be sorted), and Immutable for a sorted.

## Interfaces

Dynamic Interface

```go
// interface checking 
import "bytes"

// doesn't work in jupyter
buf := bytes.Buffer{}
if f, ok := &buf.(interface { 
    Write(p []byte) (n int, err error)
}); ok {
    f.Write([]byte("foobar"))
}
```

## Testing & Benching

`#golang perf` tip: use `benchcmd` along with `benchstat` to benchmark and compare the performance of a standalone program or script.

For example, I'm using it now to benchmark the speed of a Go code generator!

- https://github.com/golang/tools/tree/master/cmd/benchcmp
- https://github.com/aclements/go-misc/blob/master/benchcmd/main.go

## Continues Integration

### `github actions`

#### Using local builders with [`act`](https://github.com/nektos/act)

1. `butuzov/dots` -> `butuzov/act-go`
2. `.actrc`

   ```
   --platform ubuntu-latest=butuzov/act-go:latest --env DRY_RUN=1
   ```

#### Using TIP version in Github actions

_(GitHub Search Terms: `go build tip matrix`)_

```yaml
---
name: Clean
on:
  schedule:
    # run every day
    - cron: "0 0 * * *"

jobs:
  clean-caches:
    name: Clean caches
    timeout-minutes: 5

    strategy:
      fail-fast: false
      matrix:
        go-version:
          - 1.15.x
        os: [ubuntu-latest]
        may-fail: [false]
        include:
          - go-version: tip
            os: ubuntu-latest
            may-fail: true

    continue-on-error: ${{ matrix.may-fail }}
    runs-on: ${{ matrix.os }}

    env:
      GOFLAGS: -v -mod=readonly

    steps:
      - name: Set up Go release
        if: matrix.go-version != 'tip'
        env:
          # to avoid error due to `go version` accepting -v flag with an argument since 1.15
          GOFLAGS: ""
        uses: percona-platform/setup-go@v2
        with:
          go-version: ${{ matrix.go-version }}

      - name: Set up Go tip
        if: matrix.go-version == 'tip'
        run: |
          git clone --depth=1 https://go.googlesource.com/go $HOME/gotip
          cd $HOME/gotip/src
          ./make.bash
          echo "GOROOT=$HOME/gotip" >> $GITHUB_ENV
          echo "$HOME/gotip/bin" >> $GITHUB_PATH

      - name: Check out code into the Go module directory
        uses: percona-platform/checkout@v2
        with:
          lfs: true

      - name: Enable Go modules cache
        uses: percona-platform/cache@v2
        with:
          path: ~/go/pkg/mod
          key: ${{ matrix.os }}-go-${{ matrix.go-version }}-modules-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ matrix.os }}-go-${{ matrix.go-version }}-modules-

      - name: Enable Go build cache
        uses: percona-platform/cache@v2
        with:
          path: ~/.cache/go-build
          key: ${{ matrix.os }}-go-${{ matrix.go-version }}-build-${{ github.ref }}-${{ hashFiles('**') }}
          restore-keys: |
            ${{ matrix.os }}-go-${{ matrix.go-version }}-build-${{ github.ref }}-
            ${{ matrix.os }}-go-${{ matrix.go-version }}-build-

      - name: Clean Go modules cache
        run: go clean -modcache
```
