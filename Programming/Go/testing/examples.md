
# Testing (Examples)

```bash
# opening browser window wich will be available in a moment.
open http://localhost:3000/

# run godoc server on the port 300 with .1m update interval.
godoc -http=:3000 -index_interval=.1m
```

## Package Header 

You can structure it as 

1. Copyright text as `//` comments
2. Package comment. First line of it will ends as Synopsis for package. Rest in package help page.

```go
// Copyright text.

/*
	Package some contains example of the some package.

	This text eventually lands on your package documentation homepage.
*/
package some
```

## Examples

All examples suppose to be located in `_test.go` files. For easier navigation, make it `file_examples_tests.go` for `file.go`

### Example Go package 

```go
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

### Example GO Package `Examples`

```go
package mygopcg

import "fmt"

// Examples can be provided without output, just as example. Such tests wouldn't run.
func ExampleHello() {
	Hello()
}

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
