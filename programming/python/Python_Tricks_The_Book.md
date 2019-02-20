# [Python Tricks: The Book](https://dbader.org/products/python-tricks-book/) by @dbader

My notes on the wonderful book, so I can return later to notes. Check link in title to buy it. It worth every penny.

-----

  * [Chapter 2 - Patterns for Cleaner Python](#chapter-2---patterns-for-cleaner-python)
      * [2.1 Covering Your A** With Assertions](#21-covering-your-a-with-assertions)
      * [2.2 Complacent Comma Placement](#22-complacent-comma-placement)
      * [2.3 Context Managers and the with Statement](#23-context-managers-and-the-with-statement)
      * [2.4 Underscores, Dunders, and More](#24-underscores-dunders-and-more)
      * [2.5 A Shocking Truth About String Formatting](#25-a-shocking-truth-about-string-formatting)
      * [2.6 “The Zen of Python” Easter Egg](#26-the-zen-of-python-easter-egg)
  * [Chapter 3 - Effective Functions](#chapter-3---effective-functions)
      * [3.1 Python’s Functions Are First-Class](#31-pythons-functions-are-first-class)
      * [3.2 Lambdas Are Single-Expression Functions](#32-lambdas-are-single-expression-functions)
      * [3.3 The Power of Decorators](#33-the-power-of-decorators)
      * [3.4 Fun With `*args` and <code>**kwargs</code>](#34-fun-with-args-and-kwargs)
      * [3.5 Function Argument Unpacking](#35-function-argument-unpacking)
      * [3.6 Nothing to Return Here](#36-nothing-to-return-here)
  * [Chapter 4 - Classes &amp; OOP](#chapter-4---classes--oop)
      * [4.1 Object Comparisons: `is` vs `==`](#41-object-comparisons-is-vs-)
      * [4.2 String Conversion (Every Class Needs a `__repr__`)](#42-string-conversion-every-class-needs-a-__repr__)
      * [4.3 Defining Your Own Exception Classes](#43-defining-your-own-exception-classes)
      * [4.4 Cloning Objects for Fun and Profit](#44-cloning-objects-for-fun-and-profit)
      * [4.5 Abstract Base Classes Keep Inheritance in Check](#45-abstract-base-classes-keep-inheritance-in-check)
      * [4.6 What Named tuples Are Good For](#46-what-named-tuples-are-good-for)
      * [4.7 Class vs Instance Variable Pitfalls](#47-class-vs-instance-variable-pitfalls)
      * [4.8 Instance, Class, and Static Methods Demystified](#48-instance-class-and-static-methods-demystified)
  * [Chapter 5 - Common Data Structures in Python](#chapter-5---common-data-structures-in-python)
      * [5.1 Dictionaries, Maps, and Hashtables](#51-dictionaries-maps-and-hashtables)
      * [5.2 Array Data Structures](#52-array-data-structures)
      * [5.3 Records, Structs, and Data Transfer Objects](#53-records-structs-and-data-transfer-objects)
      * [5.4 Sets and Multi sets.](#54-sets-and-multi-sets)
      * [5.5 Stacks (LIFOs)](#55-stacks-lifos)
      * [5.6 Queues (FIFOs)](#56-queues-fifos)
      * [5.7 PriorityQueues](#57-priorityqueues)
  * [Chapter 6 - Looping &amp; Iteration](#chapter-6---looping--iteration)
      * [6.1 Writing Pythonic Loops](#61-writing-pythonic-loops)
      * [6.2 Comprehending Comprehensions](#62-comprehending-comprehensions)
      * [6.3 List Slicing Tricks and the Sushi Operator](#63-list-slicing-tricks-and-the-sushi-operator)
      * [6.4 Beautiful Iterators](#64-beautiful-iterators)
      * [6.5 Generators Are Simplified Iterators](#65-generators-are-simplified-iterators)
      * [6.6 Generator Expressions](#66-generator-expressions)
      * [6.7 Iterator Chains](#67-iterator-chains)
  * [Chapter 7 - Dictionary Tricks](#chapter-7---dictionary-tricks)
      * [7.1 Dictionary Default Values](#71-dictionary-default-values)
      * [7.2 Sorting Dictionaries for Fun and Profit](#72-sorting-dictionaries-for-fun-and-profit)
      * [7.3 Emulating Switch/Case Statements With Dicts](#73-emulating-switchcase-statements-with-dicts)
      * [7.4 The Craziest Dict Expression in the West](#74-the-craziest-dict-expression-in-the-west)
      * [7.5 So Many Ways to Merge Dictionaries](#75-so-many-ways-to-merge-dictionaries)
      * [7.6 Dictionary Pretty-Printing](#76-dictionary-pretty-printing)
  * [Chapter 8 - Pythonic Productivity Techniques](#chapter-8---pythonic-productivity-techniques)
      * [8.1 Exploring Python Modules and Objects](#81-exploring-python-modules-and-objects)
      * [8.2 Isolating Project Dependencies With Virtualenv](#82-isolating-project-dependencies-with-virtualenv)
      * [8.3 Peeking Behind the Bytecode Curtain](#83-peeking-behind-the-bytecode-curtain)
  * [Chapter 9/10 - Closing Thoughts and ...](#chapter-910---additional-things-dan-pointed-but-didnt-mention)
      * [9/10.0 Podcasts, Websites etc.](#9100-podcasts-websites-etc)
      * [9/10.1 Tenary Expression.](#9101-tenary-expression)



## Chapter 2 - Patterns for Cleaner Python

### 2.1 Covering Your A** With Assertions

1. Never use tuples in assert.
```python
assert( 1 == 2, "It's failing!") # No it's not
```

2. Do not give assert runtime checks, `assert` can be disabled.

### 2.2 Complacent Comma Placement

```python
# Having comma at the end of your list is OK.
names = [
  'John',
  'Betty',
  'Liza',
]

#  not having a comma - in strings concatenation.
str_var = (
  'A long time ago'
  'in a galaxy far, far away....'
)
```


### 2.3 Context Managers and the with Statement

#### `with ` & dunders
```python
class ManagedFile:
     def __init__(self, name):
         self.name = name
    def __enter__(self):
         self.file = open(self.name, 'w')
         return self.file
    def __exit__(self, exc_type, exc_val, exc_tb):
         if self.file:
             self.file.close()

# so we can do something like
# implementing custom write or read
or any other functionality.
with ManagedFile("list.json") as f:
    f.read()

```

#### `with` & `contextlib`
```python
from contextlib import contextmanager
@contextmanager
def managed_file(name):
  try:
    f = open(name, 'w')
    yield f
  finally:
    f.close()

>>> with managed_file('hello.txt') as f:
 ...     f.write('привет, мир!')
 ...     f.write('а теперь, пока!')

```

### 2.4 Underscores, Dunders, and More

  Type                  | Example | Why
------------------------|---------|-----
Single@Start            | `_var`  | This is private methods/functions/vars
Single@End              |  `var_` | Avoid conflicts names
Double@Start            | `__var` | Name mangling
Dunder                  |`__var__`| Magic
Single                  | `_`     | Ignore_me_variable or Last expression result.


### 2.5 A Shocking Truth About String Formatting

```python
name = "Django"
# old way
print( "Hey %s" % name )        # -> Hey Django
# using format
print( "Hey {}".format( name )) # -> Hey Django
# using f-strings
print( f"Hey {name}" )          # -> Hey Django
# using templates
from string import Template
print( Template('Hey, $name').substitute(name=name) ) #-> Hey, Django
```
Read more about format and % ways @ [Pyformat](https://pyformat.info/)


### 2.6 “[The Zen of Python](http://www.thezenofpython.com/)” Easter Egg

```python
import this
```

#### The Zen of Python, by Tim Peters

  * Beautiful is better than ugly.
  * Explicit is better than implicit.
  * Simple is better than complex.
  * Complex is better than complicated.
  * Flat is better than nested.
  * Sparse is better than dense.
  * Readability counts.
  * Special cases aren't special enough to break the rules.
  * Although practicality beats purity.
  * Errors should never pass silently.
  * Unless explicitly silenced.
  * In the face of ambiguity, refuse the temptation to guess.
  * There should be one-- and preferably only one --obvious way to do it.
  * Although that way may not be obvious at first unless you're Dutch.
  * Now is better than never.
  * Although never is often better than *right* now.
  * If the implementation is hard to explain, it's a bad idea.
  * If the implementation is easy to explain, it may be a good idea.
  * Namespaces are one honking great idea -- let's do more of those!


## Chapter 3 - Effective Functions

### 3.1 Python’s Functions Are First-Class

#### Functions are objects
```python
def yell(test):
  return test.toupper() + '!'
bark = yell

bark("Bark-Bark")
del yell
yell('Hello?') # Raise NameError

bark("hey")
print(bark.__name__) # yell

# list of functions
[ str.lower, str.upper ]

def greet(func_):
  return func_("¡Hola!")

print(greet(str.upper))

# nested functions (1st class citizens)
# we can use it, abstract it or return.
def speak(text):
  def loud(t):
    return t.upper()
  return loud(text)

# state and context
# grab the local scope
def get_speak_func(text, volume):
  def whisper():
    return text.lower() + '...'
  def yell():
    return text.upper() + '!'
  if volume > 0.5:
    return yell
  else:
return whisper
>>> get_speak_func('Привет, Мир', 0.7)()
```

#### Objects as functions

Just add __call__ to it.

```python
class Adder:
  def __init__(self, n):
    self.n = n
  def __call__(self, x):
    return self.n + x
>>> plus_3 = Adder(3)
>>> plus_3(4)
```

Is object `callable`? Use `callable`!

```python
callable("some text")
>>> False
callable(Adder)
>>> True
```

### 3.2 Lambdas Are Single-Expression Functions

```python
add = lambda x, y: x + y
add(1,2)
>>> 3
```
#### Functional expression

```
( lambda x, y: x+y )( 3, 4 )
>>>7
```

We can use lambdas with:

  * `sorted`
  * `map`
  * `reduce`
  * `filter`

Lexical Closure

```python
def summe(n):
  return lambda x: x + n

summe(1)(2)
>>> 3
```
#### THINK TWICE BEFORE WRITING LAMBDA - ITS CAN BE HARD TO READ YOUR CODE.


### 3.3 The Power of Decorators

#### Proposes
 * Protocoling decorators
 * Authentefications and Access Control
 * API rate limiting
 * caching/memoization
 * Adapters/Wraper

```python
@dec_1
@dec_2
def method():
  return "Yo!"
>> dec_1( dec_2( method( ) ) )

def decorator(func):
  def wraper(*args, **kwargs):
    # some incredible logic here
    return func(*args, **kwargs)
  return wraper

@decorator
def demo(say, my_name):
  return say(my_name)
# or

@decorator demo(str.upper, "Oleg")
```

#### debug
```python
# allow us to trace with __name__ and __doc__
import functools
def uppercase(func):
  @functools.wraps(func)
  # will move to wraper name and doc string
  def wrapper():
    return func().upper()
  return wrapper
```


### 3.4 Fun With `*args` and `**kwargs`

Non Required Arguments

  * `*args` - positional arguments
  * `**kwargs` - named arguments

### 3.5 Function Argument Unpacking

We can unpack lists/sets/tuples into function parameters

```python
def func(one, two):
  return ( f"I am count: {one} - you hide,"
            "I am count: {two} - you run.")

print( func(*["one", "two"]) )
```

### 3.6 Nothing to Return Here
```python
# return None!

def foo():
  return

def bar():
  pass

def barfoo():
  if global_var == 1:
    return True
  # here comes hidden return None

def barfoo():
  return None
```

## Chapter 4 - Classes & OOP

### 4.1 Object Comparisons: `is` vs `==`

  * `==` equivalence check
  * `is` identity check

### 4.2 String Conversion (Every Class Needs a `__repr__`)

Class to string representation with dunder `repr` (`__repr__`), simular to
`toString` (php, js/node.js, java), `func (s Structure) String() string` (go).
Note: `__str__` can overide a string representation, but directly and in collections will be used `__repr__`. Build in `str` and `repr` function will work only with `__str__` and `__repr__` if available.

```python
class SomeClass():
  def __repr__(self):
    return (f'{self.__class__.__name__}('
      f'{self.attr1!r}, {self.attr2!r})')
```

 * `__str__` should be readable
 * `__repr__` should be defined

### 4.3 Defining Your Own Exception Classes

```python
class NameTooShortError(ValueError):
  pass

def validate(name):
  if len(name) < 10:
    raise NameTooShortError(name)
```

### 4.4 Cloning Objects for Fun and Profit

Deep and shallow copies.

```python
import copy
xs = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
zs = copy.deepcopy(xs)  # deep copy
ys = copy.copy(xs)      # shallow copy
```

In addition `__copy__` and `__deepcopy__` can cover copy and shallow copy custom logic.


### 4.5 Abstract Base Classes Keep Inheritance in Check

```python
from abc import ABCMeta, abstractmethod

class Base(metaclass=ABCMeta):
  @abstractmethod
  def foo(self):
    pass
  @abstractmethod
  def bar(self):
    pass

# fails as bar not implemented
class Concrete(Base):
  def foo(self):
    pass

# correctly display inheritance
assert issubclass(Concrete, Base)
```

### 4.6 What Named tuples Are Good For

```python
from collections import namedtuple
Car = namedtuple('Auto' , 'color engine')
# Is Creates class Auto with a fields 'color' and 'engine' and
# instanciate it to Car variable
redSubaru = Car("red", "3.6Turbo")
Car._fields
>>> ('color', 'engine')
# and you can instantiate class with wraper implementing
# custom methods
class CarMethods(Car):
  def hex(self):
    return '#000' if self.color not "red" else "#f00"

# replace value
redSubaru._repalce(color=black)

# dumping values as json from OrderDict
json.dump(redSubaru._asdict())

# and finaly - factory
Car._make("red", "turbo")

```

### 4.7 Class vs Instance Variable Pitfalls

 * Class Variables - internal class variables - e.g. `cls` or simple `name` (static variables)
 * Instance Variables - variables stored in object `self`

```python
class Dog:
  num_legs = 4 # <- Переменная класса
  def __init__(self, name):
    self.name = name # <- Переменная экземпляра
```

Here goes example with a  jack dog, and modification num_legs directly, which affects also all instances. But -  it's allowed to change class variables in instances.

```python
jack.num_legs = 6
jack.num_legs, jack.__class__.num_legs
>>> (6, 4)
```

### 4.8 Instance, Class, and Static Methods Demystified
```python
class MyClass:
  def method(self):
    return 'вызван метод экземпляра', self

  @classmethod
  def classmethod(cls):
    return 'вызван метод класса', cls

  @staticmethod
  def staticmethod():
      return 'вызван статический метод'
```

## Chapter 5 - Common Data Structures in Python

### 5.1 Dictionaries, Maps, and Hashtables
  * PyCon 2017 Talk: [Modern Python Dictionaries - A confluence of a dozen great ideas](https://www.youtube.com/watch?v=npw4s1QTmPg) by Raymond Hettinger
  * collections.OrderedDict
  * collections.defaultdict - having `defaults` in your dictionary.
  * collections .ChainMap - multiple dictionaries combined (for search proposes, for example)
  * types .MappingProxyType - read only dict (frozendict).
  * keys are `hashable`'s





## Chapter 6 - Looping & Iteration

### 6.1 Writing Pythonic Loops

  * `enumerate` and `.items()` for lists and dicts as exprected.

### 6.2 Comprehending Comprehensions

  * `{}`, `[]` - dictionary, set & list comprehensions.
  ```python
  # If condition
  if_ = [  x    for x  in range(100) if x % 10 == 0 ]
  # If else condition
  elif_ = [  x   if x % 10 == 0 else "!10" for x  in range(100) ]

  print( if_, elif_ )

  # list
  print([x for x  in range(100) if x % 10 == 0])

  # set
  print({x for x  in range(100) if x % 10 == 0})

  # dictionary
  print({x:x+1 for x  in range(100) if x % 10 == 0})
  ```


### 6.4 Beautiful Iterators

  For custom data structures/class collections iteration you need to implelemt `__iter__` and `__next__`. Raise `StopIteration` to stop.

### 6.5 Generators Are Simplified Iterators

  Memory optimized "returns" - `yeald value`, return `None` to stop.

### 6.6 Generator Expressions

  `()` - same sytdax as list/set/dictionalry comprehension. But, its `generator` expression.

### 6.7 Iterator Chains

  Generators can be chained in the same way as the function can be chained.

  ```python
  # its more readable way, but it's can be inlined.
  print(
      list(
          (-i for i in
              (i*i for i in
                  (i for i in range(10) if i % 3 == 0 )
              )
          )
      )
  )
  ```

## Chapter 7 - Dictionary Tricks
### 7.1 Dictionary Default Values

  ```python
  # using .get() method
  dictionary.get('key', []) # list is value if key entry not found

  # using default dict zero value, wouldn't work with .get()
  from collections import defaultdict
  print(defaultdict(int)['key'])
  ```


### 7.3 Emulating Switch/Case Statements With Dicts

  You can use lists or dictionaries of callbacks/functions to emulate `switch`.

### 7.4 The Craziest Dict Expression in the West

  Avoid using: `True`, 1, 1.0 in same time as keys. Implement you `__hash__` and/or `__eq__` nicly.

### 7.5 So Many Ways to Merge Dictionaries

  Pre-Python 3.5
  ```python
  zs = dict(xs, **ys)
  >>> zs
  {'a': 1, 'c': 4, 'b': 3}
  ```

  Post-Python 3.5

  ```python
  zs = {**xs, **ys}
  ```

### 7.6 Dictionary Pretty-Printing

  ```python
  mapping = {'a': 23, 'b': 42, 'c': 0xc0ffee}

  # simple debug dump using json (file pointer not required)
  import json, sys

  json.dump(mapping, fp=sys.stdout, indent=4, sort_keys=True)

  # using icecream
  # https://github.com/gruns/icecream
  from icecream import ic

  ic(mapping)

  # using pretty print
  import pprint

  pprint.pprint(mapping)
  ```

## Chapter 8 - Pythonic Productivity Techniques
### 8.1 Exploring Python Modules and Objects

  Use REPL more offen. `dir()`, `help()` for simple reflection.

### 8.2 Isolating Project Dependencies With Virtualenv

  `python3 -m venv .venv` by Raimond and Dan.

  ```bash
  # create virtual environment
  $ > python3 -m venv .venv
  # activate it, instead of . you also can use source
  $ > . .venv/bin/activate
  # install whatever you need
  (.venv) $ > python3 -m pip install pandas
  # want to stop work? deactivate it
  (.venv) $ > deactivate
  # and back to your shell
  $ >
  ```

### 8.3 Peeking Behind the Bytecode Curtain
  ```python
  import dis

  def test(name):
      return f"Hello {name}!"

  print(test.__code__.co_code)
  print(test.__code__.co_consts)
  print(test.__code__.co_varnames)

  dis.dis(test)
  ```
## Chapter 9/10 - Closing Thoughts and ...

### 9/10.0 Podcasts, Websites etc.

### 9/10.1 Tenary Expression.

  ```python
  # as we discovered sets are faster, but its ok to use two elements list.
  # ( false_result, true_result )[ bool expression ]
  ('NOPE', 'YES')[ True ]
  # or something like this.
  (1, 2)[ 2 > 1]
  ```
