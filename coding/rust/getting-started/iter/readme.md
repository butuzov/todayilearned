# Iterator

TL;DR

- Use `.iter()` when you only need to look at elements.
- Use `.iter_mut()` when you want to modify elements in place.
- Use `.into_iter()` when you need ownership and want to consume or move elements.

[iter@reddit.com](https://www.reddit.com/r/rust/comments/1nskq8q)

## Summary

| Method         | Returns  | Ownership      | Typical Use Case         |
| -------------- | -------- | -------------- | ------------------------ |
| `.iter()`      | `&T`     | Borrow only    | read elements            |
| `.iter_mut()`  | `&mut T` | Mutable borrow | modify elements in place |
| `.into_iter()` | `T`      | Ownership      | consume/move elements    |

## Example

```rust
pub fn transformer(input: Vec<(String, Command)>) -> Vec<String> {
	input.into_iter().map(|(s, c)| match c {
		Command::Uppercase => s.to_uppercase(),
		Command::Trim => s.trim().to_string(),
		Command::Append(amount) => s + &"bar".repeat(amount),
	}) .collect()
}
```

### `.map(|(s, c)| match c { ... })`

- Applies a function to each element.
- Destructures the tuple (s, c) into the string s and the command c.
- Depending on the command, produces a new String:
  - `Uppercase` → converts the string to uppercase.
  - `Trim` → removes leading and trailing whitespace.
  - `Append(amount)` → appends "bar" amount times.

### `.collect()`

- Collects the results of the iterator.
- Builds a new vector: `Vec<String>`.

### `.iter_iter()`

- Creates a consuming iterator.
- The original vector gives up its elements → we now own the Strings.
- Important because s + "bar" consumes the string (ownership).
- After calling this, the original vector can no longer be used.

## `Trait Iterator`

Implementing iterators with [`Trait Iterator`](https://doc.rust-lang.org/std/iter/trait.Iterator.html).

```rust

```
