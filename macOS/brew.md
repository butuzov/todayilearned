# Brew

Managing Multiple version of Python/Go.

```bash
# get repo
git clone https://github.com/Homebrew/homebrew-core .
git log --pretty=oneline --graph -- Formula/go.rb

# or with shortcuts
git search Formula/go.rb

# general update
brew update

# brew unlink
brew unlink go

# brew install
brew install https://raw.githubusercontent.com/Homebrew/homebrew-core/64aeb88840618d7ad3fb12407f9be8d295406444/Formula/go.rb
```

brew install
| Version | Hash                                     |
|---------|------------------------------------------|
| Go 1.12 | 64aeb88840618d7ad3fb12407f9be8d295406444 |