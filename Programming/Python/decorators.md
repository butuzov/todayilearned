

# Decorators

[PEP 3129](https://www.python.org/dev/peps/pep-3129/), [PEP 318](https://www.python.org/dev/peps/pep-0318/)

### Resources
* [PythonWiki: PythonDecorators](https://wiki.python.org/moin/PythonDecorators)
* [PythonWiki: PythonDecoratorLibrary](https://wiki.python.org/moin/PythonDecoratorLibrary)
* [Stackoverflow: How to make a chain of function decorators?
](https://stackoverflow.com/questions/739654/)
* https://github.com/lord63/awesome-python-decorator

### Talks
* [PyCon2019: Practical decorators](https://youtu.be/MjHpMCIvwsY) by Reuven M. Lerner - [slides](https://speakerdeck.com/pycon2019/reuven-m-lerner-practical-decorators) 
* [EuroPytop 2018: A Taxonomy of Decorators: A-E](https://youtu.be/pEL1THG6ysY) by Andy Fundinger - [slides](https://github.com/Ciemaar/decorator-taxonomy)

### Example

```python
def decorator(func):
    def wrapper(*args, **kwargs):
        return f"{func(*args, **kwargs)}"
    
    return wrapper

@decorator
def add(a, b):
    return a+b
print(add(1,2))
output > '3'
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
print(pow(5))
output > 100
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
print(name('oleg'))
output > 'My name is slim shady'
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

%timeit fib1(35)
%timeit fib2(35)

# cache info
print(fib2.cache_info())
output > CacheInfo(hits=81111143, misses=36, maxsize=35, currsize=35)
```

    6.23 s ± 22.3 ms per loop (mean ± std. dev. of 7 runs, 1 loop each)
    117 ns ± 0.485 ns per loop (mean ± std. dev. of 7 runs, 10000000 loops each)

####  `functools.wraps` - allows to mask wrapper function name

* [What does `functools.wraps` do?](https://stackoverflow.com/questions/308999)

```python
def loud(func): 
    def wrapper(*args, **kwargs):
        args = map(lambda x: x.upper(), (args))
        return func(*args, **kwargs)
    return wrapper

@loud
def mprint(string):
    return string
print(mprint.__name__)
output > 'wrapper'
print(mprint.__name__)
output > 'mprint'
print(mprint("yo"))
output > 'YO'

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
print(mprint.__name__)
output > 'mprint'
print(mprint("YO"))
output > 'yo'
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

print(Weight('5 kg') < Weight('4999 gr'))
output > False
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
```

```python
fun("its not so funny", True)
```

    Let me just say, its not so funny

```python
fun(12312, True)
```

    Strength in numbers, eh? 12312

```python
fun([123,12,3,123,12,3,12,3], True)
```

    Enumerate this:
    0 123
    1 12
    2 3
    3 123
  12
    5 3
    6 12
    7 3

```python
# list available implementations
print(fun.registry.keys())
output > dict_keys([<class 'object'>, <class 'int'>, <class 'list'>])

# --- checking implementations....

# default implementation
print(fun.dispatch(str))
output > <function fun at 0x10f2766a8>
print(fun.dispatch(float))
output > <function fun at 0x10f2766a8>
# existing implementation
print(fun.dispatch(int))
output > <function _ at 0x10f276378>
```

    <function __main__._(arg: int, verbose=False)>

### `@decorators` used in Object-Oriented Programming

* `@statickmethod` - static methods
* `@classmethod` - class creation
* `@abstractmethod` - static methods 

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
    print(Power(2, 2))
    output > Power(2^2) is 4
print(Power.cube(2))
output > Power(2^4) is 16
print(Power.cube(2))
output > Power(2^4) is 16
print(Power.root(4))
output > Power(4^0.5) is 2.0 
print(Power.cube(2))
output > Power(2^4) is 16
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
print(root)
output > <__main__.Root object at 0x10f1ef630>
print(root.n)
output > 4
print(root.n)
output > 16
print(root.root)
output > 2.0
print(root.root)
output > 8.0
root.n=16
print(root.n)
output > 16
print(root.root)
output > 8.0
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
output >  Ciao...
```