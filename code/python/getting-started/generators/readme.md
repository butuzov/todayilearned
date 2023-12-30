# generators

## infinite sequence

```python
list(range(5))

result >>> [0, 1, 2, 3, 4]
```

```python
def infinite_sequence():
    num = 0
    while True:
        yield num
        num += 1

gen = infinite_sequence()
[next(gen), next(gen), next(gen), next(gen), next(gen)]

result >>> [0, 1, 2, 3, 4]
```

```python
# Simple sequence

def sequence(n):
    num = 0
    while num < n:
        yield num
        num += 1
        
list(sequence(5))

result >>> [0, 1, 2, 3, 4]
```

```python
# Simple sequence

def sequence_(n):
    """ will stop with StopIter """
    num = 0
    while True:
        yield num
        num += 1
        if num > n:
            raise StopIteration

for it in sequence_(5):
    print()

> 0
> 1
> 2
> 3
> 4
> 5
```

## Profiling Generator Performance

```python
import sys

# case list comprehension
sys.getsizeof([i ** 2 for i in range(10000)]), sys.getsizeof(i ** 2 for i in range(10000))

result >>> (85176, 208)
```

```python
import cProfile
```

```python
cProfile.run('sum([i * 2 for i in range(10000)])')

> 5 function calls in 0.001 seconds
> 
> Ordered by: standard name
> 
> ncalls  tottime  percall  cumtime  percall filename:lineno(function)
> 1    0.001    0.001    0.001    0.001 <string>:1(<listcomp>)
> 1    0.000    0.000    0.001    0.001 <string>:1(<module>)
> 1    0.000    0.000    0.001    0.001 {built-in method builtins.exec}
> 1    0.000    0.000    0.000    0.000 {built-in method builtins.sum}
> 1    0.000    0.000    0.000    0.000 {method 'disable' of '_lsprof.Profiler' objects}
> 
>
```

```python
cProfile.run('sum((i * 2 for i in range(10000)))')

> 10005 function calls in 0.003 seconds
> 
> Ordered by: standard name
> 
> ncalls  tottime  percall  cumtime  percall filename:lineno(function)
> 10001    0.001    0.000    0.001    0.000 <string>:1(<genexpr>)
> 1    0.000    0.000    0.003    0.003 <string>:1(<module>)
> 1    0.000    0.000    0.003    0.003 {built-in method builtins.exec}
> 1    0.001    0.001    0.003    0.003 {built-in method builtins.sum}
> 1    0.000    0.000    0.000    0.000 {method 'disable' of '_lsprof.Profiler' objects}
> 
>
```

## Advanced Generator Methods

### `.send()`

### `.close()`

### `.throw()`