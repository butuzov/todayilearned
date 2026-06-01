<!-- weight: 4 -->
<!-- menu: Tooling -->

# Go

> `go` is a [tool for managing](https://pkg.go.dev/cmd/go) Go source code.

- `bug` - start a bug report
- `build` - compile packages and dependencies
- [`build`](#build) - compile packages and dependencies
- `clean`- remove object files and cached files
- [`doc`](#doc) - show documentation for package or symbol
- `env` - print Go environment information
- `fix` - update packages to use new APIs
- `fmt` - gofmt (reformat) package sources
- `generate` - generate Go files by processing source
- `get` - add dependencies to current module and install them
- `install` - compile and install packages and dependencies
- [`list`](#go-list) - list packages or modules
- [`mod`](#go-mod) - module maintenance
- `work` - workspace maintenance
- `run` - compile and run Go program
- `test` - test packages
- [`tool`](#go-tool) - run specified go tool
- `version` - print Go version
- `vet` - report likely mistakes in packages

## `build`

### Examples

```bash

# show $WORK directory
go build --work .

# Dry run + -work
go build -n .

# Explicit output of the build process
go build -x .

# Escape analysis
go build -gcflags="-m"
# more info
go build -gcflags="-m=5"

# go ASM
go build -gcflags="-S"

# disbale optimizations
go build -gcflags="-N" .
# disable inlining
go build -gcflags="-l" .
# disable both
go build -gcflags="-N -l" .
```

### SSA

- [GopherCon 2017: Generating Better Machine Code with SSA - Keith Randall](https://www.youtube.com/watch?v=uTMvKVma5ms)
- [GopherCon 2015: Static Code Analysis Using SSA - Ben Johnson](https://www.youtube.com/watch?v=D2-gaMvWfQY)

```
# bce checks https://go101.org/article/bounds-check-elimination.html
go build -gcflags="-d=ssa/check_bce/debug=1"

# generate ssa dump
GOSSAFUNC=main go build .

# Garbage collector liveness bitmap generation.
# https://docs.studygolang.com/src/cmd/compile/internal/liveness/plive.go
# go build -gcflags="-live" .
go build -gcflags="-live=2" .


# adding extras
go build -gcflags="-bench=bench.out"
go build -gcflags="-race" # race detector
go build -gcflags="-memprofile=profile.out"
```

## `doc`

Doc prints the documentation comments associated with the item identified by its arguments (a package, const, func, type, var, method, or struct field) followed by a one-line summary of each of the first-level items "under" that item (package-level declarations for a package, methods for a type, etc.).

```bash
# show short info about json package
go doc json
```

## `tool`

Tool runs the go tool command identified by the arguments.

```shell

# Компілює Go-файл у об’єктний .o файл
go tool compile
go tool link
go tool vet
go tool pprof
go tool objdump

# Dissable (see symbols)
go tool nm ./bin/foobar

# Coverage
go tool cover

go tool trace
# Go toolchain
go tool dist
# Print All supported GOOS/GOARCH
go tool dist list

# Assembler
go tool asm
go tool addr2line



```

## `mod`

Go Modules a explained in separate [page](../modules/)

## `list`

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

Go list opeates with `PackagePublic`

```go
type PackagePublic struct {
	// Note: These fields are part of the go command's public API.
	// See list.go. It is okay to add fields, but not to change or
	// remove existing ones. Keep in sync with ../list/list.go
	Dir           string                `json:",omitempty"` // directory containing package sources
	ImportPath    string                `json:",omitempty"` // import path of package in dir
	ImportComment string                `json:",omitempty"` // path in import comment on package statement
	Name          string                `json:",omitempty"` // package name
	Doc           string                `json:",omitempty"` // package documentation string
	Target        string                `json:",omitempty"` // installed target for this package (may be executable)
	Shlib         string                `json:",omitempty"` // the shared library that contains this package (only set when -linkshared)
	Root          string                `json:",omitempty"` // Go root, Go path dir, or module root dir containing this package
	ConflictDir   string                `json:",omitempty"` // Dir is hidden by this other directory
	ForTest       string                `json:",omitempty"` // package is only for use in named test
	Export        string                `json:",omitempty"` // file containing export data (set by go list -export)
	BuildID       string                `json:",omitempty"` // build ID of the compiled package (set by go list -export)
	Module        *modinfo.ModulePublic `json:",omitempty"` // info about package's module, if any
	Match         []string              `json:",omitempty"` // command-line patterns matching this package
	Goroot        bool                  `json:",omitempty"` // is this package found in the Go root?
	Standard      bool                  `json:",omitempty"` // is this package part of the standard Go library?
	DepOnly       bool                  `json:",omitempty"` // package is only as a dependency, not explicitly listed
	BinaryOnly    bool                  `json:",omitempty"` // package cannot be recompiled
	Incomplete    bool                  `json:",omitempty"` // was there an error loading this package or dependencies?

	DefaultGODEBUG string `json:",omitempty"` // default GODEBUG setting (only for Name=="main")

	// Stale and StaleReason remain here *only* for the list command.
	// They are only initialized in preparation for list execution.
	// The regular build determines staleness on the fly during action execution.
	Stale       bool   `json:",omitempty"` // would 'go install' do anything for this package?
	StaleReason string `json:",omitempty"` // why is Stale true?

	// Source files
	// If you add to this list you MUST add to p.AllFiles (below) too.
	// Otherwise file name security lists will not apply to any new additions.
	GoFiles           []string `json:",omitempty"` // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
	CgoFiles          []string `json:",omitempty"` // .go source files that import "C"
	CompiledGoFiles   []string `json:",omitempty"` // .go output from running cgo on CgoFiles
	IgnoredGoFiles    []string `json:",omitempty"` // .go source files ignored due to build constraints
	InvalidGoFiles    []string `json:",omitempty"` // .go source files with detected problems (parse error, wrong package name, and so on)
	IgnoredOtherFiles []string `json:",omitempty"` // non-.go source files ignored due to build constraints
	CFiles            []string `json:",omitempty"` // .c source files
	CXXFiles          []string `json:",omitempty"` // .cc, .cpp and .cxx source files
	MFiles            []string `json:",omitempty"` // .m source files
	HFiles            []string `json:",omitempty"` // .h, .hh, .hpp and .hxx source files
	FFiles            []string `json:",omitempty"` // .f, .F, .for and .f90 Fortran source files
	SFiles            []string `json:",omitempty"` // .s source files
	SwigFiles         []string `json:",omitempty"` // .swig files
	SwigCXXFiles      []string `json:",omitempty"` // .swigcxx files
	SysoFiles         []string `json:",omitempty"` // .syso system object files added to package

	// Embedded files
	EmbedPatterns []string `json:",omitempty"` // //go:embed patterns
	EmbedFiles    []string `json:",omitempty"` // files matched by EmbedPatterns

	// Cgo directives
	CgoCFLAGS    []string `json:",omitempty"` // cgo: flags for C compiler
	CgoCPPFLAGS  []string `json:",omitempty"` // cgo: flags for C preprocessor
	CgoCXXFLAGS  []string `json:",omitempty"` // cgo: flags for C++ compiler
	CgoFFLAGS    []string `json:",omitempty"` // cgo: flags for Fortran compiler
	CgoLDFLAGS   []string `json:",omitempty"` // cgo: flags for linker
	CgoPkgConfig []string `json:",omitempty"` // cgo: pkg-config names

	// Dependency information
	Imports   []string          `json:",omitempty"` // import paths used by this package
	ImportMap map[string]string `json:",omitempty"` // map from source import to ImportPath (identity entries omitted)
	Deps      []string          `json:",omitempty"` // all (recursively) imported dependencies

	// Error information
	// Incomplete is above, packed into the other bools
	Error      *PackageError   `json:",omitempty"` // error loading this package (not dependencies)
	DepsErrors []*PackageError `json:",omitempty"` // errors loading dependencies, collected by go list before output

	// Test information
	// If you add to this list you MUST add to p.AllFiles (below) too.
	// Otherwise file name security lists will not apply to any new additions.
	TestGoFiles        []string `json:",omitempty"` // _test.go files in package
	TestImports        []string `json:",omitempty"` // imports from TestGoFiles
	TestEmbedPatterns  []string `json:",omitempty"` // //go:embed patterns
	TestEmbedFiles     []string `json:",omitempty"` // files matched by TestEmbedPatterns
	XTestGoFiles       []string `json:",omitempty"` // _test.go files outside package
	XTestImports       []string `json:",omitempty"` // imports from XTestGoFiles
	XTestEmbedPatterns []string `json:",omitempty"` // //go:embed patterns
	XTestEmbedFiles    []string `json:",omitempty"` // files matched by XTestEmbedPatterns
}
```
