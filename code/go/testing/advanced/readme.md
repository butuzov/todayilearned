<!-- menu: Advanced -->
# Advanced Testing

## Talks

### `Advanced Testing with Go` by  Mitchell Hashimoto

- [`video`](https://www.youtube.com/watch?v=yszygk1cpEc)
- [`slides`](https://speakerdeck.com/mitchellh/advanced-testing-with-go)
- [`code`](./advanced-testing)

__Notes__:

- `golden_test.go` Testing with golden files and updating them (while using `testdata`).
    ```shell
    # udpate
    > go test -update
    ```
- `testing.go` Having `testing.go` or `testing_.go` (or `file_testing.go`) to provide Testing API.
- `network_testing.go` server listner fixture (see https://github.com/hashicorp/go-plugin ).
- `subprocessing_test.go` testing subprocesses with Entrypoint

### `Testing Techniques` by Andrew Gerrand

- [video](https://www.youtube.com/watch?v=ndmB0bj7eyw)
- [slides](https://go.dev/talks/2014/testing.slide#1)


__Notes___:
- Table-driven Tests

### `Go Testing By Example` by Russ Cox

1. Make it easy to add new test cases.
1. Use test coverage to find untested code.
1. Coverage is no substitute for thought.
1. Write exhaustive tests.
1. Separate test cases from test logic.
1. Look for special cases.
1. If you didn’t add a test, you didn’t fix the bug.
1. Not everything fits in a table.
1. Test cases can be in testdata files.
1. Compare against other implementations.
1. Make test failures readable.
1. If the answer can change, write code to update them.
1. Use txtar for multi-file test cases.
1. Annotate existing formats to create testing mini-languages.
1. Write parsers and printers to simplify tests.
1. Code quality is limited by test quality.
1. Scripts make good tests.
1. Try rsc.io/script for your own script-based test cases.
1. Improve your tests over time.
1. Aim for continuous deployment.



## Testing Patterns

> A test is not a unit test if:
> * It talks to the database
> * It communicates across the network
> * It touches the file system
> * It can't run at the same time as any of your other unit tests
> * You have to do special things to your environment to run it.


## Reading

* [Writing Table Driven Tests in Go](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go) by Dave Cheney
* [5 advanced testing techniques in Go](https://segment.com/blog/5-advanced-testing-techniques-in-go/) [ru](https://habr.com/ru/company/otus/blog/452772/)
* [ru] [httptest](https://dou.ua/lenta/articles/golang-httptest/)
* [Testing Web Apps in Go](http://markjberger.com/testing-web-apps-in-golang/)
* https://golang.org/pkg/cmd/go/internal/test/
