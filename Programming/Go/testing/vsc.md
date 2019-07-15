
# Visual Studio Code

Some of the tips and tricks to make go testign even better...

## Go Extension

<kbd>⌘</kbd> + <kbd>P</kbd> - `Go: Test File`

## Go Test Explorer

<kbd>⌘</kbd> + <kbd>P</kbd> - `go Test Explorer: ...`

## Code Runner

`runme_go`
```bash
#  contents
# @ https://github.com/butuzov/dots/blob/master/bin/runme_go
runme_go(){
   FILE="${1}"
    if [[ "${FILE}" =~ "_test.go" ]]; then
        cd "$(dirname "${FILE}")"
        go test -timeout 30s -run $(cat "${FILE}" | sed -n 's/func.*\(Test.*\)(.*/\1/p' | xargs | sed 's/ /|/g') -v
        exit $?
    fi

    if [[ $(basename "${FILE}") == "main.go" ]]; then

        cd "$(dirname "${FILE}")"
        go run .
        exit $?
    fi

    go run "${FILE}"
    exit $?
}

runme_go $1
```
`settings.json`
```json
...
"code-runner.executorMap": {
    "go": "runme_go"
},
...
```
