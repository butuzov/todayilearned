use criterion::{black_box, criterion_group, criterion_main, Criterion};
use hello_add::add;

fn bench_addition(c: &mut Criterion) {
    c.bench_function("add two numbers", |b| {
        b.iter(|| add(black_box(2), black_box(2)))
    });
}

criterion_group!(benches, bench_addition);
criterion_main!(benches);
