<!-- menu: Go -->
# Go(lang)?

## Hello World

```go
package main

func main(){
  println("Hello World")
}
```

## Install

### `brew`

```shell
brew isntall go
brew upgrade go
```

### `gvm`

```shell
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

gvm install go1.21 -B
gvm use go1.21
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.7
```
