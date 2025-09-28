<!-- weight: 5 -->
<!-- menu: Testing -->

# Testing

## Future Reading

<!-- TODO(butuzov): read about advanced testing in rust -->

- [Advanced Rust testing](https://rust-exercises.com/advanced-testing/) + [git: workshop](https://github.com/mainmatter/rust-advanced-testing-workshop)
- [Mastering Testing in Rust](https://codezup.com/mastering-testing-in-rust-best-practices-techniques/)
- [New Testing Frameworks for Rust in 2025](https://markaicode.com/testing-frameworks-2025-beyond-standard-library/)
- [How Doctests Work — Guillaume Gomez](https://www.youtube.com/watch?v=NmgNi6kFXZI)

## Testing

```shell
# Run Tests
cargo test
cargo test should_fail
cargo test -- --nocapture # Will show staout of println!
cargo test -- --ignored # Run tests with #[ignore]
```

```rust
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    #[should_panic]
    fn should_fail() {
        panic!("oops");
    }
}
```

## Integration Tests

```shell
# Run With
cargo test
```

```rust
# tests/test.rs
use my_crate::add;

#[test]
fn adds_two() {
    assert_eq!(add(2, 2), 4);
}
```

## Fuzzing

```shell
cargo install cargo-fuzz
cargo fuzz init
# Run Fuzzer
cargo fuzz run my_target
```

```rust
// fuzz/my_target.rs
#![no_main]
use libfuzzer_sys::fuzz_target;

fuzz_target!(|data: &[u8]| {
    // fuzz logic
    if data.len() == 3 && data[0] == b'R' && data[1] == b'S' && data[2] == b'T' {
        panic!("Found magic input!");
    }
});
```

## Benchmarking

```shell
cargo bench
```

```toml
# Cargo.toml
[dev-dependencies]
criterion = "0.5"
```

```rust
// benches/my_bench.rs:
use criterion::{black_box, Criterion, criterion_group, criterion_main};

fn bench_addition(c: &mut Criterion) {
    c.bench_function("addition", |b| {
        b.iter(|| {
            let x = black_box(2);
            let y = black_box(2);
            x + y
        })
    });
}

criterion_group!(benches, bench_addition);
criterion_main!(benches);
```
