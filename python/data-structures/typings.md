# Typings

* [PEP 484](https://www.python.org/dev/peps/pep-0484), 
  [PEP 526](https://www.python.org/dev/peps/pep-0526), 
  [PEP 544](https://www.python.org/dev/peps/pep-0544), 
  [PEP 563](https://www.python.org/dev/peps/pep-0563),
  [PEP 560](https://www.python.org/dev/peps/pep-0560),
  
* https://docs.python.org/3/library/typing.html

---
```python
def hello(world: str) -> str:
    return 'Hello {}'.format(world)
```
---

### Talks

* [PyCon2019: Type hinting (and mypy)](https://www.youtube.com/watch?v=hTrjTAPnA_k) by Bernat Gabor (`mypy`) - [slides](https://gaborbernat.github.io/pycon-us-2019/)
* [PyCon2019: Getting to Three Million Lines of Type-Annotated Python](https://www.youtube.com/watch?v=mh9jQSxzv0c) by Michael Sullivan (`mypy`)
* [PyCon2019: Leveraging the Type System to Write Secure Applications](https://www.youtube.com/watch?v=ZplZ8ZBwu0Q) by Shannon Zhu (`pyre`)

```python
import typingM
```

## Typing Generator

* [MonkeyType](https://github.com/instagram/MonkeyType)

## Static Checkers

* `mypy`     [github](https://github.com/python/mypy), [docs](http://mypy-lang.org/)
* `pyre`     [github](https://github.com/facebook/pyre-check), [docs](https://pyre-check.org/docs/overview.html)
* `pyright`  [github](https://github.com/Microsoft/pyright)
* `pytype`   [github](https://github.com/google/pytype)

## `mypy` - optional typechecker

```ini
# setup.cfg contents
# MyPy
# https://mypy.readthedocs.io/en/stable/config_file.html
[mypy]
# Import discovery
ignore_missing_imports                            = True
follow_imports                                    = silent
follow_imports_for_stubs                          = True

# Disallow dynamic typing
# - There are bug in mypy at the moment so we cant really
#   disallow Any for a moment.
disallow_any_unimported                           = False
disallow_any_expr                                 = False
disallow_any_decorated                            = False
disallow_any_explicit                             = False
disallow_any_generics                             = False
disallow_subclassing_any                          = False

# Untyped definitions and calls
disallow_untyped_calls                            = True
disallow_untyped_defs                             = True
check_untyped_defs                                = True
disallow_incomplete_defs                          = True
disallow_untyped_decorators                       = True

# None and optional handling
no_implicit_optional                              = True
strict_optional                                   = True

# Configuring warnings
warn_unused_ignores                               = True
warn_no_return                                    = True
warn_return_any                                   = True

# Suppressing errors
show_none_errors                                  = True

# Miscellaneous strictness flags
allow_redefinition                                = True
strict_equality                                   = True

# Import discovery
namespace_packages                                = True

# Incremental mode
incremental                                       = True
skip_version_check                                = True
cache_dir                                         = .cache/mypy

# Configuring error messages
show_error_context                                = True
show_column_numbers                               = True

# Advanced options
pdb                                               = True
show_traceback                                    = True
warn_incomplete_stub                              = True

# Miscellaneous
verbosity                                         = 0
warn_redundant_casts                              = True
warn_unused_configs                               = True
```

## Generic Functions

```python
from typing import TypeVar, Union

# T represents generic parameter 
T = TypeVar('T', int, float)
def add(a: T, b:T) -> T:
    return a+b

print(add(1., 2))
output > 3.0

print(add(1., 2.))
output > 3.0
```

## Generators

```python
from typing import Generator

def fib(n: int) -> Generator[int, None, None]:
    a, b = 0, 1
    while a < n:
        yield a
        a, b = b, a+b

print([x for x in fib(6)])
output >  [0, 1, 1, 2, 3, 5]
```

## Iterators

```python
from typing import Iterator

def fib(n: int) -> Iterator[int]:
    a, b = 0, 1
    while a < n:
        yield a
        a, b = b, a+b

print([x for x in fib(6)])
output >  [0, 1, 1, 2, 3, 5]
```

## Cheating

```python
from typing import Any

def add(a: Any, b: Any) -> Any:
    return a+b

```

