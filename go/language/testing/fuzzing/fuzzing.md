
# Fuzzing

* Fuzzing the new unit testing - [`video`](https://www.youtube.com/watch?v=EJVp13f_aIs), [`slides`](https://go-talks.appspot.com/github.com/dvyukov/go-fuzz/slides/fuzzing.slide#1). [`git`](https://github.com/dvyukov/go-fuzz)
* [go-fuzz github.com/arolek/ase](https://medium.com/@dgryski/3c74d5a3150c) by Damian Gryski
* https://mijailovic.net/2017/07/29/go-fuzz/
* [DNS parser, meet Go fuzzer](https://blog.cloudflare.com/dns-parser-meet-go-fuzzer/)

* parsiya go-fuzz articles [iprange (part 1)](https://parsiya.net/blog/2018-04-29-learning-go-fuzz-1-iprange/), [goexif (part 2)](https://parsiya.net/blog/2018-05-05-learning-go-fuzz-2-goexif2/)

## Quick Start

```bash
# install go-fuzz and go-fuzz-build
go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build

# donload samples
git clone https://github.com/dvyukov/go-fuzz-corpus
cd go-fuzz-corpus/png

# fuzzing time
go-fuzz-build
go-fuzz
```

## Continuous Fuzzing

Just as unit-testing, fuzzing is better done continuously.

Currently there are 2 services that offer continuous fuzzing based on go-fuzz:

- [fuzzit.dev](https://fuzzit.dev/) ([tutorial](https://github.com/fuzzitdev/example-go))
- [fuzzbuzz.io](https://fuzzbuzz.io/) ([tutorial](https://docs.fuzzbuzz.io/getting-started/find-your-first-bug-in-go))

## How to...

1. `fuzz.go` with `func Fuzz([]byte) int` in package directory.
2. Write a case in `func Fuzz`

```go
package main

import "github.com/butuzov/somepkg"

func Fuzz(data []byte) int {
    r, err := Functionality(data)
    if err != 0 {
        return 0
    }
    return 1
}
```
3. Create a sample corpus file (in `corpus` directory) (1 per example)
4. Run `go-fuzz-build`
5. Run `go-fuzz` 

## Analyzing the Crash

* `suppressions` - crash logs
* `crashers` - found __bad__ input
  * `hash` - input
  * `hash.quoted` - input as string
  * `hash.output` - crash dump

## Reproducing crash

```go
// Sample app to test crash a5 for xor-gate/goexif2.
package main

import (
	"fmt"
	"os"

	"github.com/butuzov/somepkg"
)

func main() {
	f, err := os.Open("crashers/HASH")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_ = Fuzz(f)
}
```
