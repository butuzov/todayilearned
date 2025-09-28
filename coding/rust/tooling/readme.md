<!-- weight: 4 -->
<!-- menu: Tooling -->

# Tooling

## Cargo

### Creating new project

```shell
cargo new
cargo new hello_world
```

### Building && Running

```shell
# Build in debug mode
cargo build 
# Build in release mode
cargo build --release
cargo run
```

#### Multiple Binaries

```toml
# Cargo.toml
[[bin]]
name = "daemon"
path = "src/daemon/bin/main.rs"

[[bin]]
name = "client"
path = "src/client/bin/main.rs"
```

```shell
# Run Cargo
cargo run --bin example
```

### Managing Dependencies

```shell
# Installing Package Verson
cargo install --version=0.8.2 sqlx-cli
cargo install sqlx-cli@0.8.2
cargo install --version=0.8.2 sqlx-cli --no-default-features --features postgres

# Using Git (Main Branch)
cargo install --git repo_url
cargo install --git repo_url --branch|tag|rev branch_name|tag|commit_hash

# List All Installed
cargo install --list

# Add Dependency
cargo add chrono

# Update all
cargo update
```

## Jupyter

https://github.com/evcxr/evcxr/blob/main/evcxr_jupyter/

```
cargo install --locked evcxr_jupyter --features edition2024
evcxr_jupyter --install
rustup component add rust-src
```

## Helper Tools

### `nextest`

```shell
cargo binstall cargo-nextest
# Running Test
cargo nextest run
cargo nextest run commands::login::tests
cargo nextest rin engine::streams::row::tests::encode  --nocapture
```

### `watch`

```shell
cargo binstall cargo-watch
cargo watch -c -q -x 'run'
cargo watch -c -q -x 'nextest run'
# Excluding via .gitignore
```
