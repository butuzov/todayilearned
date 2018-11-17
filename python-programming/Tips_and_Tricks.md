# Python: Tips & Tricks

>Cumulative Notes on Python Programming

# Table of Contents
[TOC]

---
## Idiomatic Code

### [PEP 8 - Style Guide for Python Code](https://www.python.org/dev/peps/pep-0008/)
* Use `space`s, no `tab`s. Always 4.
* Lines should be no more then 79 characters in length or less, if you need to
  * make string shorter simplifieing meaningfull variables.
  * add `\` to split code to the new line.
  * continuations of long expressions, should be added by one indentation (4 spaces) of their normal indentation level.
* Functions and Classes should be represented by two blank lines.
* Class methods should be separated by one blank lines.
* don't put spaces around list indexes, function calls, or keyword assignments.
* Put one, and only one space before and after variable assignment.
* etc... etc...
More info at [`pep8` site](https://pep8.org/) and [official page](https://www.python.org/dev/peps/pep-0008/)

##### Install PEP8 and pylint in addition to the VSC Python Extension
```bash
python3 -m pip install --upgrade pep8 pylint
```

---
## Sequences

### Unpacking

```python
data = ['ACME', 50, 91.1, (2012, 12, 21)]

# simple unpacking
name, shares, price, date = data
name
> 'ACME'

# tuple unpacking
name, shares, price, (year, mon, day) = data
year
> 2012

# nested unpacking permited
name, *_, (year, *_) = data
year
> 2012

# we can unpack any sequence of same length
s = 'hello'
a, b, s, d, f = s
b
> 'e'

# skiping some of the unpackined values (works in any position start, middle or end)
a, *b, f = s
a, b, f
> ('h', ['e', 'l', 'l'], 'o')

# we can unpack into function
print(*b, sep=':')
> e:l:l

# if this is a dictionary, we can unpack named arguments, alongside with positional.
print(*b, **{'sep':':'})
> e:l:l
```

---
## Dictionaries

### Data Structures - Dictionaries

Order in dictionaries depends on it's [implementation](https://www.youtube.com/watch?v=npw4s1QTmPg)

```python
a = {'a':1, 'b':3, 'c':2}
a
# in python 3.7, but it's depends on implementation
> {'a': 1, 'b': 3, 'c': 2}

# ordered dict helps to keep it really ordered
from collections import OrderedDict
b = OrderedDict(a)
b
> OrderedDict([('a', 1), ('b', 2), ('c', 2)])

# we also can use sorted to sort doctionary by value
sorted(a.items(), key=lambda x: x[1])
> [('a', 1), ('c', 2), ('b', 3)]
# or by key
sorted(a.items(), key=lambda x: x[0])
> [('a', 1), ('b', 3), ('c', 2)]
```

---
## Data Structures - Slices

```python
# slice has simular syntax as range
# slice(stop)
# slice(start, stop)
# slice(start, stop, step)

first_ten_numbers = list(range(10))
print(first_ten_numbers)
> [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

# named slice - slice(stop)
first_two_numbers = slice(2)
print(first_ten_numbers[first_two_numbers])
> [0, 1]

# named slice - slice(start, stop)
rest_of_numbers = slice(2,len(first_ten_numbers))
print(first_ten_numbers[rest_of_numbers])
> [2, 3, 4, 5, 6, 7, 8, 9]

# simple slicing using indexes
print(first_ten_numbers[0:2])
> [0, 1]

# reverse printing
print(first_ten_numbers[::-1])
> [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]
```


## Packages

### Local Package Development

You can develop local packages installing package localy (or even in `venv`)

```bash
python3 -m venv .venv
. ./venv/bin/activate
python3 setup.py develop
```

setup.py sample
```python3
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

### TDD Development

It's easy using [pytest](https://docs.pytest.org/en/latest/index.html).

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

##### some_cases.py
```python3

cases = []
cases.appen( ( 1, 2, 3 ) )
cases.appen( ( 5, 2, 7 ) )
```

##### some_test.py
```python3

from package.some_cases import cases

@pytest.mark.parametrize('a,b', cases)
test_add(a,b,c):
  assert c == add(a,b)
```

##### some.py
```python3
def add(a,b):
  return a+b
```
###### @todo: Buy and Read [Python Testing with pytest](https://pragprog.com/book/bopytest/python-testing-with-pytest) by Brian Okken
###### @todo: checkout [Hypothesis](https://hypothesis.readthedocs.io/en/latest/)
###### @todo: checkout [flake8](https://pypi.org/project/flake8/)
