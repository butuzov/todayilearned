

```python
from icecream import ic
import sys

def jupyter(*args):
    print(*args, file=sys.stdout)

ic.configureOutput(prefix='> ', outputFunction=jupyter)
```

# Programming in Python (notes)

[TOC]

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

### Pyflakes

`pyflakes` code analyzer

### `Pylint`

`pylint` shows us a styling and logical errors

### Vulture

`vulture` is searching for dead code.

### `codecov`


```python

```

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


```python

```

### black


```python

```


```python

```


```python

```


```python

```


```python

```


```python

```


```python

```

## Data Structures

### Counter


```python
from collections import Counter

c1 = Counter();
c1['apples'] += 1
c1['apples'] += 2
c1['peaches'] += 10

# Most Common Elements (sorted by `count`)
ic(c1.most_common())

# Counter Example
c2 = Counter()
lipsum = "Un dos tres quatro sinco seis"
words = [w.strip(".,") for w in lipsum.lower().split(' ')]
for word in words:
    c2[word]+=1

ic(c2.most_common(2))

# List
ic(list(c2.elements()))

# Unique Elements
ic(list(c2))

# Additoional methods
c3 = Counter(["dos"])
c3.update(["dos", "tres", "sinco", "seis", "ocho"])
c3.subtract(["ocho"])
ic(c3)

```

    > c1.most_common(): [('peaches', 10), ('apples', 3)]
    > c2.most_common(2): [('un', 1), ('dos', 1)]
    > list(c2.elements()): ['un', 'dos', 'tres', 'quatro', 'sinco', 'seis']
    > list(c2): ['un', 'dos', 'tres', 'quatro', 'sinco', 'seis']
    > c3: Counter({'dos': 2, 'tres': 1, 'sinco': 1, 'seis': 1, 'ocho': 0})





    Counter({'dos': 2, 'tres': 1, 'sinco': 1, 'seis': 1, 'ocho': 0})



    time: 203 ms


### OrderedDict


```python
# in python 3.7 order will be preserved,
# but it's depends on dict implementation
a = {'a':1, 'b':3, 'c':2}
ic(a)

# ordered dict helps to keep it really ordered
from collections import OrderedDict
b = OrderedDict(a)
ic(b)

# we also can use sorted to sort doctionary by value
ic(sorted(a.items(), key=lambda x: x[1]))

# or by key
ic(sorted(a.items(), key=lambda x: x[0]))
```

    > a: {'a': 1, 'b': 3, 'c': 2}
    > b: OrderedDict([('a', 1), ('b', 3), ('c', 2)])
    > sorted(a.items(), key=lambda x: x[1]): [('a', 1), ('c', 2), ('b', 3)]
    > sorted(a.items(), key=lambda x: x[0]): [('a', 1), ('b', 3), ('c', 2)]





    [('a', 1), ('b', 3), ('c', 2)]



### Slices


```python
# slice has simular syntax as range
# slice(stop)
# slice(start, stop)
# slice(start, stop, step)

first_ten_numbers = list(range(10))
ic(first_ten_numbers)

# named slice - slice(stop)
first_two_numbers = slice(2)
ic(first_ten_numbers[first_two_numbers])

# named slice - slice(start, stop)
rest_of_numbers = slice(2,len(first_ten_numbers))
ic(first_ten_numbers[rest_of_numbers])

# simple slicing using indexes
ic(first_ten_numbers[0:2])

# reverse printing
ic(first_ten_numbers[::-1])
```

### Sequences (Unpacking)


```python
data = ['ACME', 50, 91.1, (2012, 12, 21)]
ic(data)

# Simple unpucking
name, shares, price, date = data
ic(name)

# Tuple unpacking
name, shares, price, (year, mon, day) = data
ic(year)

# Nested Unpacking
name, *_, (year, *_) = data
ic(year)


# we can unpack any sequence of same length
s = 'hello'
a, b, c, d, f = s
ic(b)

# skiping some of the unpackined values (works in any position start, middle or end)
a, *b, f = s
ic(a, b, f)

# we can unpack into function
print(*b, sep=':')

# if this is a dictionary, we can unpack named arguments,
# alongside with positional.
print(*b, **{'sep':':'})
```

    > data: ['ACME', 50, 91.1, (2012, 12, 21)]
    > name: 'ACME'
    > year: 2012
    > year: 2012
    > b: 'e'
    > a: 'h', b: ['e', 'l', 'l'], f: 'o'
    e:l:l
    e:l:l


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


```python

```


```python

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


## Object Oriented Programming


```python

```


```python

```


```python

```


```python

```


```python

```


```python

```


```python

```


```python

```

## SECTION


```python

```
