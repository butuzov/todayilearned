<!-- weight: 4 -->
<!-- menu: Testing -->

# Testing

```shell
# Coverage for local packages except osme of them...

export PKGS=$(go list ./... | grep -vE "(gotests/gotests|.*data|templates)" | tr -s '\n' ',' | sed 's/.\{1\}$//')
go test -v -covermode=count -coverpkg=$PKGS -coverprofile=coverage.cov

# Find what tests were skipped
go test -v . | grep SKIP

# run tests 10 times + verbose output (also cleans cache)
go test -v -test.count 10 .

# clean cache
go clean -testcache

# run on 2 cores
go test -v -test.count 10 -test.cpu 2 .

# run tests (filter by name)
go test -v -run S .

# run 4 runners tests
go test -v -parallel 4 .

# run 4 runners tests
go test -race .

# json
go test -v --json .

# just compiling test code
go test --exec=/bin/true ./...
go test -c pkg

# Using jq to filter output of json based export.
go test -json | jq -s 'map(select(.Test != null)) | sort_by(.Elapsed)'

# benchmarks (+ memory)
go test -json -benchmem -run=^$ -bench .
```

## Test Examples

{{% list "testing/add.go,testing/add_test.go,testing/subtract.go,testing/subtract_test.go" %}}

## Continuous integration

### github

```yaml
#.github/workflows/test.yaml
jobs:
  Tests:
    runs-on: ubuntu-latest
    strategy:
      fail-fast: false
    steps:
      # INFO: Tip of the Test Matrix
      - name: Set Conditional Environment Variable
        if: ${{ matrix.golang == '1.25' }}
        run: echo "RUN_ONCE=true" >> $GITHUB_ENV

      - uses: actions/checkout@v5
      - uses: actions/setup-go@v6
        with:
          go-version: ${{ matrix.golang }}

      # ---  📊 Dependencies -----------------------------------------------------
      # INFO: 📊 Running buils & tests (with coverage) ---------------------------
      - run: go test --cover -v -coverpkg=github.com/user/package/... -coverprofile=coverage.out github.com/user/package/... -json | tparse
```

```yaml
# taskfile.yaml
```

## Tooling

### Tests Generation `cweill/gotests`

Tests generation with https://github.com/cweill/gotests, is by defacto standard.

```json
// .vscode/settings.json
{
  "go.generateTestsFlags": ["-named", "-parallel", "10"]
}
```

### Go tests wrappers

Go tests wrappers https://github.com/mfridman/tparse, https://github.com/gotestyourself/gotestsum, https://github.com/timtatt/sift

#### `tparse`

```yaml
# taskfile.yaml
tasks:
  deps:install:tparse:
    desc: "Install tparse binary."
    cmds:
      - cmd: |
          URL="https://github.com/mfridman/tparse/releases/download/v0.18.0/tparse_linux_x86_64" && \
          sudo curl -sSL -o /usr/local/bin/tparse "${URL}" && sudo chmod +x /usr/local/bin/tparse
        vars: { VERSION: "v0.18.0" }
        platforms: [linux]

      - cmd: |
          if [ $(which tparse) == "" ]; then
             go install github.com/mfridman/tparse@latest
          fi
        platforms: [darwin]
```

```shell
go test --cover -v -coverpkg=github.com/user/package/... -coverprofile=coverage.out github.com/user/package/... -json | tparse
```

#### `gotestsum`

https://github.com/gotestyourself/gotestsum

```shell
gotestsum
# With Coverage
gotestsum -- -coverprofile=cover.out ./...

gotestsum \
  --jsonfile tmp.json.log \
  --post-run-command "bash -c '
    echo; echo Slowest tests;
    gotestsum tool slowest --num 10 --jsonfile tmp.json.log'"
```

### Testing `stretchr/testify`

[`stretchr/testify`](./testify)
