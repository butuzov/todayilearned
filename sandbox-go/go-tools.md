# Go Tools

## `go tool compile` 

https://golang.org/cmd/compile/

####  `$GOOS`
The list of valid GOOS values includes android, darwin, dragonfly, freebsd, linux, nacl, netbsd, openbsd, plan9, solaris, windows, and zos. 

#### `$GOARCH`

On the other hand, the list of valid GOARCH values includes 386, amd64, amd64p32, arm, armbe, arm64, arm64be, ppc64, ppc64le, mips, mipsle, mips64, mips64le, mips64p32, mips64p32le, ppc, s390, s390x, sparc, and sparc64.

```bash
# Print assembly listing to standard output
GOOS=darwin GOARCH=amd64 go tool compile -S main.go

# Debug parse tree after type checking.
go tool compile -W main.go

# Compline only package 
go tool compile package.go && ls -la
> package.o
```


## `go build`


https://golang.org/cmd/buildid/

```bash

# Verbose output of the build process
go build -x main.go

```

## `go mod` - Go Modules

The commands are:

* `download` - download modules to local cache
* `edit` - edit go.mod from tools or scripts
* `graph`  - print module requirement graph
* `init` - initialize new module in current directory
* `tidy` - add missing and remove unused modules
* `vendor` - make vendored copy of dependencies
* `verify` - verify dependencies have expected content
* `why` - explain why packages or modules are needed


```bash
# go mod <command> [arguments]
go mod init <name>
go get -u ./...

# vendore deps
go mod vendor

# update dependency
go get -u <url>
go mod vendor

# why we have this dependency?
go mod why github.com/sirupsen/logrus
# or 
go mod graph | grep logrus

# editing
# remove from require
go mod edit -droprequire github.com/sirupsen/logrus
# add to repalce block
go mod edit -replace github.com/sirupsen/logrus=./logrus/local
```

## `godoc` - Documentation

Go Documentation + Package Documentation 

```bash
# running local documentation
godoc -http :8000
```

## `go doc` 


Doc prints the documentation comments associated with the item identified by its arguments (a package, const, func, type, var, method, or struct field) followed by a one-line summary of each of the first-level items "under" that item (package-level declarations for a package, methods for a type, etc.).


```bash
# show short info about json package
go doc json
```

## `go env`  Environment information 

```bash
go env
go env -json
go env -w "GOPATH=/www/path/"
```

## `go list`  - List lists the named packages 

https://dave.cheney.net/2014/09/14/go-list-your-swiss-army-knife

```bash
# std
go list std

# The -json flag tells you everything the go tool knows about a package:
# checkout
#     go help list
# to see info about Package struct
go list -json std

# Custom formated output
go list -f '{{.Doc}}' std

# Show Package Files
go list -f '{{join .GoFiles " "}}' bytes

# No Docs
go list -f '{{if not .Doc}}{{.ImportPath}}{{end}}' std

# Find Packages that depands of golang.org/x/oauth2 package
go list -f '{{range .Deps}}{{if eq . "golang.org/x/oauth2"}}{{$.ImportPath}}{{end}}{{end}}' all

# Broken
go list -e -f '{{with .Error}}{{.}}{{end}}' all

# go list printing line counts
for pkg in $(go list golang.org/x/oauth2/...); do
    wc -l $(go list -f '{{range .GoFiles}}{{$.Dir}}/{{.}} {{end}}' $pkg) | \
        tail -1 | awk '{ print $1 " '$pkg'" }'
done | sort -nr

> 617 golang.org/x/oauth2/google
> 600 golang.org/x/oauth2
> 357 golang.org/x/oauth2/internal
> 160 golang.org/x/oauth2/jws
> 147 golang.org/x/oauth2/jwt
> 112 golang.org/x/oauth2/clientcredentials
> 22 golang.org/x/oauth2/paypal
> 16 golang.org/x/oauth2/vk
> 16 golang.org/x/oauth2/odnoklassniki
> 16 golang.org/x/oauth2/linkedin
> 16 golang.org/x/oauth2/github
> 16 golang.org/x/oauth2/facebook


# Dependency Graph
(   echo "digraph G {"
    go list -f '{{range .Imports}}{{printf "\t%q -> %q;\n" $.ImportPath .}}{{end}}' \
        $(go list -f '{{join .Deps " "}}' time) time
    echo "}"
) | dot -Tsvg -o time-deps.svg

 
deps() {
        go list -f '{{ join .Deps  "\n"}}' . | grep pkgname
}

deps

# printing paltform specific files
env GOOS=darwin go list -f '{{ .GoFiles }}' github.com/pkg/term
[term.go term_bsd.go]
```


## `go test`
[Go testing Package](testing)

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