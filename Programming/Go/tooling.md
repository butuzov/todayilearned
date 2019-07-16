
# Go Tools

## `go tool compile` 

https://golang.org/cmd/compile/

### $GOOS
The list of valid GOOS values includes android, darwin, dragonfly, freebsd, linux, nacl, netbsd, openbsd, plan9, solaris, windows, and zos. 

### $GOARCH

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

## `go mod`

Go Modules

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
```

## `godoc` Documentation

Go Documentation + Package Documentation 

```bash
# running local documentation
godoc -http :8000
```

## `go doc` 

**Documentation**

Doc prints the documentation comments associated with the item identified by its arguments (a package, const, func, type, var, method, or struct field) followed by a one-line summary of each of the first-level items "under" that item (package-level declarations for a package, methods for a type, etc.).

```bash
# show short info about json package
go doc json
```

## `go env` 

**Env prints Go environment information**

```bash
go env
# or
go env -json
```

## `go list` 

**List lists the named packages.**

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

```

## `go test`

[Go testing Package](testing)
