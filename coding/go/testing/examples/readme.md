# `func Example()`

All examples suppose to be located in `_test.go` files. For easier navigation, make it `file_examples_tests.go` for `file.go`

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

import "../advanced/examples/fmt"./../tooling/testing/examples/fmt"

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
