# Github Actions

## Using local builders with [`act`](https://github.com/nektos/act)

  1) `butuzov/dots` -> `butuzov/act-go`
  2) `.actrc`

  ```
  --platform ubuntu-latest=butuzov/act-go:latest --env DRY_RUN=1
  ```

  3) `act push` to run test

### Simple pipeline example

```yaml

```


### Using TIP version in Github actions

_(GitHub Search Terms: `go build tip matrix`)_

```yaml
---
name: Clean
on:
  schedule:
    # run every day
    - cron: "0 0 * * *"

jobs:
  clean-caches:
    name: Clean caches
    timeout-minutes: 5

    strategy:
      fail-fast: false
      matrix:
        go-version:
          - 1.15.x
        os: [ubuntu-latest]
        may-fail: [false]
        include:
          - go-version: tip
            os: ubuntu-latest
            may-fail: true

    continue-on-error: ${{ matrix.may-fail }}
    runs-on: ${{ matrix.os }}

    env:
      GOFLAGS: -v -mod=readonly

    steps:
      - name: Set up Go release
        if: matrix.go-version != 'tip'
        env:
          # to avoid error due to `go version` accepting -v flag with an argument since 1.15
          GOFLAGS: ""
        uses: percona-platform/setup-go@v2
        with:
          go-version: ${{ matrix.go-version }}

      - name: Set up Go tip
        if: matrix.go-version == 'tip'
        run: |
          git clone --depth=1 https://go.googlesource.com/go $HOME/gotip
          cd $HOME/gotip/src
          ./make.bash
          echo "GOROOT=$HOME/gotip" >> $GITHUB_ENV
          echo "$HOME/gotip/bin" >> $GITHUB_PATH

      - name: Check out code into the Go module directory
        uses: percona-platform/checkout@v2
        with:
          lfs: true

      - name: Enable Go modules cache
        uses: percona-platform/cache@v2
        with:
          path: ~/go/pkg/mod
          key: ${{ matrix.os }}-go-${{ matrix.go-version }}-modules-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ matrix.os }}-go-${{ matrix.go-version }}-modules-

      - name: Enable Go build cache
        uses: percona-platform/cache@v2
        with:
          path: ~/.cache/go-build
          key: ${{ matrix.os }}-go-${{ matrix.go-version }}-build-${{ github.ref }}-${{ hashFiles('**') }}
          restore-keys: |
            ${{ matrix.os }}-go-${{ matrix.go-version }}-build-${{ github.ref }}-
            ${{ matrix.os }}-go-${{ matrix.go-version }}-build-

      - name: Clean Go modules cache
        run: go clean -modcache
```
