# Go Tips & Tricks

Check on weekly basics in twitter

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

## unsorted

#golang perf tip: use benchcmd along with benchstat to benchmark and compare the performance of a standalone program or script.

For example, I'm using it now to benchmark the speed of a Go code generator!

https://github.com/golang/tools/tree/master/cmd/benchcmp
https://github.com/aclements/go-misc/blob/master/benchcmd/main.go

ci at tip
https://github.com/go-reform/reform/blob/2342d78b8a0dc22240fe5649b9293cfc71c02535/.github/workflows/ci.yaml