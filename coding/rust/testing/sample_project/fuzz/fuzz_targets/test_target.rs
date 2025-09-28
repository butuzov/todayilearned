#![no_main]
use libfuzzer_sys::fuzz_target;
use hello_add::add;

fuzz_target!(|data: (i32, i32)| {
    let (a, b) = data;
    let _ = add(a, b);
});
