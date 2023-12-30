# Test Tooling

## Helpers

### `gotests`
- https://github.com/cweill/gotests

### `gotestsum`
- https://github.com/gotestyourself/gotestsum

### `tparse`
- https://github.com/mfridman/tparse

## Libraries

* `testify` - [ipynb](testify.ipynb), [markdown](testify.md)
* `ginkgo` & `gomega`


## Editors Integration

**Visual Studio Code**

```json
    // gotests wraper, allow to pass own options to gotests
    // the next options will generate parallel tests as maps using testify template
    "go.generateTestsFlags" : [ "-template=testify", "-named", "-parallel" ],
    // run tests in verbose with count=1 (cleans cache)
    "go.testFlags" : ["-v", "-count=1"],
```
