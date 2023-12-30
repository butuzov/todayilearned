<!-- menu: golangci-lint  -->
<!-- github: https://github.com/golangci/  -->
# golangci-lint

[`golangci-lint`](https://github.com/golangci/golangci-lint) is a best linter in whole world PERIOD

## Quick Start

```shell

golangci-lint run ...
# with config
golangci-lint run -c .golangci.default.yml testdata
# test particular linter with default configuration (verbose)
golangci-lint run -v --no-config --disable-all -E lintername  ...
```

## Working With "Suite"

Suite is designed to quickly check workability of linter.

```shell
task
task gocritic -w
task golangci-lint -w
task revive -w
```

I am cehcking each linter on the basis of usefulnes it can bring.

## Opinionted linter config.

### `golangci-lint``


Some linters requier detailed work on choosing one suitable:

- gomodguard, importas, depguard, revive@import-alias-naming...
- grouper, decoder
- zerologling, loggerlint, sloglint
- unused, deadcode


#### General Approach
- reset  `limits=` at `scripts/golangci-lint`
- remove `// OK_with_config` from all files.
- create required new tests.
- pepare new list of linters or use "fresh one".
- run `task golangci-lint`

### `revive`

[`revive`](https://github.com/mgechev/revive/) is a powerfull and really good linter by iteself, and `golangci-lint` contrinutors made its best to support it.

### `gocritic`

[`gocritic`](https://github.com/go-critic/go-critic) has wide range or styling or performance rules, however it's to unrelilable (as for me). And I would prefer not to use it, if i can.

## Integrations

### With Visual Studio Code

Settings provided by [`golang.go`](https://marketplace.visualstudio.com/items?itemName=golang.go)

```json
// settings.json || .vscode/settings.json
{
  "go.lintTool": "golangci-lint",
  "go.lintFlags": [ "-v", "--sort-results" ]
}
```

### With `pre-commit`

```yaml
repos:
- repo: https://github.com/golangci/golangci-lint
  rev: v1.54.2
  hooks:
  - id: golangci-lint
  # With arguments, if needed.
  # args: [ "--no-config", "--disable-all", "-E", "maligned"]

```

### With Github Action

Can be used with in you gothub CI pipelines [`golangci/golangci-lint-action@v3`](https://github.com/golangci/golangci-lint-action). Ludovic recommends to use it separately.


```yaml
# .github/workflows/action.yaml
name: golangci-lint
on:
  push:
    branches:
    - main
  pull_request:

permissions:
  contents: read

jobs:
  golangci:
    name: lint
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
          cache: false
      - name: golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: v1.55
          args: -c .github/.golangci.yml
```





