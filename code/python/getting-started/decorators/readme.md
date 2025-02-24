# Decorators

[PEP 3129](https://www.python.org/dev/peps/pep-3129/), [PEP 318](https://www.python.org/dev/peps/pep-0318/)

### Resources

- [PythonWiki: PythonDecorators](https://wiki.python.org/moin/PythonDecorators)
- [PythonWiki: PythonDecoratorLibrary](https://wiki.python.org/moin/PythonDecoratorLibrary)
- [Stackoverflow: How to make a chain of function decorators?
  ](https://stackoverflow.com/questions/739654/)
- https://github.com/lord63/awesome-python-decorator

### Talks

- [PyCon2019: Practical decorators](https://youtu.be/MjHpMCIvwsY) by Reuven M. Lerner - [slides](https://speakerdeck.com/pycon2019/reuven-m-lerner-practical-decorators)
- [EuroPytop 2018: A Taxonomy of Decorators: A-E](https://youtu.be/pEL1THG6ysY) by Andy Fundinger - [slides](https://github.com/Ciemaar/decorator-taxonomy)

### Example

```python
def decorator(func):
    def wrapper(*args, **kwargs):
        return f"{func(*args, **kwargs)}"

    return wrapper

@decorator
def add(a, b):
    return a+b

ic(add(1,2))

> ic> add(1,2): '3'

result >>> '3'
```

## Decorators with params

```python
def decorator(func, argument):
    def wrapper(number):
        return func(number+argument)
    return wrapper

more_at_5 = functools.partial(decorator, argument=5)

@more_at_5
def pow(a):
    return a**2

ic(pow(5))

> ic> pow(5): 100

result >>> 100
```

### Using `Class` as a decorator

```python
class slim_shady:

    def __init__(self, func):
        self.func = func

    def __call__(self, name):
        return self.func(self.__class__.__name__.replace("_", " "))

@slim_shady
def name(name):
    return "My name is {}".format(name)

ic(name('oleg'))

> ic> name('oleg'): 'My name is slim shady'

result >>> 'My name is slim shady'
```

Alternate `repr()` implementation provided by [reprlib](https://docs.python.org/3/library/reprlib.html)

```python
from reprlib import recursive_repr

class ascii():
    def __init__(self, char:str) -> None:
        if len(char) > 1:
            raise ValueError("Accepting only single chars")

        self.char = char
    def __repr__(self) -> str:
        return "{} is {}".format(self.char, ord(self.char))

class lister(list):

    @recursive_repr()
    def __repr__(self):
        return '<' + ' , '.join(map(repr, self)) + '>'

l = lister()
l.append(ascii('a'))
l.append(ascii('b'))
l.append(ascii('c'))
print(l)

> <a is 97 , b is 98 , c is 99>
```

or using a wrapper inside wrapper inside decorator

## `functools`'s decorators

`functools.lru_cache` - caching for LRU objects/calls

```python
def fib1(n):
    return n if n in (0, 1) else (fib1(n-1)+fib1(n-2))

@functools.lru_cache(maxsize=35)
def fib2(n):
    return n if n < 2 else (fib2(n-1)+fib2(n-2))

# no limits
@functools.lru_cache(maxsize=None)
def fib3(n):
    return n if n < 2 else (fib3(n-1)+fib3(n-2))

%timeit fib1(12)
%timeit fib2(12)
%timeit fib3(12)

# cache info
ic(fib2.cache_info())
ic(fib3.cache_info())

> 57.8 µs ± 788 ns per loop (mean ± std. dev. of 7 runs, 10000 loops each)
> 88.3 ns ± 6.6 ns per loop (mean ± std. dev. of 7 runs, 10000000 loops each)
> 74.7 ns ± 5.27 ns per loop (mean ± std. dev. of 7 runs, 10000000 loops each)
> ic> fib2.cache_info(): CacheInfo(hits=81111120, misses=13, maxsize=35, currsize=13)
> ic> fib3.cache_info(): CacheInfo(hits=81111120, misses=13, maxsize=None, currsize=13)

result >>> CacheInfo(hits=81111120, misses=13, maxsize=None, currsize=13)
```

#### `functools.wraps` - allows to mask wrapper function name

- [What does `functools.wraps` do?](https://stackoverflow.com/questions/308999)

```python
def loud(func):
    def wrapper(*args, **kwargs):
        args = map(lambda x: x.upper(), (args))
        return func(*args, **kwargs)
    return wrapper

@loud
def mprint(string):
    return string


ic(mprint.__name__)
ic(mprint("yo"))

# -------------------------------------------------
del mprint

# -------------------------------------------------
def wisper(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        args = map(lambda x: x.lower(), (args))
        return func(*args, **kwargs)
    return wrapper

@wisper
def mprint(string):
    return string

ic(mprint.__name__)
ic(mprint("YO"))

> ic> mprint.__name__: 'wrapper'
> ic> mprint("yo"): 'YO'
> ic> mprint.__name__: 'mprint'
> ic> mprint("YO"): 'yo'

result >>> 'yo'
```

#### `functools.total_ordering`

Given a class defining one or more rich comparison ordering methods, this class decorator supplies the rest. This simplifies the effort involved in specifying all of the possible rich comparison operations:

The class must define one of `__lt__()`, `__le__()`, `__gt__()`, or `__ge__()`. In addition, the class should supply an `__eq__()` method.

```python
@functools.total_ordering
class Weight:
    def __init__(self, weight:str):
        self.number, self.units = weight.split(" ")

    @property
    def value(self):
        return self.converts_to_kilograms()

    def converts_to_kilograms(self):
        mesurments = {}
        mesurments['gr'] = 0.001
        mesurments['kg'] = 1
        mesurments['quintal'] = 100
        mesurments['pound'] = 0.5
        mesurments['ton'] = 1000

        return float(self.number)*mesurments.get(self.units)

    def __gt__(self, other):
        return self.value > other.value

    def __eq__(self, other):
        return self.value == other.value

ic(Weight('5 kg') == Weight('5000 gr'))
ic(Weight('5 kg') <  Weight('4999 gr'))
ic(Weight('5 kg') >  Weight('4999 gr'))

# clean output
None

> ic> Weight('5 kg') == Weight('5000 gr'): True
> ic> Weight('5 kg') <  Weight('4999 gr'): False
> ic> Weight('5 kg') >  Weight('4999 gr'): True
```

### `functools.singledispatch`

```python
from typing import List

list_of_ints = List[int]

@functools.singledispatch
def fun(arg, verbose=False):
    if verbose:
        print("Let me just say,", end=" ")
    print(arg)

# and overloading fucntions
@fun.register
def _(arg: int, verbose=False):
    if verbose:
        print("Strength in numbers, eh?", end=" ")
    print(arg)

@fun.register(list)
def _(arg, verbose=False):
    if verbose:
        print("Enumerate this:")
    for i, elem in enumerate(arg):
        print(i, elem)

import collections
@fun.register(collections.abc.Sequence)
def _(arg, verbose=False):
    if verbose:
        print("Sequence:")
    for i, elem in enumerate(arg):
        print(i, elem)
```

```python
fun("its not so funny", True)

> Sequence:
> 0 i
> 1 t
> 2 s
> 3
> 4 n
> 5 o
> 6 t
> 7
> 8 s
> 9 o
> 10
> 11 f
> 12 u
> 13 n
> 14 n
> 15 y
```

```python
fun(12312, True)

> Strength in numbers, eh? 12312
```

```python
fun([123,12,3,123,12,3,12,3], True)

> Enumerate this:
> 0 123
> 1 12
> 2 3
> 3 123
> 4 12
> 5 3
> 6 12
> 7 3
```

```python
# list available implementations
ic(fun.registry.keys())

# --- checking implementations....

# default implementation
ic(fun.dispatch(str))
ic(fun.dispatch(float))
# existing implementation
ic(fun.dispatch(int))

> ic> fun.registry.keys(): dict_keys([<class 'object'>, <class 'int'>, <class 'list'>, <class 'collections.abc.Sequence'>])
> ic> fun.dispatch(str): <function _ at 0x7ffb243bf1f0>
> ic> fun.dispatch(float): <function fun at 0x7ffb242f5ee0>
> ic> fun.dispatch(int): <function _ at 0x7ffb243bf550>

result >>> <function __main__._(arg: int, verbose=False)>
```

### `@decorators` used in Object-Oriented Programming

- `@statickmethod` - static methods
- `@classmethod` - class creation
- `@abstractmethod` - static methods

Note about `@abstractmethod`, in most cases you will use method that raises `NotImplemented` exception **only**^ which is not `abstract method`!

```python
from abc import abstractmethod

class AbstractPower:
    @abstractmethod
    def power(number, power):
        raise NotImplementedError("ddd")

class Power(AbstractPower):

    def __init__(self, number, power):
        self.number = number
        self.power = power

    @classmethod
    def cube(cls, number):
        return cls(number, 4)

    @classmethod
    def root(cls, number):
        return cls(number, 0.5)

    @staticmethod
    def power(x, power):
        return x**power

    def __repr__(self):
        return "{}({}^{}) is {}".format(
            self.__class__.__name__,
            self.number,
            self.power,
            self.__class__.power(self.number, self.power)
        )

ic(Power(2, 2))
ic(Power.cube(2))
ic(Power.root(4))
ic(Power.cube(2))

> ic> Power(2, 2): Power(2^2) is 4
> ic> Power.cube(2): Power(2^4) is 16
> ic> Power.root(4): Power(4^0.5) is 2.0
> ic> Power.cube(2): Power(2^4) is 16

result >>> Power(2^4) is 16
```

### Getters and Setters

```python
class Root:
    def __init__(self, number):
        self._n = number

    @property
    def n(self):
        return self._n

    @n.setter
    def n(self, n):
        self._n = n

    @property
    def root(self):
        return self._n * .5

root = Root(4)
ic(root)
ic(root.n)
ic(root.root)
root.n=16
ic(root.n)
ic(root.root)

> ic> root: <__main__.Root object at 0x7ffb2425ca30>
> ic> root.n: 4
> ic> root.root: 2.0
> ic> root.n: 16
> ic> root.root: 8.0

result >>> 8.0
```

### `@dataclass`

```python
import dataclasses

@dataclasses.dataclass
class Point:
    x: float
    y: float
    z: float = 0.0

ic(Point(1.5, 2.5))

> ic> Point(1.5, 2.5): Point(x=1.5, y=2.5, z=0.0)

result >>> Point(x=1.5, y=2.5, z=0.0)
```

### Context Managers and `@contextmanager`

You can make your own context manager without `__enter__` and `__exit__`

```python
from contextlib import contextmanager

@contextmanager
def managed_file(name):
    try:
        f = open(name, 'w')
        yield f
    finally:
        f.close()

with managed_file('hello.txt') as f:
    f.write('¡hola!')
```

### Exit function

```python
import atexit

@atexit.register
def goodbye():
    print("Ciao...")

exit(0)

> Ciao...
```
