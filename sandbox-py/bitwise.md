### Bitwise algorithmes and hacks

* http://bitwisecmd.com/
* https://graphics.stanford.edu/~seander/bithacks.html
* https://www.hyfr.life/bitwise-operators-tricks/
* https://catonmat.net/low-level-bit-hacks
* https://github.com/knoxknox/bit-hacks

## Python

* https://wiki.python.org/moin/BitwiseOperators

```python
def b(n: int) -> str:
    return "Binary 0x{0:08b} / Decimal {0:d}".format(n)
```

```python
# display int in binary format
b(10)

result >>> 'Binary 0x00001010 / Decimal 10'
```

### `x << y`

```python
for i,v in enumerate([1, 3, 5]):
    print(f"{i+1})", "in=>", b(v), "out=>" , b(v << 1))

stdout >>> 1) in=> Binary 0x00000001 / Decimal 1 out=> Binary 0x00000010 / Decimal 2
stdout >>> 2) in=> Binary 0x00000011 / Decimal 3 out=> Binary 0x00000110 / Decimal 6
stdout >>> 3) in=> Binary 0x00000101 / Decimal 5 out=> Binary 0x00001010 / Decimal 10
```

### `x >> y`

```python
for i,v in enumerate([64, 32, 16, 8, 4, 2, 1]):
    print(f"{i+1})", "in=>", b(v), "out=>" , b(v >> 1))

stdout >>> 1) in=> Binary 0x01000000 / Decimal 64 out=> Binary 0x00100000 / Decimal 32
stdout >>> 2) in=> Binary 0x00100000 / Decimal 32 out=> Binary 0x00010000 / Decimal 16
stdout >>> 3) in=> Binary 0x00010000 / Decimal 16 out=> Binary 0x00001000 / Decimal 8
stdout >>> 4) in=> Binary 0x00001000 / Decimal 8 out=> Binary 0x00000100 / Decimal 4
stdout >>> 5) in=> Binary 0x00000100 / Decimal 4 out=> Binary 0x00000010 / Decimal 2
stdout >>> 6) in=> Binary 0x00000010 / Decimal 2 out=> Binary 0x00000001 / Decimal 1
stdout >>> 7) in=> Binary 0x00000001 / Decimal 1 out=> Binary 0x00000000 / Decimal 0
```

### `x & y`

### `x | y`

### `~x`

Returns the complement of `x` - the number you get by switching each 
* `1` for a `0`.
* `0` for a `1`.

This is the same as `-x - 1`.

```python
for i in [('1', 1), ('~1', ~1)]: print(f"{i[0]:>4s}:", b(int(i[1])))

stdout >>>    1: Binary 0x00000001 / Decimal 1
stdout >>>   ~1: Binary 0x-0000010 / Decimal -2
```

### `x ^ y`

### Examples

`add_bitwise_operator.py`

Adding two positive integers without `+` operator

```python
def add_bitwise_operator(x, y):
    while y:
        carry = x & y
        print("{:>12s}:".format('carry'), b(carry))
        
        x = x ^ y
        print("{:>12s}:".format('x'), b(x))
        
        y = carry << 1
        print("{:>12s}:".format('y'), b(y))

        print("="*45)
    return x

add_bitwise_operator(12, 5)

stdout >>>        carry: Binary 0x00000100 / Decimal 4
stdout >>>            x: Binary 0x00001001 / Decimal 9
stdout >>>            y: Binary 0x00001000 / Decimal 8
stdout >>> =============================================
stdout >>>        carry: Binary 0x00001000 / Decimal 8
stdout >>>            x: Binary 0x00000001 / Decimal 1
stdout >>>            y: Binary 0x00010000 / Decimal 16
stdout >>> =============================================
stdout >>>        carry: Binary 0x00000000 / Decimal 0
stdout >>>            x: Binary 0x00010001 / Decimal 17
stdout >>>            y: Binary 0x00000000 / Decimal 0
stdout >>> =============================================
result >>> 17
```

`swap_pair.py`

**Swap_pair**

A function swap odd and even bits in an integer with as few instructions as possible (Ex bit and bit 1 are swapped, bit 2 and bit 3 are swapped)

For example:
* 22: `0x010110`  --> 41: `0x101001`
* 10: `0x1010`    --> 5 : `0x0101`

Workflow 

* We can approach this as operating on the odds bit first, and then the even bits.
* We can mask all odd bits with `10101010` in binary ('AA') then shift them right by 1
* Similarly, we mask all even bit with `01010101` in binary ('55') then shift them left
by 1. 
* Finally, we merge these two values by OR operation.

```python
def swap_pair(num):
    print("Swap Pair of {}".format(num))
    print()
    # odd bit arithmetic right shift 1 bit
    print("Odd Bits")
    print("{:>12s}:".format("10(AA)"), b(int('AAAAAAAA', 16)))
    print("{:>12s}:".format("n & AA"),  b(num & int('AAAAAAAA', 16)))
    print("{:>12s}:".format("n & AA >> 1"),  b((num & int('AAAAAAAA', 16)) >> 1))
    odd = (num & int('AAAAAAAA', 16)) >> 1
    print()
    # even bit left shift 1 bit
    print("Even Bits")
    print("{:>12s}:".format("01(55)"), b(int('55555555', 16)))
    print("{:>12s}:".format("n & 55"),  b(num & int('55555555', 16)))
    print("{:>12s}:".format("n & 55 << 1"),  b((num & int('55555555', 16)) << 1))
    even = (num & int('55555555', 16)) << 1
    
    print()
    print("{:>12s}:".format("odd | even"),  odd | even)
    return odd | even

swap_pair(3)

stdout >>> Swap Pair of 3
stdout >>> 
stdout >>> Odd Bits
stdout >>>       10(AA): Binary 0x10101010101010101010101010101010 / Decimal 2863311530
stdout >>>       n & AA: Binary 0x00000010 / Decimal 2
stdout >>>  n & AA >> 1: Binary 0x00000001 / Decimal 1
stdout >>> 
stdout >>> Even Bits
stdout >>>       01(55): Binary 0x1010101010101010101010101010101 / Decimal 1431655765
stdout >>>       n & 55: Binary 0x00000001 / Decimal 1
stdout >>>  n & 55 << 1: Binary 0x00000010 / Decimal 2
stdout >>> 
stdout >>>   odd | even: 3
result >>> 3
```

## Square Root

```python
def sqrt(n: float) -> float:
    result: float = 0
        
    # The second-to-top bit is set: 1 << 30 for 32 bits
    bit = 1 << 30
    print("Bit", b(bit))
    
    while bit > n :
        bit = bit >> 2
        print("Decrising bit ", b(bit))
    
    while bit != 0:
        
        if n >= result + bit :
            n -= result + bit
            print("n ", b(n))
            result = bit + (result >> 1)
            print("result_in ", b(result))
        else:
            result >>= 1
            print("result_out ", b(result))
            
        bit >>= 2
    
    return result

sqrt(25)

stdout >>> Bit Binary 0x1000000000000000000000000000000 / Decimal 1073741824
stdout >>> Decrising bit  Binary 0x10000000000000000000000000000 / Decimal 268435456
stdout >>> Decrising bit  Binary 0x100000000000000000000000000 / Decimal 67108864
stdout >>> Decrising bit  Binary 0x1000000000000000000000000 / Decimal 16777216
stdout >>> Decrising bit  Binary 0x10000000000000000000000 / Decimal 4194304
stdout >>> Decrising bit  Binary 0x100000000000000000000 / Decimal 1048576
stdout >>> Decrising bit  Binary 0x1000000000000000000 / Decimal 262144
stdout >>> Decrising bit  Binary 0x10000000000000000 / Decimal 65536
stdout >>> Decrising bit  Binary 0x100000000000000 / Decimal 16384
stdout >>> Decrising bit  Binary 0x1000000000000 / Decimal 4096
stdout >>> Decrising bit  Binary 0x10000000000 / Decimal 1024
stdout >>> Decrising bit  Binary 0x100000000 / Decimal 256
stdout >>> Decrising bit  Binary 0x01000000 / Decimal 64
stdout >>> Decrising bit  Binary 0x00010000 / Decimal 16
stdout >>> n  Binary 0x00001001 / Decimal 9
stdout >>> result_in  Binary 0x00010000 / Decimal 16
stdout >>> result_out  Binary 0x00001000 / Decimal 8
stdout >>> n  Binary 0x00000000 / Decimal 0
stdout >>> result_in  Binary 0x00000101 / Decimal 5
result >>> 5
```

```python
b(3), b(~3)

result >>> ('Binary 0x00000011 / Decimal 3', 'Binary 0x-0000100 / Decimal -4')
```

```python
def inverse(n):
    for i in range(0,8):
        if n < 1<<i: 
            s = 1<<i
            print(s, b(s))
            print(s-1, b(s-1))
            break
inverse(8)

stdout >>> 16 Binary 0x00010000 / Decimal 16
stdout >>> 15 Binary 0x00001111 / Decimal 15
```

```python
b(1)

result >>> 'Binary 0x00000001 / Decimal 1'
```

```python
from typing import List
```