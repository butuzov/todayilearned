# Tooling

## Package Managers

- [`pip`](./pip/readme.md)
- [`poetry`](./poetry/readme.md)

## Code Quality

### `radon`

[Radon](https://radon.readthedocs.io/en/latest/) is a Python tool which computes various code metrics. Supported metrics are:

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
* https://radon.readthedocs.io/en/latest/