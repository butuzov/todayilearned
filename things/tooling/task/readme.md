<!-- menu: "`task`" -->
# Taskfile

[taskfile.dev](https://taskfile.dev/) runner I like to use instead of Make.

## Install

```shell
brew install go-task
```

## Recipes

### Cross Platform Dev

```yaml
tasks:
  default:
    sources:
    - "./*.py"
    - "./*.kv"
    method: timestamp
    cmds:
    - cmd: ../.venv/bin/python app.py
      platforms: [darwin]
    - cmd: .win-venv/Scripts/python app.py
      platforms: [windows]
```

### Short Style calls

```yaml
tasks:
  build:        go build -trimpath -o bin/app ./cmd/app/
  build-race:   go build -race -trimpath -o bin/app ./cmd/app/
  install:      go install -trimpath -v -ldflags="-w -s" ./cmd/app/
  install-race: go install -race -trimpath -v -ldflags="-w -s" ./cmd/app/
  cover:        go tool cover -html=coverage.cov
  lints:        golangci-lint run --no-config ./... -D deadcode
```

### Shuting down web service

```yaml
  cmds:
    - go build -o ./api-server ./cmd/api-server/
    - cmd: lsof -t  -i:5000 | xargs kill -9
      silent: true
      ignore_error: true
    - ./api-server
```

## Tips & Tricks

```shell
# If default task set to monitor sources, use `watch` alias.
alias watch='task -w -v -f'
```


## Unsorted

```yaml
env:
  GOBIN: '{{.PWD}}/bin'
#   PATH:  '{{.PWD}}/bin:{{.PATH}}' wouldn;t work
```
