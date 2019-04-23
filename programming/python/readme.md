

# Python Notes

## Language

### Coveing Your A** With Assertions

```python
# Don't do
# Becouse of tuple its will always will be true
assert( 1 == 2, "It's failing!") 
```

### Complacent Comma Placement

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

### Context Managers and the with Statement

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

# So we can do something like implementing custom write or read
# or any other functionality.
with ManagedFile("list.json") as f:
    f.read()

""" or using @contextmanager """  
    
from contextlib import contextmanager

@contextmanager
def managed_file(name):
  try:
    f = open(name, 'w')
    yield f
  finally:
    f.close()

with managed_file('hello.txt') as f:
    f.write('¡holla!')
    f.write('¡adios!')
```

### (Unpacking) Seaquences 

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

# skiping some of the unpackined values (works in any position start, middle or end)
a, *b, f = s

# we can unpack into function
print(*b, sep=':')

# if this is a dictionary, we can unpack named arguments, 
# alongside with positional.
print(*b, **{'sep':':'})
```

### `slice`ing lists

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

## Tooling and Standards

### PEP8 

**Summary**

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

**Readme more**:
* https://pep8.org/
* https://www.python.org/dev/peps/pep-0008/
* http://pep8online.com/

### @pyflackes

Code analyzer 

### @pylint

Shows us a styling and logical errors

### @codecov

Code coverage

### @vulture

Dead code

### `flake8`

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

### `venv` or `virtualenv`

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

