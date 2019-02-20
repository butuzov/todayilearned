

# Programming in Python (notes)

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

## Code Quality

### Code Standards - PEP8

* https://pep8.org/
* https://www.python.org/dev/peps/pep-0008/
* http://pep8online.com/

**Notes**

* Use `space`, no `tab`. Always 4.
* Lines should be no more then 79 characters in length or less, if you need to
  - make string shorter simplifieing meaningfull variables.
  - add `\` to split code to the new line.
  - continuations of long expressions, should be added by one indentation (4 spaces) 
    of their normal indentation level.
* Functions and Classes should be represented by two blank lines.
* Class methods should be separated by one blank lines.
* don't put spaces around list indexes, function calls, or keyword assignments.
* Put one, and only one space before and after variable assignment.
* etc... etc...

### Code Analizers

* [ ] `pep8` PEP8
* [ ] `pyflackes` code analyzer
* [ ] `pylint` shows us a styling and logical errors
* [ ] `codecov` code coverage
* [ ] `vulture` dead code
* [ ] `flake8` all in one
* [ ] `black` beutifier and code standards.

### Flake8

```bash
git init
git add example.py
git commit -m "first commit"
> ....

python -m flake8 --install-hook git
git config --bool flake8.strict true

echo "# new comment" >> example.py
git commit -m "second commit"
> error messages...
```

## Data Structures

### Sequences (Unpacking)

```python
data = ['ACME', 50, 91.1, (2012, 12, 21)]
print(data)
output > ['ACME', 50, 91.1, (2012, 12, 21)]

# Simple unpucking
name, shares, price, date = data
print(name)
output > 'ACME'

# Tuple unpacking
name, shares, price, (year, mon, day) = data
print(year)
output > 2012
print(year)
output > 2012

# Nested Unpacking
name, *_, (year, *_) = data
print(year)
output > 2012

# we can unpack any sequence of same length
s = 'hello'
a, b, c, d, f = s
print(b)
output > 'e'
print(b)
output > OrderedDict([('a', 1), ('b', 3), ('c', 2)])

# skiping some of the unpackined values (works in any position start, middle or end)
a, *b, f = s

# we can unpack into function
print(*b, sep=':')

# if this is a dictionary, we can unpack named arguments, 
# alongside with positional.
print(*b, **{'sep':':'})
```

### Dictionaries, Maps, and Hashtables

[Modern Python Dictionaries - A confluence of a dozen great ideas (PyCon 2017)](https://www.youtube.com/watch?v=npw4s1QTmPg) by Raymond Hettinger

* [x] `collections.OrderedDict`
* [ ] `collections.defaultdict`
* [ ] `collections.ChainMap` - multiple dictionaries combined (for search proposes, for example)
* [ ] types .MappingProxyType - read only dict (frozendict).

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

```python
# default values of dictionary
```

### Array Data Structures

* [x] slicing `list`s
* [ ] `list`s
* [ ] `tuple`s
* [x] `array.array` - typed lists
* [ ] `str` - immutable strings list (unicode)
* [ ] `bytes` - immutable bytes container
* [ ] `bytearray` - mutable bytes container

```python
# Slice has simular syntax as range
# slice(stop)
# slice(start, stop)
# slice(start, stop, step)

first_ten_numbers = list(range(10))
print(first_ten_numbers)
output > [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

# named slice - slice(stop)
first_two_numbers = slice(2)
print(first_ten_numbers[first_two_numbers])
output > [0, 1]

# named slice - slice(start, stop)
rest_of_numbers = slice(2,len(first_ten_numbers))
print(first_ten_numbers[rest_of_numbers])
output > [2, 3, 4, 5, 6, 7, 8, 9]

# simple slicing using indexes
print(first_ten_numbers[0:2])
output > [0, 1]

# reverse printing
print(first_ten_numbers[::-1])
output > [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]

# using slices as indexes
q = [1,2,3,4,5,6,7,8,9,0]
print(q)
output > [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]

# all without last element
print(q[:-1])
output > [1, 2, 3, 4, 5, 6, 7, 8, 9]    

# last element, start from end.
print(q[-1])
output > 0     

# start from 3rd, till end with step of 2
print(q[2::4])
output > [3, 7]   

# start from 3rd... and reverse it!
print(q[2::-1])
output > [3, 2, 1]  

# start from 9th .. till 4rd.. and reverse it!
print(q[8:3:-1])
output > [9, 8, 7, 6, 5] 
```

```python
# C-types arrays  array.array 
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

### Records, Structs, and Data Transfer Objects

  * [ ] `dict`s
  * [ ] `tuple`s & `collections.namedtuple`s
  * [ ] `typing.NamedTuple` improved named `tuple`'s
  * [x] [`struct.Struct`] 
  * [ ] `types.SimpleNamespace`

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

### Sets
  * [ ] `sets` `{1,2,3}`
  * [ ] `frozenset` - immutable `set`
  * [x] `collections.Counter` - weird set (allow multiple values), but has counter on it.

```python
# Sets
s = set([1,2,1,2,3,1])
print(s)
output > {1, 2, 3}
```

```python
# collections.Counter example

from collections import Counter

c1 = Counter();
c1['apples'] += 1
c1['apples'] += 2
c1['peaches'] += 10

# Most Common Elements (sorted by `count`)
print(c1.most_common())
output > [('peaches', 10), ('apples', 3)]

# Counter Example
c2 = Counter()
lipsum = "Un dos tres quatro sinco seis"
words = [w.strip(".,") for w in lipsum.lower().split(' ')]   
for word in words:
    c2[word]+=1
print(c2.most_common(2))
output > [('un', 1), ('dos', 1)]

# List
print(list(c2.elements()))
output > ['un', 'dos', 'tres', 'quatro', 'sinco', 'seis']

# Unique Elements
print(list(c2))
output > ['un', 'dos', 'tres', 'quatro', 'sinco', 'seis']

# Additoional methods
c3 = Counter(["dos"])
c3.update(["dos", "tres", "sinco", "seis", "ocho"])
c3.subtract(["ocho"])
print(c3)
output > Counter({'dos': 2, 'tres': 1, 'sinco': 1, 'seis': 1, 'ocho': 0}) 

```

### Queues (FIFOs)

* [ ] `list` as example of stack
* [x] `collections.deque` double linked list.
* [ ] `queue.Queue` paralel cumputation blocking.
* [ ] `multiprocessing.Queue`

### Stacks (LIFOs)

* `list` as example of stack
* `collections.deque` double linked list.
* `queue.Queue` paralel cumputation blocking.
* `multiprocessing.Queue`
  

```python
# Dequeue Small example 

from collections import deque as Deque

d = Deque(["one", "two", "three"])
print(d)
output > deque(['one', 'two', 'three'])
print(d)
output > deque(['one', 'two', 'three', 'four'])
print(d)
output > deque(['zero', 'one', 'two', 'three', 'four'])
print(d)
output > deque(['two', 'three', 'one'])

d.append("four")
print(d)
output > deque(['one', 'two', 'three', 'four'])
print(d)
output > deque(['zero', 'one', 'two', 'three', 'four'])
print(d)
output > deque(['two', 'three', 'one'])

d.appendleft("zero")
print(d)
output > deque(['zero', 'one', 'two', 'three', 'four'])
print(d)
output > deque(['two', 'three', 'one'])

d_poped = d.pop()
print(d_poped)
output > 'four'

d_popedleft = d.popleft()
print(d_popedleft)
output > 'zero'

# rotatiob
d.rotate(2)
print(d)
output > deque(['two', 'three', 'one'])
```

### PriorityQueues

  * `list`, but it depends
  * `heapq`
  * `queue.PriorityQueue`

## Environments

### Python Packages 

* https://www.pypa.io/en/latest/

```bash
# Python 2.7 (mac system)

# https://stackoverflow.com/questions/17271319
sudo easy_install pip

# Version and location
pip -V
```

* Activate virtual environment before working with pip
* If not use `python -m pip` form instead of `pip` or `pip3`. 
* default cache location on mac `~/Library/Caches/pip`

```bash 

# Version and location
pip -V

# reset (alias)
pip list | awk 'NR>2 {print $1}' | grep -Ev "pip|setuptools|wheel" | xargs -I {} sh -c "python -m pip uninstall {} -y"

# Using Custom cache directory
pip install --download-cache /cache/directory matplotlib

# Create Requirments
pip freeze > requirments.txt

# Install packages from requirments.txt
pip install -r requirements.txt

# Specific version
pip install flask==0.9
pip install 'Django>2.0'

# Help
pip <command> --help

# info about installed package
pip show tensorflow

# list outdated packages
pip list -o
```

### Virtual Environment

```bash
# Python 2.7 
$ sudo python -m pip install virtualenv

# Create Virtual Env
$ python -m virtualenv Py2.7

# Activate Environment
$ . Py2.7/bin/activate 

(Py2.7) $ python -V
> Python 2.7.10

# Deactivate Environment
$ deactivate 

# Python 3.6 (or any else)
$ python3 -m venv Py3.7

# Activate Environment
$ . Py3.7/bin/activate
(Py3.7) $ python -V
> Python 3.7.2

# Check ${VIRTUAL_ENV}
$ echo "${VIRTUAL_ENV}"
> /Users/butuzov/Desktop/Py3.7

# Deactivate Environment
$ deactivate 
```

### Local Development

```bash
(venv) setup.py develop
```

##### `setup.py`

```python
from setuptools import setup, find_packages

setup(
    name="doc.md",
    version="0.0.1",
    description="Single Page Help Generator",
    license="MIT",
    author="Oleg Butuzov",
    author_email='butuzov@made.ua',
    packages=find_packages(),
    install_requires=[],
    scripts=['bin/docmd'],
    zip_safe=False,
    classifiers=[
        "Programming Language :: Python",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: 3.7",
    ]
)
```

### Jupyter

https://www.dataquest.io/blog/jupyter-notebook-tips-tricks-shortcuts/

* Extensions *

```bash
(venv) > pip install install jupyter jupyter_contrib_nbextensions
(venv) > pip install git+git://github.com/cpcloud/ipython-autotime
(venv) > jupyter contrib nbextension install --sys-prefix
(venv) > jupyter nbextension enable codefolding/main
(venv) > jupyter nbextension disable codefolding/main
```

* Nice extensions *

 1. jupyter nbextension enable execute_time/ExecuteTime
 1. jupyter nbextension enable collapsible_headings/main 

* Loading Extensions into cell *
```
# Loading extension
# pip install git+git://github.com/cpcloud/ipython-autotime
%load_ext autotime
```

```python
# import sys

# sys.getrefcount(1)
# sys.getsizeof(1)
```

## Development and Testing

### PyTest

https://docs.pytest.org/en/latest/index.html

```ini
[pytest]
norecursedirs = .venv .vscode
addopts = --verbose
testpaths = doctor
console_output_style = progress
python_files = *_test.py
python_classes = Test*
python_functions = test_*
```

#### Hello World - Simple Test

```python
# filename: cases.py
cases = []
cases.appen( ( 1, 2, 3 ) )
cases.appen( ( 5, 2, 7 ) )
```
----

```python
# filename: add_test.py
from package.some_cases import cases

@pytest.mark.parametrize('a,b', cases)
test_add(a,b,c):
  assert c == add(a,b)
```
----
```python
# filename: add.py
def add(a,b):
  return a+b
```

## Resources

### Books

* [Python Tricks: The Book](https://dbader.org/products/python-tricks-book/)

### Podcasts

* [Talk Python to me](https://talkpython.fm/) by [Michael Kennedy](https://twitter.com/mkennedy)
* [Podcast.__init__('Python')](https://www.podcastinit.com/) by [Tobias Macey](https://twitter.com/TobiasMacey)
* [TestAndCode](https://www.podcastinit.com/) by [Brian Okken](https://twitter.com/brianokken)

### WebSites

  * [PythonistaCafe](www.pythonistacafe.com) - comunity website.
  * [/r/Python](https://www.reddit.com/r/Python/) @ Reddit
  * [Dan Bader](https://dbader.org/) Dan Bader's website.

