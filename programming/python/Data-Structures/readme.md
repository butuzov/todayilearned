

# Python Data Structures & Types

## Data Types

```python
### int
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

### `bytearray`

`mutable` bytes seqence

```python
# Same way to initialize it was done with byte

# Mutating array
b = bytearray("Dots over і.", 'utf8')
b.extend(bytes('I said put dots over i.', 'utf8'))

b[0:5] = b'Comas '

# bytes supports strings like opperations
words = b"Dont comes easy"

# text transform

# text split to list

# and join it back

```

### `array.array`

`C`- types arrays

```python
# https://docs.python.org/3/library/array.html
from array import array 

# Unicode (Py_UNICODE)

# signed long

# doubles

# floats 

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

```

```python
### `queue.PriorityQueue`
```

### `queue.deque`

```python
import collections 

d = collections.deque(["one", "two", "three"])
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try'])
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])

d.append("four")
print(d)
output > deque(['si', 'san', 'ni', 'iti', 'one', 'two', 'four', 'dva', 'try'])
print(d)
output > deque(['dva', 'try', 'si', 'san', 'ni', 'iti', 'one', 'two', 'four'])

d.extend(["dva", "try", "chotyry"])

d.extendleft(["iti", "ni", "san", "si"])

d.appendleft("zero")

d_poped = d.pop()

d_popedleft = d.popleft()

# rotatiob

d.rotate(2)

```

## Sets

### `set`

```python
s = set([1,2,1,2,3,1])

s.add(19)
s.add(19)

# intersection

```

```python
### `frozenset`
```

### `collections.Counter`

```python
# collections.Counter example

from collections import Counter

c0 = Counter({'uno':1, 'dos':2});

c1 = Counter(uno=1)

c2 = Counter();
c2['apples'] += 1
c2['apples'] += 2
c2['peaches'] += 10

# Most Common Elements (sorted by `count`)

# Counter Example
c3 = Counter()
lipsum = "Un dos tres quatro cinco seis siete"
words = [w.strip(".,") for w in lipsum.lower().split(' ')]   
for word in words:
    c3[word]+=1

# List

# Unique Elements

# Additoional methods
c4 = Counter(["dos"])
c4.update(["dos", "tres", "cinco", "seis", "ocho"])
c4.subtract(["ocho"])

# background implementation - itertools chain

# adding counters

# subtraction

# maximum value (union)

# minimum value (intersection)

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

# And you can instantiate class with wraper implementing custom methods
class CarMethods(Car):
    def hex(self):
        return ('#000', "#f00")[self.color != "red"]

# replace value
redSubaru._replace(color="#000")

# show tuple as dictionary

# and finaly - factory

```

## Enum

### `enum.IntEnum`

```python
from enum import IntEnum

class RequestType(IntEnum):
    POST = 1
    GET = 2

request_type = RequestType.POST

```

-------

## rework

-----

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

# - Classes defined as dataclasses are still classes...

# - They still appear as normal "type" types

# - They have additional members
# ic('__annotations__ ........ %s' % Person.__annotations__)
# ic('__dataclass_fields__ ... %s' % Person.__dataclass_fields__)
# ic('__dataclass_params__ ... %s' % Person.__dataclass_params__)

# - And have some that have default implementations provided:
p1 = Person('Persona', 'Aname')

p2 = Person('Personica', 'Bname', None, 'bname@gmail.com')

# - order=True has to be passed as an argument to the 
#   dataclass decorator for these to work:

```

![image.png](attachment:image.png)

