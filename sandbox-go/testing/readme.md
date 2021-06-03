# Testing

## Advanced Testing with Go by Mitchell Hashimoto

[`video`](https://www.youtube.com/watch?v=yszygk1cpEc) & [`slides`](https://speakerdeck.com/mitchellh/advanced-testing-with-go)


* `golden_test.go` Testing with golden files and updating them (while using `testdata`).
    ```shell
    # udpate
    > go test -update
    ```

* `testing.go` Having `testing.go` or `testing_.go` (or `file_testing.go`) to provide Testing API.
* `network_testing.go` server listner fixture (see [hashicorp/go-plugin](hashicorp/go-plugin) ).
* `subprocessing_test.go` testing subprocesses with Entrypoint

## Testing Patterns


> A test is not a unit test if:
> * It talks to the database
> * It communicates across the network
> * It touches the file system
> * It can't run at the same time as any of your other unit tests
> * You have to do special things to your environment to run it.

## Videos:

* Andrew Gerrand's [Testing Techniques](https://www.youtube.com/watch?v=ndmB0bj7eyw) talk from 2014 [`slides`](https://talks.golang.org/2014/testing.slide)

 ## Standard Library `testing`

* `Examples` and generating documentations [ipynb](docs.ipynb), [markdown](docs.md)
* `Benchmarks` [ipynb](benchmarking.ipynb), [markdown](benchmarking.md)
* Testing
  * `testing/quick` 

## Reading

* [Writing Table Driven Tests in Go](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go) by Dave Cheney
* [5 advanced testing techniques in Go](https://segment.com/blog/5-advanced-testing-techniques-in-go/) [ru](https://habr.com/ru/company/otus/blog/452772/)
* [ru] [httptest](https://dou.ua/lenta/articles/golang-httptest/)
* [Testing Web Apps in Go](http://markjberger.com/testing-web-apps-in-golang/)
* https://golang.org/pkg/cmd/go/internal/test/

### Third Party Libraries

* `testify` - [ipynb](testify.ipynb), [markdown](testify.md)
* `ginkgo` & `gomega`

### Tools

* [`gotests`](https://github.com/cweill/gotests) generator
* [`gotestsum`](https://github.com/gotestyourself/gotestsum) parser
* [`tparse`](https://github.com/mfridman/tparse) parser

### Editors Integration

**Visual Studio Code** 

```json
    // gotests wraper, allow to pass own options to gotests
    // the next options will generate parallel tests as maps using testify template
    "go.generateTestsFlags" : [ "-template=testify", "-named", "-parallel" ],
    // run tests in verbose with count=1 (cleans cache)
    "go.testFlags" : ["-v", "-count=1"],
```