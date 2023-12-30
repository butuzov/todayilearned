# Dunders In Python

## PEP 560: Core Support for typing module and Generic Types

The PEP introduces two special methods `__class_getitem__()` and `__mro_entries__`, these methods are now used by most classes and special constructs in typing. As a result, the speed of various operations with types increased up to 7 times, the generic types can be used without metaclass conflicts, and several long standing bugs in typing module are fixed.

## Callables with `__call__`

There is a way to make instance callable.

```python
class prophecy:
    
    def __call__(self, type_of_prefecy):
        return self.grim() if type_of_prefecy == 'grim' else "no prophecy for you"
    
    def grim(self):
        return "we all gona die!"
    
p = prophecy()
print(p('grim'))
print(p("other_type_prophecy"))

> we all gona die!
> no prophecy for you
```

## Object Oriented Programming `__init__` & `__del__`

Object constors and descructors.

```python
class constr_destr:
    def __init__(self): 
        print("i am alive") 
    def __del__(self): 
        print("i am dead")
    
c = constr_destr()
del c

> i am alive
> i am dead
```

### Ordering and Comparing 

Ordering and Comparing things with `__lt__`, `__le__`, `__gt__`, `__ge__` and `__eq__`. `functools.total_ordering` decorator. Implementations any of this methods allows to compare objects. Usage of `functools.total_ordering` allows to skipp full implmentation and limit only to two methods `__eq__` and anyone of the list `__lt__`, `__le__`, `__gt__`, `__ge__`.



```python
from functools import total_ordering

@total_ordering
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
    
Weight('5 kg') == Weight('5000 gr'), Weight('5 kg') <  Weight('4999 gr'), Weight('5 kg') >  Weight('4999 gr')

result >>> (True, False, True)
```

### Representation methods `__repr__` & `__str__`

```python
class ExampleRepr:
    "represents x number "
    def __init__(self, x = int):
        self.x = x
        
    def __repr__(self) -> str:
        return "{}({})".format(self.__class__.__name__, self.x)

class ExampleStr(ExampleRepr):
    def __str__(self):
        return "{}".format(self.x)
    
class ExampleHTML(ExampleStr):
    " method for classes to be used in jupyter to prepresent html "
    def _repr_html_(self):
        return "Value of x is <em>{}</em>".format(self.x)
        

# ic(ExampleRepr(2))
print(ExampleRepr(2))

# ic(ExampleStr(2))
print(ExampleStr(2))

ExampleHTML(2)

> ExampleRepr(2)
> 2
```

Value of x is <em>2</em>

## Context Managers with `__enter__` and `__exit__`

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
with ManagedFile('hello.txt') as f:
    f.write('¡holla!')
```

## Math Operations (regular ones) - overloading

```python
class Math:
    
    def __init__(self, a):
        self._a = a
    
    # --- regular math operators overloading     
    def __add__(self, b):
        """ Addition with - """
        return self.__class__(self.value + b.value)
        
    def __sub__(self, b):
        """ Subtraction with - """
        return self.__class__(self.value - b.value)
    
    def __mul__(self, b):
        """ Multiplication with * """
        return self.__class__(self.value * b.value)
    
    def __mod__(self, b):
        """ Modulus operation % """
        return self.__class__(self.value % b.value)
    def __rpow__(self, b):
        print("not affected")
        
    def __pow__(self, b):
        """ powering value with ** """
        return self.__class__(self.value ** b.value)
    
    def __floordiv__(self, b):
        """ Division with // """
        return self.__class__(self.value // b.value)
    
    def __truediv__(self, b):
        """ Division with / """
        return self.__class__(self.value / b.value)

    def __divmod__(self, b):
        return self.__floordiv__(b), self.__mod__(b)

    def __matmul__(self, b):
        """ matrix multiplication @ """ 
        return self.__class__(self.value * b.value)
    
    def __repr__(self):
        return f"{self._a}"
    
    @property
    def value(self):
        return self._a
    
    
Math(1)+Math(2),Math(1)-Math(2), Math(1)*Math(2), Math(1)/Math(2), Math(1)//Math(2), Math(1)**Math(2), Math(1)%Math(2), Math(1) @ Math(2), divmod(Math(1), Math(2)),

result >>> (3, -1, 2, 0.5, 0, 1, 1, 2, (0, 1))
```

## Math Operation (inplace and right-operated)

```python
#todo
```

## TypeCasting

```python
class TypeInt:
    def __init__(self, n):
        self.n = n
        
    def __int__(self):
        return self.n
    
    def __str__(self):
        return "one" if self.n == 1 else "I dont know, its complicated number".format(self.n)
    
    def __float__(self):
        return float(self.n)
    
    def __bool__(self):
        return True
```

```python
int(TypeInt(1)), str(TypeInt(1)), float(TypeInt(1)), bool(TypeInt(0))

result >>> (1, 'one', 1.0, True)
```