

# Decorators

* [PyCon2019: Practical decorators](https://www.youtube.com/watch?v=MjHpMCIvwsY) by Reuven M. Lerner - [slides](https://speakerdeck.com/pycon2019/reuven-m-lerner-practical-decorators)  

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
print(Power.root(4))
output > Power(4^0.5) is 2.0
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
