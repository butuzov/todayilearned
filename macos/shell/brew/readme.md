<!-- tags: package manager, mac, macosx -->
<!-- menu: "`brew`" -->
<!-- seotitle: "Homebrew" -->
# Homebrew `brew`

## Installing it

1. via shell call

```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

2. Using [`.pkg` Installer](https://github.com/Homebrew/brew/releases)

## Usage

### ENV variables

- `brew --prefix` same as  `brew --repository`

### Command Line Recipes

```shell
brew --prefix         # show home dir

# List packages
brew list -1 # 1 app per line
brew list -c # casks

# Software Lyfecycle Install/Upgrade/Uninstall
brew install --cask nameapp # install cask
brew bundle --file=Brewfile # from brew file
brew install name/name/name # tap & install
brew install pkgname@8.4.1  # specifing package
brew reinstall --build-from-source pkgname
brew install https://raw.github.com/.../pkgname.rb # formula

brew uninstall pkgname
brew pin pkgname            # lock this version
brew unpin pkgname          # unlock (can be upgraded)

# Info
brew search pkgname        # search for taps
brew info pkgname          # info
brew switch pkgname 9.1.5  # switching to version 9.1.5

# Brew iteself
brew update                # Updates Homebrew repo
brew doctor                # Should be run sometimes (maintanance)
```
## Services

You can use brew to manage services

```shell
brew services
brew install grafana prometheus
brew services start grafana
brew services start prometheus
brew services stop prometheus
brew services list
brew services run prometheus
brew services kill prometheus
brew services restart prometheus
brew services info prometheus
brew services cleanup
```

https://www.raycast.com/erics118/brew-services


## Mac OS App

- [Cork.app](https://github.com/buresdv/Cork) is nice @ [corkmac.app](https://corkmac.app/)


## Packaging

- TODO: creating formula
- TODO: `goreleaser`
