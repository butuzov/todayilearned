```go
$go get github.com/davecgh/go-spew/spew

import "github.com/davecgh/go-spew/spew"

stderr >>> go: downloading github.com/davecgh/go-spew v1.1.1
```

## `reflect` package 

```go
import "reflect"
import "fmt"
import "strings"
```

```go
// tiny formating objects
func f(s interface{}) string {
    d := fmt.Sprintf("%#+v", s)
    d = strings.Replace(d, "{", "{\n  ", -1) 
    d = strings.Replace(d, "}", "\n}\n", -1) 
    d = strings.Replace(d, ", ", ",\n  ", -1)  
    return d
}
```

Our test objects

```go
type Test struct {
    Hi string `json:"hi,omitempty"`
    hi string 
}
```

```go
h := Test{Hi:"Ola!", hi:"privit"}
```

```go
package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Hi string `json:"hi,omitempty"`
	hi string
}

func main() {
	h := Test{Hi: "Ola!", hi: "privit"}
	fmt.Println(reflect.ValueOf(h).FieldByName("hi").String())
}
```

```go
reflect.ValueOf(h).FieldByName("hi").String()
```

```go
func (t *Test) SayHello(name string, age int) (ok bool) {
    return true
}
```

Let's play!

```go
var t = Test{
    Hi: "!HOLA!",
    hi: "Say my name %s",
}
```

```go
spew.Dump(t)

stdout >>> (struct { Hi string "json:\"hi,omitempty\""; 𒀸hi string }) {
stdout >>>  Hi: (string) (len=6) "!HOLA!",
stdout >>>  𒀸hi: (string) (len=14) "Say my name %s"
stdout >>> }
```

```go
f(t)

result >>> struct {
result >>>    Hi string "json:\"hi,omitempty\""; 𒀸hi string 
result >>> }
result >>> {
result >>>   Hi:"!HOLA!",
result >>>   𒀸hi:"Say my name %s"
result >>> }
```

```go
spew.Dump(reflect.TypeOf(t))

stdout >>> (*reflect.rtype)(0xc00058a780)(struct { Hi string "json:\"hi,omitempty\""; 𒀸hi string })
```

```go
f(reflect.ValueOf(t))

result >>> struct {
result >>>    Hi string "json:\"hi,omitempty\""; 𒀸hi string 
result >>> }
result >>> {
result >>>   Hi:"!HOLA!",
result >>>   𒀸hi:"Say my name %s"
result >>> }
```

```go
f(reflect.KindOf(t))
```

### `Type`

```go
reflect.TypeOf(int64(1)).Kind()

result >>> int64
```

```go
value := int64(1)
```

```go
reflect.TypeOf(value).Name()

result >>> int64
```

```go
reflect.String == reflect.TypeOf(value).Kind()

result >>> false
```

```go
reflect.Int64 == reflect.TypeOf(value).Kind()

result >>> true
```

```go
reflect.TypeOf(reflect.String).Kind() == reflect.TypeOf(reflect.Uint8).Kind()

result >>> true
```

### `Setting Value`

```go
var x float64 = 1.1
s := reflect.ValueOf(&x)
_, _ = fmt.Printf("%s - %t", s.Type(), s.CanSet())

// s.SetFloat(2.3)

stdout >>> *float64 - false
```

```go
// moving by pointer to value
v := s.Elem()
v.SetFloat(2.3)
_, _ = fmt.Printf("%s - %t", v.Type(), v.CanSet())
x

stdout >>> float64 - true
result >>> 2.3
```

### `Structs`

```go
t := Test{
    Hi: "!HOLA!",
    hi: "Say my name %s",
}
x := reflect.ValueOf(t)
```

```go
f(t)

result >>> struct {
result >>>    Hi string "json:\"hi,omitempty\""; 𒀸hi string 
result >>> }
result >>> {
result >>>   Hi:"!HOLA!",
result >>>   𒀸hi:"Say my name %s"
result >>> }
```

```go
x.NumField()

result >>> 2
```

```go
// filed by index
f(x.Field(1))

result >>> "Say my name %s"
```

```go
// private field will fail
x.FieldByName("hi")

result >>> <invalid reflect.Value>
```

```go
// filed by index
f(x.Field(0))

result >>> "!HOLA!"
```

```go
// private field will fail
x.FieldByName("Hi")

result >>> !HOLA!
```

```go
tf := reflect.TypeOf(t)
f(tf)

result >>> &reflect.rtype{
result >>>   size:0x20,
result >>>   ptrdata:0x18,
result >>>   hash:0xcc19137c,
result >>>   tflag:0x0,
result >>>   align:0x8,
result >>>   fieldAlign:0x8,
result >>>   kind:0x19,
result >>>   equal:(func(unsafe.Pointer,
result >>>   unsafe.Pointer) bool)(0x8d6780),
result >>>   gcdata:(*uint8)(0xc00020cc30),
result >>>   str:-190,
result >>>   ptrToThis:0
result >>> }
```

```go
tf.NumField()

result >>> 2
```

```go
tf.Field(0)

result >>> {Hi  string json:"hi,omitempty" 0 [0] false}
```

```go
tf.Field(1)

result >>> {𒀸hi  string  16 [1] false}
```

```go
tf.FieldByName("hi")

result >>> {  <nil>  0 [] false} false
```

```go
tf.FieldByName("Hi")

result >>> {Hi  string json:"hi,omitempty" 0 [0] false} true
```

```go
v1 := tf.FieldByName("Hi")
f(v1)

stdout >>> // warning: expression returns 2 values, using only the first one: [reflect.StructField bool]
result >>> reflect.StructField{
result >>>   Name:"Hi",
result >>>   PkgPath:"",
result >>>   Type:(*reflect.rtype)(0x1634720),
result >>>   Tag:"json:\"hi,omitempty\"",
result >>>   Offset:0x0,
result >>>   Index:[]int{
result >>>   0
result >>> }
result >>> ,
result >>>   Anonymous:false
result >>> }
```

```go
tag := v1.Tag
```

```go
tag.Get("json")

result >>> hi,omitempty
```

```go
tag.Get("yaml")

result >>>
```

```go
v1,e1 := tag.Lookup("json")
fmt.Sprintf("[%+v] - %t", v1, e1)

result >>> [hi,omitempty] - true
```

### `Functions`

```go
x2 := reflect.TypeOf(t.SayHello)
```

```go
x1 := reflect.TypeOf(t)
for i := 0; i < x1.NumField(); i++ {
    fmt.Printf("%s\n", f(x1.Field(i)))  
}

stdout >>> reflect.StructField{
stdout >>>   Name:"Hi",
stdout >>>   PkgPath:"",
stdout >>>   Type:(*reflect.rtype)(0x1634720),
stdout >>>   Tag:"json:\"hi,omitempty\"",
stdout >>>   Offset:0x0,
stdout >>>   Index:[]int{
stdout >>>   0
stdout >>> }
stdout >>> ,
stdout >>>   Anonymous:false
stdout >>> }
stdout >>> 
stdout >>> reflect.StructField{
stdout >>>   Name:"𒀸hi",
stdout >>>   PkgPath:"",
stdout >>>   Type:(*reflect.rtype)(0x1634720),
stdout >>>   Tag:"",
stdout >>>   Offset:0x10,
stdout >>>   Index:[]int{
stdout >>>   1
stdout >>> }
stdout >>> ,
stdout >>>   Anonymous:false
stdout >>> }
stdout >>>
```

```go
//Count inbound parameters
x2.NumIn()

result >>> 2
```

```go
//Count outbound parameters
x2.NumOut()

result >>> 1
```

```go
// Method
x2.String()

result >>> func(string, int) bool
```

```go
x2.IsVariadic()

result >>> false
```

```go
x2.PkgPath()

result >>>
```

```go
var out []string
for i := 0; i < x2.NumIn(); i++ {
    inV := x2.In(i)
    res := fmt.Sprintf("param[%d]  = kind: (%v) name: (%v)",i, inV.Kind(), inV.Name()) 
    out = append(out, res)
}

f(out)

result >>> []string{
result >>>   "param[0]  = kind: (string) name: (string)",
result >>>   "param[1]  = kind: (int) name: (int)"
result >>> }
```

```go
var out []string
for i := 0; i < x2.NumOut(); i++ {
    inV := x2.In(i)
    res := fmt.Sprintf("param[%d]  = kind: (%v) name: (%v)",i, inV.Kind(), inV.Name()) 
    out = append(out, res)
    fmt.Printf("%s", f(inV))
}

f(out)

stdout >>> &reflect.rtype{
stdout >>>   size:0x10,
stdout >>>   ptrdata:0x8,
stdout >>>   hash:0xe0ff5cb4,
stdout >>>   tflag:0x7,
stdout >>>   align:0x8,
stdout >>>   fieldAlign:0x8,
stdout >>>   kind:0x18,
stdout >>>   equal:(func(unsafe.Pointer,
stdout >>>   unsafe.Pointer) bool)(0x8169c0),
stdout >>>   gcdata:(*uint8)(0x18368b0),
stdout >>>   str:23238,
stdout >>>   ptrToThis:434144
stdout >>> }
result >>> []string{
result >>>   "param[0]  = kind: (string) name: (string)"
result >>> }
```