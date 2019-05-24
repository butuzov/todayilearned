

# Python Data Structures & Types

## Data Types

### int

Supported Methods:
`__abs__`, `__add__`, `__and__`, `__bool__`, `__ceil__`, `__class__`, `__delattr__`, `__dir__`, `__divmod__`, `__doc__`, `__eq__`, `__float__`, `__floor__`, `__floordiv__`, `__format__`, `__ge__`, `__getattribute__`, `__getnewargs__`, `__gt__`, `__hash__`, `__index__`, `__init__`, `__init_subclass__`, `__int__`, `__invert__`, `__le__`, `__lshift__`, `__lt__`, `__mod__`, `__mul__`, `__ne__`, `__neg__`, `__new__`, `__or__`, `__pos__`, `__pow__`, `__radd__`, `__rand__`, `__rdivmod__`, `__reduce__`, `__reduce_ex__`, `__repr__`, `__rfloordiv__`, `__rlshift__`, `__rmod__`, `__rmul__`, `__ror__`, `__round__`, `__rpow__`, `__rrshift__`, `__rshift__`, `__rsub__`, `__rtruediv__`, `__rxor__`, `__setattr__`, `__sizeof__`, `__str__`, `__sub__`, `__subclasshook__`, `__truediv__`, `__trunc__`, `__xor__`, `bit_length`, `conjugate`, `denominator`, `from_bytes`, `imag`, `numerator`, `real`, `to_bytes`

```python
i: int = 1
    print(sys.getsizeof(i))
    output > 28
    print(sys.getsizeof(i))
    output > 28
    print(sys.getsizeof(i))
    output > 36
print(i.bit_length())
output > 1

# number separation with underscores
i: int = 1_000_000
    print(sys.getsizeof(i))
    output > 28
    print(sys.getsizeof(i))
    output > 36

# big int
i: int = 2**64
    print(sys.getsizeof(i))
    output > 36
```

```python
### float
```

```python
### str
```

```python
### bytes
```

### `struct.Struct`

structured data as bytes 

```python
# Struct
# https://docs.python.org/3/library/struct.html
from struct import Struct

MyStruct = Struct('i?f')
compact_data = MyStruct.pack(23, False, 42.0)
print(compact_data)
output > b'\x17\x00\x00\x00\x00\x00\x00\x00\x00\x00(B'
# and unpacking
print(MyStruct.unpack(compact_data))
output > (23, False, 42.0)
```

## Dictionary and Maps

### `dict`

**Talks**:

* [Modern Python Dictionaries - A confluence of a dozen great ideas (PyCon 2017)](https://www.youtube.com/watch?v=npw4s1QTmPg) by Raymond Hettinger

```python
ys={'z':9, 'y':1, 'x':5}
xs={'a':1, 'b':3, 'c':2}
print(xs)
output > {'a': 1, 'b': 3, 'c': 2}

# Merging Dictionaries
# pre python 3.5
print(dict(xs, **ys))
output > {'a': 1, 'b': 3, 'c': 2, 'x': 5, 'y': 1, 'z': 9}

# post python 3.5
print({**xs, **ys})
output > {'a': 1, 'b': 3, 'c': 2, 'x': 5, 'y': 1, 'z': 9}
```

    {'a': 1, 'b': 3, 'c': 2, 'z': 9, 'y': 1, 'x': 5}

### `collections.OrderedDict`

```python
# Ordered Dict

# in python 3.7 order will be preserved, 
# but it's depends on dict implementation
# see Hettinger's video.
a = {'a':1, 'b':3, 'c':2}
print(a)
output > {'a': 1, 'b': 3, 'c': 2}

# ordered dict helps to keep it really ordered
from collections import OrderedDict
b = OrderedDict(a)
print(b)
output > OrderedDict([('a', 1), ('b', 3), ('c', 2)])
print(b)
output > b'literal'
print(b)
output > bytearray(b'Dots over \xd1\x96.I said put dots over i.')
print(b)
output > bytearray(b'Comas over \xd1\x96.I said put dots over i.')

# we also can use sorted to sort doctionary by value
s1 = sorted(a.items(), key=lambda x: x[1])
print(s1)
output > [('a', 1), ('c', 2), ('b', 3)]

# or by key
s2 = sorted(a.items(), key=lambda x: x[0])
print(s2)
output > [('a', 1), ('b', 3), ('c', 2)]

#  or usign comprehension
s3 = {k:v for k, v in sorted(a.items(), key=lambda x: x[1], reverse=False)}
print(s3)
output > {'a': 1, 'b': 3, 'c': 2}
```

    {'a': 1, 'c': 2, 'b': 3}

### `collections.ChainMap`

```python
from collections import ChainMap
from collections import defaultdict as DefaultDict

# default options
default_options = dict({
    'option':'default_value'
})

# options from default dict
custom_source = DefaultDict(lambda:"default_dictionaly_option_placeholder")
custom_source['option'] = 'default_dict_option'

options_source = ChainMap(default_options, custom_source)
print(options_source['option'])
output > 'default_value'
print(options_source['nosuchoptions'])
output > 'default_dictionaly_option_placeholder'

# merging dictionaries
d1 = {'a': 1, 'b': 2}
d2 = {'c': 3, 'd': 4}
 print({**ChainMap(d1, d2)})
 output > {'a': 1, 'b': 2, 'c': 3, 'd': 4}
```

    {'c': 3, 'd': 4, 'a': 1, 'b': 2}

### `collections.defaultdict`

```python
import collections 

my_dict = {}

# default value (for value that not exists) is list
print(my_dict.get('key_that_not_exists', []))
output > []

# with default dictionary
print(collections.defaultdict(list)['key_that_not_exists'])
output > []
```

### `types.MappingProxyType`

* https://www.python.org/dev/peps/pep-0416
* https://docs.python.org/3/library/types.html

```python
from types import MappingProxyType

Colors = MappingProxyType({
    'red'  : '#FF0000',
    'green': '#008000',
})

# show colors
print(Colors)
output > mappingproxy({'red': '#FF0000', 'green': '#008000'})

# unexisting key
print(Colors.get('navy'))
output > None

# updating keys (shoudl fail)
error_message = "Cant Update frozen/proxy value of dictionary"
try:
    Colors.update({'navy':'#000080'})
except AttributeError:
    print(error_message)
    output > 'Cant Update frozen/proxy value of dictionary'
```

## Lists and Arrays

### `bytes`

`immutable` bytes sequence

```python
b = b'literal'
# bytes literal - restricted to ascii symbols (except \ and control codes)
print(b)
output > b'literal'
print(b)
output > bytearray(b'Dots over \xd1\x96.I said put dots over i.')
print(b)
output > bytearray(b'Comas over \xd1\x96.I said put dots over i.')

# int (code at ascii) if indexing 
print(b[1])
output > 105

# int (code at ascii) if indexing 
print(b[0:3])
output > b'lit'

# bytes to string convertion (with escape characters)
print(b'dots over \xd1\x96'.decode())
output > 'dots over і'

# creating sequence qith zero value
print(bytes(10))
output > b'\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'

# if initializer sequence of integers ...
print(bytes(range(97, 97+26)))
output > b'abcdefghijklmnopqrstuvwxyz'

# or sequence of utf8 aswell.
print(bytes("dots over і", 'utf8'))
output > b'dots over \xd1\x96'
```

### `array.array`

`C`- types arrays

```python
# https://docs.python.org/3/library/array.html
from array import array 

# Unicode (Py_UNICODE)
print(array('u', 'hello \u2641'))
output > array('u', 'hello ♁')

# signed long
print(array('l', (1, 2, 3, 4, 5)))
output > array('l', [1, 2, 3, 4, 5])

# doubles
print(array('d', (1.0, 2.0, 3.14)))
output > array('d', [1.0, 2.0, 3.14])

# floats 
print(array('f', (1.0, 1.5, 2.0, 2.5)))
output > array('f', [1.0, 1.5, 2.0, 2.5])
```

### `bytearray`

`mutable` bytes seqence

```python
# Same way to initialize it was done with byte
print(bytearray())
output > bytearray(b'')
print(bytearray(4))
output > bytearray(b'\x00\x00\x00\x00')
print(bytearray("Dots over і.", 'utf8'))
output > bytearray(b'Dots over \xd1\x96.')

# Mutating array
b = bytearray("Dots over і.", 'utf8')
b.extend(bytes('I said put dots over i.', 'utf8'))
print(b)
output > bytearray(b'Dots over \xd1\x96.I said put dots over i.')
print(b)
output > bytearray(b'Comas over \xd1\x96.I said put dots over i.')

b[0:5] = b'Comas '
print(b)
output > bytearray(b'Comas over \xd1\x96.I said put dots over i.')

# bytes supports strings like opperations
words = b"Dont comes easy"

# text transform
print(words.upper())
output > b'DONT COMES EASY'

# text split to list
print(words.split())
output > [b'Dont', b'comes', b'easy']

# and join it back
print(bytearray(b" ").join(words.split()))
output > bytearray(b'Dont comes easy')
```

## Linked Lists, Heaps and Queues

### `heapq`

```python
import heapq

# priority queue realization using heapqueue
class PriorityQueue(object):
    def __init__(self):
        self._q = []

    def add(self, value, priority=0):
        heapq.heappush(self._q, (priority, value))

    def pop(self):
        return heapq.heappop(self._q)[-1]
    
pq = PriorityQueue()
pq.add("hello", priority=2)
pq.add("world", priority=1)
print(pq.pop())
output > 'world'
print(pq.pop())
output > 'hello'
print(pq.pop())
output > 'hello'
```

```python
### `queue.PriorityQueue`
```

### `queue.deque`

```python
import collections 

d = collections.deque(["one", "two", "three"])
print(d)
output > deque(['one', 'two', 'three'])
print(d)
output > deque(['one', 'two', 'three', 'four'])
print(d)
output > deque(['one', 'two', 'four'])
print(d)
output > deque(['one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['zero', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try'])
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])

d.append("four")
print(d)
output > deque(['one', 'two', 'three', 'four'])
print(d)
output > deque(['one', 'two', 'four'])
print(d)
output > deque(['one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['zero', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try'])
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])
print(d.remove("three"))
output > None
print(d)
output > deque(['one', 'two', 'four'])
print(d)
output > deque(['one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['zero', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try'])
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])

d.extend(["dva", "try", "chotyry"])
print(d)
output > deque(['one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['zero', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try'])
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])

d.extendleft(["iti", "ni", "san", "si"])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['zero', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try'])
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])

d.appendleft("zero")
print(d)
output > deque(['zero', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try', 'chotyry'])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try'])
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])

d_poped = d.pop()
print(d_poped)
output > 'chotyry'

d_popedleft = d.popleft()
print(d_popedleft)
output > 'zero'

# rotatiob
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try'])
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])
d.rotate(2)
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])
```

## Sets

### `set`

```python
s = set([1,2,1,2,3,1])
print(s)
output > {1, 2, 3}
print(s)
output > {19, 1, 2, 3}

s.add(19)
s.add(19)
print(s)
output > {19, 1, 2, 3}

# intersection
print(s & set([2, 19]))
output > {2, 19}
```

```python
### `frozenset`
```

### `collections.Counter`

```python
# collections.Counter example

from collections import Counter

c0 = Counter({'uno':1, 'dos':2});
print(c0)
output > Counter({'dos': 2, 'uno': 1})

c1 = Counter(uno=1)
print(c1)
output > Counter({'uno': 1})

c2 = Counter();
c2['apples'] += 1
c2['apples'] += 2
c2['peaches'] += 10

# Most Common Elements (sorted by `count`)
print(c2.most_common())
output > [('peaches', 10), ('apples', 3)]

# Counter Example
c3 = Counter()
lipsum = "Un dos tres quatro cinco seis siete"
words = [w.strip(".,") for w in lipsum.lower().split(' ')]   
for word in words:
    c3[word]+=1
print(c3.most_common(2))
output > [('un', 1), ('dos', 1)]

# List
print(list(c3.elements()))
output > ['un', 'dos', 'tres', 'quatro', 'cinco', 'seis', 'siete']

# Unique Elements
print(list(c3))
output > ['un', 'dos', 'tres', 'quatro', 'cinco', 'seis', 'siete']

# Additoional methods
c4 = Counter(["dos"])
c4.update(["dos", "tres", "cinco", "seis", "ocho"])
c4.subtract(["ocho"])
print(c4)
output > Counter({'dos': 2, 'tres': 1, 'cinco': 1, 'seis': 1, 'ocho': 0})
print(c4)
output > Counter({'dos': 2, 'tres': 1, 'cinco': 1, 'seis': 1, 'ocho': 0}) 

# background implementation - itertools chain
print(c4.elements())
output > <itertools.chain object at 0x10954bda0>

# adding counters
print(c4+c3)
output > Counter({'dos': 3, 'tres': 2, 'cinco': 2, 'seis': 2, 'un': 1, 'quatro': 1, 'siete': 1})
print(c4)
output > Counter({'dos': 2, 'tres': 1, 'cinco': 1, 'seis': 1, 'ocho': 0})
print(c3)
output > Counter({'un': 1, 'dos': 1, 'tres': 1, 'quatro': 1, 'cinco': 1, 'seis': 1, 'siete': 1})
# subtraction
print(c4-c3)
output > Counter({'dos': 1})
print(c3-c4)
output > Counter({'un': 1, 'quatro': 1, 'siete': 1})

# maximum value (union)
print(c3|c4)
output > Counter({'dos': 2, 'un': 1, 'tres': 1, 'quatro': 1, 'cinco': 1, 'seis': 1, 'siete': 1})

# minimum value (intersection)
print(c3&c4)
output > Counter({'dos': 1, 'tres': 1, 'cinco': 1, 'seis': 1})
```

## Tuples

### tuple

`immutable` sequence (except if it contains `list`)

```python
one = "uno"
dos = "two"
t = (one, dos)

try:
    t[1] = "dva"
except TypeError as e:
    print("Error: ", e, file=sys.stderr)
```

    Error:  'tuple' object does not support item assignment

### `collections.namedtuple`

```python
import collections, json 

Car = collections.namedtuple('Auto' , 'color engine')
# Is Creates class Auto with a fields 'color' and 'engine' and
# instanciate it to Car variable
redSubaru = Car("red", "3.6Turbo")
print(Car._fields)
output > ('color', 'engine')

# And you can instantiate class with wraper implementing custom methods
class CarMethods(Car):
    def hex(self):
        return ('#000', "#f00")[self.color != "red"]

# replace value
redSubaru._replace(color="#000")
print(redSubaru)
output > Auto(color='red', engine='3.6Turbo')

# show tuple as dictionary
print(json.dumps(redSubaru._asdict()))
output > '{"color": "red", "engine": "3.6Turbo"}'

# and finaly - factory
red_turbo_car = Car._make(("red", "turbo"))
print(red_turbo_car)
output > Auto(color='red', engine='turbo')

# Conversions
print(tuple(red_turbo_car))
output > ('red', 'turbo')
print(red_turbo_car._asdict())
output > OrderedDict([('color', 'red'), ('engine', 'turbo')])
```

## enum.Enum

### `enum.IntEnum`

* https://docs.python.org/3/library/enum.html

```python
import enum

class Requests(enum.Enum):
    POST    = 1
    GET     = 2
    HEAD    = 3
    OPTIONS = 4
    DELETE  = 5
    PUSH    = 6

request = Requests.POST
print(request)
output > <Requests.POST: 1>
print(isinstance(request, Requests))
output > True
print(request.name)
output > 'POST' 

# Programmatic access to enumeration members and their attributes
print(Requests(2))
output > <Requests.GET: 2>

types = [str(request_type.name) for request_type in Requests]
print(", ".join(types))
output >     type object 'HTTPErrorCodes' has no attribute '__members__'

# ensuring unique values
try:
    @enum.unique
    class HTTPErrorCodes:
        NotFound = 404
        Forbiden = 404
except AttributeError  as e:  
    print(e)
    

# iota / automatic values
# 
class HTTPErrorCodes(enum.Enum):
    Bad_Request      = 400
    Unauthorized     = enum.auto()
    Payment_Required = enum.auto()
    Forbiden         = enum.auto()

codes = [str(codes.name) for codes in HTTPErrorCodes]
print(", ".join(codes))
output > 'Bad_Request, Unauthorized, Payment_Required, Forbiden'
print(", ".join(codes))
output > '400, 401, 402, 403'

values = [str(codes.value) for codes in HTTPErrorCodes]

# custumizing enum
class AutoName(enum.Enum):
    def _generate_next_value_(name, start, count, last_values):
        return name

class Ordinal(AutoName):
    NORTH = enum.auto()
    SOUTH = enum.auto()
    EAST = enum.auto()
    WEST = enum.auto()

directions = [str(direction.value) for direction in Ordinal]
print(", ".join(directions))
output > 'NORTH, SOUTH, EAST, WEST'

# Using API
Animals = enum.Enum('Animal', 'ANT BEE CAT DOG')
animals = [str(animal.name) for animal in Animals]
print(", ".join(animals))
output > 'ANT, BEE, CAT, DOG'

# using additional property
class Planets(enum.Enum):
    EARTH   = (5.976e+24, 6.37814e6)
    MARS    = (6.421e+23, 3.3972e6)
    def __init__(self, mass, radius):
        self.mass = mass       # in kilograms
        self.radius = radius   # in meters
        
    @property
    def gravity(self):
        # universal gravitational constant  (m3 kg-1 s-2)
        G = 6.67300E-11
        return G * self.mass / (self.radius * self.radius)
    
planets = [str((planet.name, planet.gravity)) for planet in Planets]
print(", ".join(planets))
output > "('EARTH', 9.802652743337129), ('MARS', 3.7126290961053403)"
```

### `enum.IntFlag`

```python
class Perm(enum.IntFlag):
    R = 4
    W = 2
    X = 1
    print(Perm.R | Perm.W)
    output > <Perm.R|W: 6>
print(Perm.R + Perm.W)
output > 6
print(bool(Perm.R & Perm.R))
output > True
print(bool(Perm.R & Perm.W))
output > False
print(bool(Perm.R & (Perm.W | Perm.R)))
output > True
print(Perm.R in (Perm.W | Perm.R))
output > True

```

### `enum.Flag`

```python
class Colors(enum.Flag):
    RED = enum.auto()
    BLUE = enum.auto()
    GREEN = enum.auto()
    WHITE = RED | BLUE | GREEN
    
colors_names = [str(color.name) for color in Colors]
print(", ".join(colors_names))
output > 'RED, BLUE, GREEN, WHITE'

color_values = [str(color.value) for color in Colors]
print(", ".join(color_values))
output > '1, 2, 4, 7'
```

## `dataclasses.dataclass` 

* [Dataclasses: The code generator to end all code generators - PyCon 2018](https://www.youtube.com/watch?v=T-TwcmT6Rcw) by Raymond Hettinger [slides](https://www.dropbox.com/s/te4q0xf46zkuu21/hettinger_dataclasses_pycon_2018.zip), [pdf](https://www.dropbox.com/s/m8pwkkz43qz5pgt/HettingerPycon2018.pdf)

```python
import dataclasses

@dataclasses.dataclass(order=True)
class Person:
    """
A *very* basic representation of a person, created as a dataclass
"""
    given_name: str
    family_name: str
    birth_date: str = None
    email_address: str = None
        print(Person)
        output > <class '__main__.Person'>
        print(Person)
        output > <class '__main__.Person'>

# - Classes defined as dataclasses are still classes...
print(Person)
output > <class '__main__.Person'>
# - They still appear as normal "type" types
print(type(Person))
output > <class 'type'>
# - They have additional members
# ic('__annotations__ ........ %s' % Person.__annotations__)
# ic('__dataclass_fields__ ... %s' % Person.__dataclass_fields__)
# ic('__dataclass_params__ ... %s' % Person.__dataclass_params__)

# - And have some that have default implementations provided:
p1 = Person('Persona', 'Aname')
print(p1)
output > Person(given_name='Persona', family_name='Aname', birth_date=None, email_address=None)

p2 = Person('Personica', 'Bname', None, 'bname@gmail.com')
print(p2)
output > Person(given_name='Personica', family_name='Bname', birth_date=None, email_address='bname@gmail.com')

# - order=True has to be passed as an argument to the 
#   dataclass decorator for these to work:
print('p1 == p2 ... %s' % (p1 == p2))
output > 'p1 == p2 ... False'
print('p1 != p2 ... %s' % (p1 != p2))
output > 'p1 != p2 ... True'
print('p1 < p2 .... %s' % (p1 < p2))
output > 'p1 < p2 .... True'
print('p1 <= p2 ... %s' % (p1 <= p2))
output > 'p1 <= p2 ... True'
print('p1 > p2 .... %s' % (p1 > p2))
output > 'p1 > p2 .... False'
print('p1 >= p2 ... %s' % (p1 >= p2))
output > 'p1 >= p2 ... False'

```

## Comparison

 Dataclass                        | NamedTuple
----------------------------------|-------------
`replace()` function	          | `_replace()` method
`asdict()` function	              | `_asdict()` method
converts to regular `dict`        | converted to OrderedDict
`astuple()` function              | `tuple()` function
Mutable                           | Frozen
Unhashable                        | Hashable
Non-iterable                      | Iterable and unpackable
No comparison methods             | Sortable
Underlying store: instance `dict` | Underlying store: tuple
168 bytes                         | 72 bytes
33 ns access                      | 61 ns access

