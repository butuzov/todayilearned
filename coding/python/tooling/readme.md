# Tooling

## Package Managers

- [`uv`](./uv/readme.md) ⭐⭐⭐⭐⭐
- `poetry` ⭐⭐⭐
- [`pip`](./pip/readme.md) ⭐⭐ 


## Formatter   

- [`ruff`](./ruff/readme.md) ⭐⭐⭐⭐⭐

## Code Quality (Linters)

### `ruff`

See [`ruff`](./ruff/readme.md)

### `radon`

[Radon](https://github.com/rubik/radon) is a Python tool which computes various code metrics. Supported metrics are:

- raw metrics: SLOC, comment lines, blank lines, &c.
- Cyclomatic Complexity (i.e. McCabe’s Complexity)
- Halstead metrics (all of them)
- the Maintainability Index (a Visual Studio metric)

```bash
# miantanability index
radon mi file.ry -s
radon ss file.py
radon hal file.py 
```

http://radon.readthedocs.org/

### `wily`

A command-line application for tracking, reporting on complexity of Python tests and applications.

```
wily build files.py
wili index
wily report files.py
wily diff files.py
wily graph files.py
```

* https://github.com/tonybaloney/wily
