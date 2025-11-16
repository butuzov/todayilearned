<!-- weight: 300 -->
<!-- menu: Dependancy Management -->

# Dependancy Management with `go.mod`

## Reading

- https://go.dev/wiki/Modules
- https://go.dev/wiki/PackagePublishing
- [Go Package, Go](https://www.johnsto.co.uk/blog/go-package-go/)
- [using replace in `go.mod`](https://thewebivore.com/using-replace-in-go-mod-to-point-to-your-local-module/)

## `go mod`

```shell
# go mod <command> [arguments]
go mod init <name>
# Updates all dependencies to lastest minor.patch version. (https://stackoverflow.com/q/66753231)
go get -u ./...

# vendore deps
go mod vendor

# update dependency (to minor.patch)
go get -u <url>

# why we have this dependency?
go mod why github.com/sirupsen/logrus
go mod graph | grep logrus

# editing
# remove from require
go mod edit -droprequire github.com/sirupsen/logrus
# add to repalce block
go mod edit -replace github.com/sirupsen/logrus=./logrus/local

# Dependencies
# downlaod missing
go mod tidy

# verify dependencies have expected content
go mod verify
```

## Troubleshooting

If case if `go mod why` useless, first resolve all `go.mod` versions on dependecies, then find what uses it.

{{% list "./dependancies/check_go_mod.py" %}}

```shell
# Then find whats uses this dependency.
cd ~/go/pkg && grep -R "github.com/user/pkg"
```

## GOPROXY

- Using `GOPRIVATE` environment variable controls which modules the go command considers to be private (not available publicly).
- Using `GOPROXY` environment variable we can tell what proxy for dependencies can be used.
- `GONOPROXY` and `GONOSUMDB` environment variables accept the same kind of glob list and override GOPRIVATE for the specific decision of whether to use the proxy and checksum database, respectively.

### Resources

- https://github.com/goproxy/goproxy

https://go.dev/wiki/GoGetProxyConfig
https://jfrog.com/blog/why-goproxy-matters-and-which-to-pick/

## Vendoring

We can save all our dependecies with `go mod vendor`.

```shell
# wouln't work with go.work, need to turn it off with GOWORK=off
GOWORK=off go mod vendor
```
