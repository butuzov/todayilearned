# tooling

## Whole Rust with brew

```shell
berw install rust
```

## Formatter

```shell
brew reinstall rustup-init
rustup toolchain install nightly
rustfmt -V
```


## Cargo

```
# binstall
curl -L --proto '=https' --tlsv1.2 -sSf \
  https://raw.githubusercontent.com/cargo-bins/cargo-binstall/main/install-from-binstall-release.sh | bash
# or from source
cargo install cargo-binstall


#
cargo binstall cargo-nextest
cargo binstall cargo-watch
```

### `watch` & `nextest`

```shell
# watch
cargo watch -c -q -x 'nextest run'
cargo watch -c -q -x 'nextest run commands::engine  --nocapture'
cargo watch -c -q -x 'nextest run engine::streams::row::tests::encode  --nocapture
```
