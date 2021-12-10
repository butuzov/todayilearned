![image.png](numpy.png)

```python
# main lib
import numpy as np

# for additional examples
import pandas as pd
import matplotlib.pyplot as plt
```

## Creating Numpy Arrays

* `array` - Convert input data (list, tuple, array, or other sequence type) to an ndarray either by inferring a dtype or explicitly specifying a dtype. Copies the input data by default.
* `asarray` - Convert input to ndarray, but do not copy if the input is already an ndarray
* `arange` - Like the built-in range but returns an ndarray instead of a list.
* `ones`, `ones_like` - Produce an array of all 1’s with the given shape and dtype. ones_like takes another array and produces a ones array of the same shape and dtype.
* `zeros`, `zeros_like` - Like ones and ones_like but producing arrays of 0’s instead
* `empty`, `empty_like` - Create new arrays by allocating new memory, but do not populate with any values like ones and zeros
* `full`, `full_like` - Create new array with predefined value.
* `eye`, `identity` - Create a square N x N identity matrix (1’s on the diagonal and 0’s elsewhere)
* `triu` and `tril` - Upper/Lower triangle of an array.

```python
# random array
np.random.random(10)

result >>> array([0.86033867, 0.5858656 , 0.79229386, 0.23516033, 0.61764072,
result >>>        0.74007677, 0.21139272, 0.2307421 , 0.53221243, 0.48946233])
```

```python
np.random.random((10, 10))

result >>> array([[0.17253409, 0.26544304, 0.56144901, 0.5460301 , 0.18790159,
result >>>         0.90708397, 0.31426915, 0.27670116, 0.41057547, 0.73606721],
result >>>        [0.6670111 , 0.16069691, 0.09398584, 0.6743464 , 0.35933136,
result >>>         0.56667779, 0.48243537, 0.46385099, 0.34801941, 0.66958033],
result >>>        [0.44738101, 0.58125635, 0.11060476, 0.28448686, 0.22333878,
result >>>         0.96944141, 0.96337812, 0.54301328, 0.86144682, 0.82978016],
result >>>        [0.72722318, 0.2808265 , 0.8334588 , 0.63111176, 0.78041068,
result >>>         0.03375935, 0.45047302, 0.94532276, 0.72344994, 0.2198788 ],
result >>>        [0.76377835, 0.17796167, 0.86239069, 0.39807612, 0.94731645,
result >>>         0.58768407, 0.95565046, 0.25787271, 0.79055057, 0.24924649],
result >>>        [0.14921119, 0.8194624 , 0.76912482, 0.98818437, 0.73004861,
result >>>         0.29726423, 0.65267886, 0.33836298, 0.40090371, 0.83055724],
result >>>        [0.95710862, 0.3814533 , 0.70365778, 0.93266901, 0.88023762,
result >>>         0.08260247, 0.10050779, 0.16522425, 0.90702809, 0.7962819 ],
result >>>        [0.30851017, 0.62733042, 0.29266634, 0.0279146 , 0.45916205,
result >>>         0.33485201, 0.38164182, 0.14503506, 0.49564023, 0.19603016],
result >>>        [0.81451348, 0.85368043, 0.73558722, 0.67570912, 0.61958051,
result >>>         0.73358838, 0.15588746, 0.76697275, 0.66820435, 0.71293139],
result >>>        [0.97029826, 0.12966517, 0.84043236, 0.36639832, 0.82877345,
result >>>         0.09147033, 0.9329385 , 0.1174552 , 0.67307674, 0.43956926]])
```

```python
# transformation
(np.random.random((10, 10))*10).astype(int)

result >>> array([[8, 9, 2, 0, 7, 6, 7, 5, 4, 6],
result >>>        [8, 8, 4, 8, 5, 9, 5, 2, 5, 9],
result >>>        [6, 7, 6, 7, 9, 3, 1, 0, 6, 6],
result >>>        [1, 3, 2, 5, 3, 3, 7, 1, 6, 0],
result >>>        [2, 2, 8, 6, 6, 1, 0, 5, 7, 6],
result >>>        [9, 1, 1, 1, 1, 8, 5, 3, 0, 8],
result >>>        [6, 3, 3, 1, 7, 4, 5, 5, 5, 5],
result >>>        [5, 2, 2, 3, 0, 5, 3, 8, 1, 8],
result >>>        [2, 5, 0, 3, 6, 2, 3, 7, 2, 9],
result >>>        [2, 4, 4, 5, 4, 7, 1, 5, 5, 5]])
```

```python
# 10 elements from 0 to 5
np.linspace(0, 5, 10)

result >>> array([0.        , 0.55555556, 1.11111111, 1.66666667, 2.22222222,
result >>>        2.77777778, 3.33333333, 3.88888889, 4.44444444, 5.        ])
```

```python
# Regular python list
np.array([1,2,3,4])

result >>> array([1, 2, 3, 4])
```

NumPy offers several functions to create arrays with initial placeholder content
Create an array of zeros with desired shape (x,y)
* x == number of rows
* y == number of columns in array

```python
np.zeros((3,4))

result >>> array([[0., 0., 0., 0.],
result >>>        [0., 0., 0., 0.],
result >>>        [0., 0., 0., 0.]])
```

```python
np.ones((3,4))

result >>> array([[1., 1., 1., 1.],
result >>>        [1., 1., 1., 1.],
result >>>        [1., 1., 1., 1.]])
```

```python
# Using dtype in order to specify the data type
np.ones((3,4),dtype=np.int16)

result >>> array([[1, 1, 1, 1],
result >>>        [1, 1, 1, 1],
result >>>        [1, 1, 1, 1]], dtype=int16)
```

```python
# np.empty() is not the same as np.zeros()
np.empty((2,3))

result >>> array([[0., 0., 0.],
result >>>        [0., 0., 0.]])
```

```python
# we also can specify type
np.full((2,3), 4, dtype=np.double)

result >>> array([[4., 4., 4.],
result >>>        [4., 4., 4.]])
```

```python
# np.eye() creates an eyedentity matrix
np.eye(3)

result >>> array([[1., 0., 0.],
result >>>        [0., 1., 0.],
result >>>        [0., 0., 1.]])
```

```python
# `*_like` will use input array as example of shape

np.full_like(np.zeros((2,4)), np.nan)

result >>> array([[nan, nan, nan, nan],
result >>>        [nan, nan, nan, nan]])
```

To create sequences of numbers, NumPy provides a function analogous to range that returns arrays instead of lists

`arange(start, stop, step, dtype)`

```python
np.arange(2, 20, 2, dtype=np.float)

result >>> array([ 2.,  4.,  6.,  8., 10., 12., 14., 16., 18.])
```

```python
np.arange(2, 20, 2)

result >>> array([ 2,  4,  6,  8, 10, 12, 14, 16, 18])
```

```python
array_2d = np.array([(2,4,6),(3,5,7)]) 
array_2d

result >>> array([[2, 4, 6],
result >>>        [3, 5, 7]])
```

```python
array_2d.shape

result >>> (2, 3)
```

```python
# Using reshape to create n dimensional arrays
np.arange(6).reshape(3,2)

result >>> array([[0, 1],
result >>>        [2, 3],
result >>>        [4, 5]])
```

The reshape will only take arguments that multiply to the number in arange function.
For example: for arange(8), the possible combinations for reshape are `(2,4)`, `(4,2)`, `(2,2,2)`

```python
np.arange(6).reshape(6,1)

result >>> array([[0],
result >>>        [1],
result >>>        [2],
result >>>        [3],
result >>>        [4],
result >>>        [5]])
```

```python
np.zeros_like(np.arange(6).reshape(6,1))

result >>> array([[0],
result >>>        [0],
result >>>        [0],
result >>>        [0],
result >>>        [0],
result >>>        [0]])
```

```python
np.ones_like(np.arange(6).reshape(6,1))

result >>> array([[1],
result >>>        [1],
result >>>        [1],
result >>>        [1],
result >>>        [1],
result >>>        [1]])
```

```python
np.triu((np.random.random((5, 5))*10).astype(int))

result >>> array([[0, 6, 4, 8, 2],
result >>>        [0, 1, 5, 6, 8],
result >>>        [0, 0, 5, 0, 8],
result >>>        [0, 0, 0, 7, 6],
result >>>        [0, 0, 0, 0, 7]])
```

```python
np.triu((np.random.random((5, 5))*10).astype(int), 1)

result >>> array([[0, 3, 1, 9, 1],
result >>>        [0, 0, 9, 2, 8],
result >>>        [0, 0, 0, 0, 3],
result >>>        [0, 0, 0, 0, 0],
result >>>        [0, 0, 0, 0, 0]])
```

```python
np.tril((np.random.random((5, 5))*10).astype(int))

result >>> array([[6, 0, 0, 0, 0],
result >>>        [8, 9, 0, 0, 0],
result >>>        [5, 3, 7, 0, 0],
result >>>        [8, 2, 8, 9, 0],
result >>>        [6, 9, 6, 6, 1]])
```

```python
np.tril((np.random.random((5, 5))*10).astype(int), 1)

result >>> array([[6, 8, 0, 0, 0],
result >>>        [0, 3, 1, 0, 0],
result >>>        [8, 9, 1, 7, 0],
result >>>        [8, 3, 3, 0, 8],
result >>>        [0, 7, 4, 4, 3]])
```

## Printitng arrays

```python
print(np.arange(10000).reshape(100,100))

stdout >>> [[   0    1    2 ...   97   98   99]
stdout >>>  [ 100  101  102 ...  197  198  199]
stdout >>>  [ 200  201  202 ...  297  298  299]
stdout >>>  ...
stdout >>>  [9700 9701 9702 ... 9797 9798 9799]
stdout >>>  [9800 9801 9802 ... 9897 9898 9899]
stdout >>>  [9900 9901 9902 ... 9997 9998 9999]]
```

```python
np.set_printoptions(threshold = 1000)
```

```python
print(np.arange(10000).reshape(100,100))

stdout >>> [[   0    1    2 ...   97   98   99]
stdout >>>  [ 100  101  102 ...  197  198  199]
stdout >>>  [ 200  201  202 ...  297  298  299]
stdout >>>  ...
stdout >>>  [9700 9701 9702 ... 9797 9798 9799]
stdout >>>  [9800 9801 9802 ... 9897 9898 9899]
stdout >>>  [9900 9901 9902 ... 9997 9998 9999]]
```

## Arithmetic Operations

If the dimensions of two arrays are dissimilar, element-to-element operations are not possible. However, operations on arrays of non-similar shapes is still possible in NumPy, because of the broadcasting capability.  

### Elementwise operations

```python
a = np.array([10,10,10])
b = np.array([5,5,5])
```

```python
a + b

result >>> array([15, 15, 15])
```

```python
a - b

result >>> array([5, 5, 5])
```

```python
a * b

result >>> array([50, 50, 50])
```

```python
a / b

result >>> array([2., 2., 2.])
```

```python
a % 3

result >>> array([1, 1, 1])
```

```python
a < 35

result >>> array([ True,  True,  True])
```

```python
a > 25

result >>> array([False, False, False])
```

```python
a ** 2

result >>> array([100, 100, 100])
```

### dot function or method

```python
A = np.array( [[1,1], [0,1]] )
B = np.array( [[2,1], [3,4]] )

print('A:\n', A)
print('B:\n', B)

stdout >>> A:
stdout >>>  [[1 1]
stdout >>>  [0 1]]
stdout >>> B:
stdout >>>  [[2 1]
stdout >>>  [3 4]]
```

```python
# This gives the matrix multiplication
A.dot(B)

result >>> array([[5, 5],
result >>>        [3, 4]])
```

```python
np.dot(A,B)

result >>> array([[5, 5],
result >>>        [3, 4]])
```

### in-place operations

```python
# Modifying an existing array rather than create a new one
a  *= 3
a

result >>> array([30, 30, 30])
```

```python
b += a
b

result >>> array([35, 35, 35])
```

```python
### Unary Operators
```

```python
ages = np.array([12,15,18,20])
```

```python
ages.sum()

result >>> 65
```

```python
ages.min()

result >>> 12
```

```python
ages.max()

result >>> 20
```

By default, these operations apply to the array as though it were a list of numbers, regardless of its shape. However, by specifying the axis parameter you can apply an operation along the specified axis of an array

```python
numbers = np.arange(12).reshape(3,4)
numbers

result >>> array([[ 0,  1,  2,  3],
result >>>        [ 4,  5,  6,  7],
result >>>        [ 8,  9, 10, 11]])
```

Row and column operations
In a `2D` array axis `#0` represents columns. Axis `#1` refers to rows

```python
# Sum up each column
numbers.sum(axis=0)

result >>> array([12, 15, 18, 21])
```

```python
# Sum up each row
numbers.sum(axis=1)

result >>> array([ 6, 22, 38])
```

```python
# Minimum of each row
numbers.min(axis=1)

result >>> array([0, 4, 8])
```

## Input and Output 

### File System

```python
a = np.array(range(10))
b = a * 3
```

```python
# save array in txt uncompressed
np.savetxt("np_a.txt", a)
print(np.loadtxt('np_a.txt'))

stdout >>> [0. 1. 2. 3. 4. 5. 6. 7. 8. 9.]
```

```python
# save in binary format
np.save("np_a", a)
print(np.load('np_a.npy'))

stdout >>> [0 1 2 3 4 5 6 7 8 9]
```

```python
# save few arrays
np.savez("np_ab", a=a, b=b)
# loadign data from uncompressed file
data = np.load('np_ab.npz')
print(data['a'])

stdout >>> [0 1 2 3 4 5 6 7 8 9]
```

```python
# save few arrays compressed
np.savez_compressed("np_ab_z", a=a, b=b)
data = np.load('np_ab_z.npz')
print(data['a'])

stdout >>> [0 1 2 3 4 5 6 7 8 9]
```

```python
!ls np*

stdout >>> np_a.npy  np_a.txt  np_ab.npz  np_ab_z.npz
```

```python
!unlink np_a.npy
!unlink np_a.txt
!unlink np_ab.npz
!unlink np_ab_z.npz
```

### Memory

https://docs.scipy.org/doc/numpy/reference/generated/numpy.memmap.html#numpy.memmap

```python
from tempfile import mkdtemp
import os.path as path
filename = path.join(mkdtemp(), 'newfile.dat')
```

```python
fp = np.memmap(filename, dtype='float32', mode='w+', shape=(3,4))
fp

result >>> memmap([[0., 0., 0., 0.],
result >>>         [0., 0., 0., 0.],
result >>>         [0., 0., 0., 0.]], dtype=float32)
```

```python
data = np.arange(12, dtype='float32')
data.resize((3,4))
fp[:] = data[:]
fp

result >>> memmap([[ 0.,  1.,  2.,  3.],
result >>>         [ 4.,  5.,  6.,  7.],
result >>>         [ 8.,  9., 10., 11.]], dtype=float32)
```

```python
fp.filename == path.abspath(filename)
fp.filename

result >>> '/var/folders/bj/r_01hm_d1x3gkrbxn_vvrmgm0000gn/T/tmpdrz_b9d6/newfile.dat'
```

```python
del fp
```

## Universal Functions

### Trigonometric Functions

```python
angles = np.array([0,30,45,60,90])
```

Angles need to be converted to radians by multiplying by pi/180 
Only then can we appy trigonometric functions to our array.

```python
angles_radians = angles * np.pi/180
angles_radians

result >>> array([0.        , 0.52359878, 0.78539816, 1.04719755, 1.57079633])
```

```python
# Sine of angles in the array:
np.sin(angles_radians)

result >>> array([0.        , 0.5       , 0.70710678, 0.8660254 , 1.        ])
```

```python
## Alternatively, use the np.radians() function to convert to radians
angles_radians = np.radians(angles)
angles_radians

result >>> array([0.        , 0.52359878, 0.78539816, 1.04719755, 1.57079633])
```

```python
# Cosine of angles in the array:
np.cos(angles_radians)

result >>> array([1.00000000e+00, 8.66025404e-01, 7.07106781e-01, 5.00000000e-01,
result >>>        6.12323400e-17])
```

```python
# Tangent of angles in the array
np.tan(angles_radians)

result >>> array([0.00000000e+00, 5.77350269e-01, 1.00000000e+00, 1.73205081e+00,
result >>>        1.63312394e+16])
```

`arcsin`, `arcos`, and `arctan` functions return the trigonometric inverse of sin, cos, and tan of the given angle. The result of these functions can be verified by `numpy.degrees()` function by converting radians to degrees.

```python
# Compute sine inverse of angles. Returned values are in radians.

np.arcsin(np.sin(angles * np.pi/180) )

result >>> array([0.        , 0.52359878, 0.78539816, 1.04719755, 1.57079633])
```

```python
# np.degrees() converts radians to degrees
np.degrees(np.arcsin(np.sin(angles * np.pi/180) ) )

result >>> array([ 0., 30., 45., 60., 90.])
```

### Statistical Functions

```python
test_scores = np.array([32.32, 56.98, 21.52, 44.32, 
                        55.63, 13.75, 43.47, 43.34])
```

```python
# Mean test scores of the students: 
np.mean(test_scores)

result >>> 38.91625
```

```python
# Median
np.median(test_scores)

result >>> 43.405
```

```python
# We will now perform basic statistical methods on real 
# life dataset. We will use salary data of 1147 European developers.

salaries = np.genfromtxt('data/salary.csv', delimiter=',')
salaries

result >>> array([60000., 58000., 56967., ..., 54647., 25000., 70000.])
```

```python
salaries.shape

result >>> (1147,)
```

```python
print(f"Mean = {np.mean(salaries)}")
print(f"Median = {np.median(salaries)}")
print(f"Standard Deviation = {np.std(salaries)}")
print(f"Variance = {np.var(salaries)}")

stdout >>> Mean = 55894.53879686138
stdout >>> Median = 48000.0
stdout >>> Standard Deviation = 55170.37550939316
stdout >>> Variance = 3043770333.8474483
```

##  Indexing and Slicing

This also follows zero based indexing like python lists

```python
a = np.arange(11)**2
a

result >>> array([  0,   1,   4,   9,  16,  25,  36,  49,  64,  81, 100])
```

```python
a[2]

result >>> 4
```

```python
a[-2]

result >>> 81
```

```python
a[2:7]

result >>> array([ 4,  9, 16, 25, 36])
```

```python
a[2:-2]

result >>> array([ 4,  9, 16, 25, 36, 49, 64])
```

```python
a[:7]

result >>> array([ 0,  1,  4,  9, 16, 25, 36])
```

```python
a[:11:2]

result >>> array([  0,   4,  16,  36,  64, 100])
```

```python
a[::-1]

result >>> array([100,  81,  64,  49,  36,  25,  16,   9,   4,   1,   0])
```

### 2D Array indexing

Consider an array students, it contains the test scores in two courses of the students against their names

```python
students = np.array([['Alice','Beth','Cathy','Dorothy'],
                     [65,78,90,81],
                     [71,82,79,92]])
```

```python
students

result >>> array([['Alice', 'Beth', 'Cathy', 'Dorothy'],
result >>>        ['65', '78', '90', '81'],
result >>>        ['71', '82', '79', '92']], dtype='<U7')
```

```python
students[0]

result >>> array(['Alice', 'Beth', 'Cathy', 'Dorothy'], dtype='<U7')
```

```python
students[1]

result >>> array(['65', '78', '90', '81'], dtype='<U7')
```

```python
students[2]

result >>> array(['71', '82', '79', '92'], dtype='<U7')
```

```python
students[0,1]

result >>> 'Beth'
```

```python
# get data using method
students.item((0,1))

result >>> 'Beth'
```

```python
# set data using method
students.itemset((0,1), 'Michel')
students

result >>> array([['Alice', 'Michel', 'Cathy', 'Dorothy'],
result >>>        ['65', '78', '90', '81'],
result >>>        ['71', '82', '79', '92']], dtype='<U7')
```

### Searching in arrays

```python
a = np.arange(16).reshape(4,4)**2
a

result >>> array([[  0,   1,   4,   9],
result >>>        [ 16,  25,  36,  49],
result >>>        [ 64,  81, 100, 121],
result >>>        [144, 169, 196, 225]])
```

```python
np.where(a > 50)

result >>> (array([2, 2, 2, 2, 3, 3, 3, 3]), array([0, 1, 2, 3, 0, 1, 2, 3]))
```

```python
np.argwhere(a > 50)

result >>> array([[2, 0],
result >>>        [2, 1],
result >>>        [2, 2],
result >>>        [2, 3],
result >>>        [3, 0],
result >>>        [3, 1],
result >>>        [3, 2],
result >>>        [3, 3]])
```

### 2D Array slicing

This will consider the rows 0 and 1, columns 2 and 3

```python
students[0:2,2:4]

result >>> array([['Cathy', 'Dorothy'],
result >>>        ['90', '81']], dtype='<U7')
```

```python
# all rows and columns 1
students[:,1:2]

result >>> array([['Michel'],
result >>>        ['78'],
result >>>        ['82']], dtype='<U7')
```

```python
# all rows and 1,2 columns
students[:,1:3]

result >>> array([['Michel', 'Cathy'],
result >>>        ['78', '90'],
result >>>        ['82', '79']], dtype='<U7')
```

```python
# All columns, rows 0 and 1
students[0:2,:]

result >>> array([['Alice', 'Michel', 'Cathy', 'Dorothy'],
result >>>        ['65', '78', '90', '81']], dtype='<U7')
```

```python
# All rows and columns
students[:]

result >>> array([['Alice', 'Michel', 'Cathy', 'Dorothy'],
result >>>        ['65', '78', '90', '81'],
result >>>        ['71', '82', '79', '92']], dtype='<U7')
```

```python
#The last row
students[-1,:]

result >>> array(['71', '82', '79', '92'], dtype='<U7')
```

```python
#3rd from last to second from last row, last two columns
students[-3:-1,-2:]

result >>> array([['Cathy', 'Dorothy'],
result >>>        ['90', '81']], dtype='<U7')
```

### dots or ellipsis(...)

Slicing can also include ellipsis (…) to make a selection tuple of the same length as the dimension of an array. The dots (...) represent as many colons as needed to produce a complete indexing tuple

```python
students[0,...]

result >>> array(['Alice', 'Michel', 'Cathy', 'Dorothy'], dtype='<U7')
```

```python
# all rows and column 1
students[...,1]

result >>> array(['Michel', '78', '82'], dtype='<U7')
```

```python
students[...,1].shape

result >>> (3,)
```

## Iteratition

### 1D Arrays

```python
a = np.arange(11)**2
a

result >>> array([  0,   1,   4,   9,  16,  25,  36,  49,  64,  81, 100])
```

```python
for i in a:
    print(i**(1/2))

stdout >>> 0.0
stdout >>> 1.0
stdout >>> 2.0
stdout >>> 3.0
stdout >>> 4.0
stdout >>> 5.0
stdout >>> 6.0
stdout >>> 7.0
stdout >>> 8.0
stdout >>> 9.0
stdout >>> 10.0
```

### Multi-Dimensional Arrays

```python
students = np.array([['Alice','Beth','Cathy','Dorothy'],
                     [65,78,90,81],
                     [71,82,79,92]])
```

```python
# Each iteration will be over the rows of the array
for i in students:
    print('i = ', i)

stdout >>> i =  ['Alice' 'Beth' 'Cathy' 'Dorothy']
stdout >>> i =  ['65' '78' '90' '81']
stdout >>> i =  ['71' '82' '79' '92']
```

**Flatten a multi-dimensional array**

```python
# If one wants to perform an operation on each element in the array, 
# one can use the flatten function which will flatten the array to a 
# single dimension. 'By default', the flattening will occur `row-wise` 
# (also knows as C order)

for element in students.flatten():
    print(element)

stdout >>> Alice
stdout >>> Beth
stdout >>> Cathy
stdout >>> Dorothy
stdout >>> 65
stdout >>> 78
stdout >>> 90
stdout >>> 81
stdout >>> 71
stdout >>> 82
stdout >>> 79
stdout >>> 92
```

```python
# To flatten a 2D array column-wise, 
# use the Fortran order
for element in students.flatten(order='F'):
    print(element)

stdout >>> Alice
stdout >>> 65
stdout >>> 71
stdout >>> Beth
stdout >>> 78
stdout >>> 82
stdout >>> Cathy
stdout >>> 90
stdout >>> 79
stdout >>> Dorothy
stdout >>> 81
stdout >>> 92
```

```python
# nditer
# Efficient multi-dimensional iterator object to iterate over arrays
x = np.arange(12).reshape(3,4)
x

result >>> array([[ 0,  1,  2,  3],
result >>>        [ 4,  5,  6,  7],
result >>>        [ 8,  9, 10, 11]])
```

```python
# This is row-wise iteration, similar to iterating over a C-order flattened array
for i in np.nditer(x):
    print(i)

stdout >>> 0
stdout >>> 1
stdout >>> 2
stdout >>> 3
stdout >>> 4
stdout >>> 5
stdout >>> 6
stdout >>> 7
stdout >>> 8
stdout >>> 9
stdout >>> 10
stdout >>> 11
```

```python
# Fortran order
# This is like iterating over an array which has been flattened column-wise

for i in np.nditer(x, order = 'F'): 
    print(i)

stdout >>> 0
stdout >>> 4
stdout >>> 8
stdout >>> 1
stdout >>> 5
stdout >>> 9
stdout >>> 2
stdout >>> 6
stdout >>> 10
stdout >>> 3
stdout >>> 7
stdout >>> 11
```

```python
# There are a number of flags which we can pass as a list to nditer. 
# Many of these involve setting buffering options <br />
# If we want iterate over each column, we can use the flag argument with value 'external_loop'

for i in np.nditer(x, order = 'F', flags = ['external_loop']): 
    print(i)

stdout >>> [0 4 8]
stdout >>> [1 5 9]
stdout >>> [ 2  6 10]
stdout >>> [ 3  7 11]
```

### Modifying Array Values

By default, the nditer treats the input array as a read-only object. To modify the array elements, you must specify either read-write or write-only mode. This is controlled with per-operand flags.

```python
# will trhoud value error
try:
    for arr in np.nditer(x):
        arr[...] = arr * arr
except ValueError as e:
    print("Error:", e)
finally:
    print(arr)

stdout >>> Error: assignment destination is read-only
stdout >>> 0
```

```python
# We set the ops_flag to make the array read-write
try:
    for arr in np.nditer(x, op_flags = ['readwrite']):
        arr[...] = arr * arr
except ValueError as e:
    print("Error:", e)
finally:
    print(arr)

stdout >>> 121
```

## Array Shape Manipulation

```python
a = np.array([("Germany","France", "Hungary","Austria"),
              ("Berlin","Paris", "Budapest","Vienna" )])
```

```python
a

result >>> array([['Germany', 'France', 'Hungary', 'Austria'],
result >>>        ['Berlin', 'Paris', 'Budapest', 'Vienna']], dtype='<U8')
```

```python
a.shape

result >>> (2, 4)
```

The `ravel()` function

The primary functional difference is that flatten is a method of an ndarray object and hence can only be called for true numpy arrays. In contrast `ravel()` is a library-level function and hence can be called on any object that can successfully be parsed. For example `ravel()` will work on a list of ndarrays, while flatten will not.

```python
a.ravel()

result >>> array(['Germany', 'France', 'Hungary', 'Austria', 'Berlin', 'Paris',
result >>>        'Budapest', 'Vienna'], dtype='<U8')
```

```python
# transposing array
a.T

result >>> array([['Germany', 'Berlin'],
result >>>        ['France', 'Paris'],
result >>>        ['Hungary', 'Budapest'],
result >>>        ['Austria', 'Vienna']], dtype='<U8')
```

```python
a.T.ravel()

result >>> array(['Germany', 'Berlin', 'France', 'Paris', 'Hungary', 'Budapest',
result >>>        'Austria', 'Vienna'], dtype='<U8')
```

`reshape()` gives a new shape to an array without changing its data.

```python
a.shape

result >>> (2, 4)
```

```python
a.reshape(4,2)

result >>> array([['Germany', 'France'],
result >>>        ['Hungary', 'Austria'],
result >>>        ['Berlin', 'Paris'],
result >>>        ['Budapest', 'Vienna']], dtype='<U8')
```

```python
np.arange(15).reshape(3,5)

result >>> array([[ 0,  1,  2,  3,  4],
result >>>        [ 5,  6,  7,  8,  9],
result >>>        [10, 11, 12, 13, 14]])
```

```python
# Reshaping a 15-element array to an 18-element one will throw an error
try:
    np.arange(15).reshape(3,6)
except ValueError as e:
    print("Error:", e)

stdout >>> Error: cannot reshape array of size 15 into shape (3,6)
```

```python
# Specify only one dimension (and infer the others) when reshaping
# Another way we can reshape is by metioning only one dimension, and -1. 
# -1 means that the length in that dimension is inferred
```

```python
countries = np.array(["Germany","France", "Hungary","Austria","Italy","Denmark"])
countries

result >>> array(['Germany', 'France', 'Hungary', 'Austria', 'Italy', 'Denmark'],
result >>>       dtype='<U7')
```

```python
countries.reshape(-1,3)

result >>> array([['Germany', 'France', 'Hungary'],
result >>>        ['Austria', 'Italy', 'Denmark']], dtype='<U7')
```

```python
countries.reshape(3,-1)

result >>> array([['Germany', 'France'],
result >>>        ['Hungary', 'Austria'],
result >>>        ['Italy', 'Denmark']], dtype='<U7')
```

```python
countries.reshape(3,2)

result >>> array([['Germany', 'France'],
result >>>        ['Hungary', 'Austria'],
result >>>        ['Italy', 'Denmark']], dtype='<U7')
```

## Splitting Arrays

`split()` an array into multiple sub-arrays. By specifying the number of equally shaped arrays to return, or by specifying the columns after which the division should occur

`split(array, indices_or_sections, axis=0)`

```python
x = np.arange(9)
```

```python
x

result >>> array([0, 1, 2, 3, 4, 5, 6, 7, 8])
```

```python
# Split the array in 3 equal-sized subarrays:
np.split(x, 3)

result >>> [array([0, 1, 2]), array([3, 4, 5]), array([6, 7, 8])]
```

```python
# The number of splits must be a divisor of the number of elements
# Or Numpy will complain that an even split is not possible

try:
    np.split(x, 4)
except TypeError as e:
    print("Error:", e)
except ValueError as e:
    print("Error:", e)

stdout >>> Error: array split does not result in an equal division
```

```python
# Split the array at positions indicated in 1-D array:
np.split(x,[4,7])

result >>> [array([0, 1, 2, 3]), array([4, 5, 6]), array([7, 8])]
```

`hsplit`. The `numpy.hsplit` is a special case of split() function where axis is 1 indicating a horizontal split regardless of the dimension of the input array. <br />
In this example, the split will be performed along a column

```python
y = np.array([("Germany","France", "Hungary","Austria"),
              ("Berlin","Paris", "Budapest","Vienna" )])
```

```python
y

result >>> array([['Germany', 'France', 'Hungary', 'Austria'],
result >>>        ['Berlin', 'Paris', 'Budapest', 'Vienna']], dtype='<U8')
```

```python
p1, p2 = np.hsplit(y, 2)
```

```python
p1

result >>> array([['Germany', 'France'],
result >>>        ['Berlin', 'Paris']], dtype='<U8')
```

```python
p2

result >>> array([['Hungary', 'Austria'],
result >>>        ['Budapest', 'Vienna']], dtype='<U8')
```

```python
np.hsplit(y,4)

result >>> [array([['Germany'],
result >>>         ['Berlin']], dtype='<U8'), array([['France'],
result >>>         ['Paris']], dtype='<U8'), array([['Hungary'],
result >>>         ['Budapest']], dtype='<U8'), array([['Austria'],
result >>>         ['Vienna']], dtype='<U8')]
```

```python
try:
    np.hsplit(y,3)
except TypeError as e:
    print("Error:", e)
except ValueError as e:
    print("Error:", e)

stdout >>> Error: array split does not result in an equal division
```

`vsplit` splits along the vertical axis

```python
p_1,p_2 = np.vsplit(y, 2)
print(p_1, p_2)

stdout >>> [['Germany' 'France' 'Hungary' 'Austria']] [['Berlin' 'Paris' 'Budapest' 'Vienna']]
```

An alternative approach is array unpacking. In this example, we unpack the array into two variables. The array unpacks by row i.e Unpacking "unpacks" the first dimensions of an array 

```python
countries,capitals = y
print('Countries: ', countries)
print('Capitals: ' , capitals)

stdout >>> Countries:  ['Germany' 'France' 'Hungary' 'Austria']
stdout >>> Capitals:  ['Berlin' 'Paris' 'Budapest' 'Vienna']
```

```python
b1,b2,b3,b4 = y.T
print(b1,b2,b3,b4)

stdout >>> ['Germany' 'Berlin'] ['France' 'Paris'] ['Hungary' 'Budapest'] ['Austria' 'Vienna']
```

We can not use the following code, reason being the first dimension of array now contains 4 rows. If we want to split in 2 arrays horizontally we should use `split` or `hsplit`.

```python
try:
    z1,z2 = y.T
except TypeError as e:
    print("Error:", e)
except ValueError as e:
    print("Error:", e)

stdout >>> Error: too many values to unpack (expected 2)
```

## View vs Copy

 - When the contents are physically stored in another location, it is called **Copy.**
 - On the other hand, a different view of the same memory content is provided, we call it as **View.**
 - Different array objects can share the same data. NumPy has **ndarray.view()** method which is a new array object that looks at the same data of the original array.

```python
fruits = np.array(["Apple","Mango","Grapes","Watermelon"])
```

```python
basket_1 = fruits.view()
basket_2 = fruits.view()
print(basket_1)
print(basket_2)

stdout >>> ['Apple' 'Mango' 'Grapes' 'Watermelon']
stdout >>> ['Apple' 'Mango' 'Grapes' 'Watermelon']
```

```python
print("ids for the arrays are different.")
print("id for fruits is : ", id(fruits))
print("id for baskets is : ", id(basket_1), id(basket_2))

stdout >>> ids for the arrays are different.
stdout >>> id for fruits is :  4583528528
stdout >>> id for baskets is :  4583529008 4583528448
```

```python
basket_1 is fruits

result >>> False
```

```python
basket_1.base is fruits

result >>> True
```

```python
# Change a few elements of basket. It changes the elements of fruits
# Here, we assign a new value to the first element of basket_2. You 
# might be astonished that the list of fruits has been "automatically" 
# changed as well. The explanation is that there has been no new assignment 
# to basket_2, only to one of its elements.
basket_2[0] = "Strawberry"
basket_2

result >>> array(['Strawberry', 'Mango', 'Grapes', 'Watermelon'], dtype='<U10')
```

```python
fruits

result >>> array(['Strawberry', 'Mango', 'Grapes', 'Watermelon'], dtype='<U10')
```

```python
basket_1

result >>> array(['Strawberry', 'Mango', 'Grapes', 'Watermelon'], dtype='<U10')
```

```python
# Change the entire elements of basket. It does not change fruits
basket_1 = np.array(["Peach","Pineapple","Banana","Orange"])

# new np array
basket_1

result >>> array(['Peach', 'Pineapple', 'Banana', 'Orange'], dtype='<U9')
```

```python
fruits

result >>> array(['Strawberry', 'Mango', 'Grapes', 'Watermelon'], dtype='<U10')
```

```python
# Change the shape of basket. It does not change the shape of fruits
basket_2.shape = 2,2
basket_2

result >>> array([['Strawberry', 'Mango'],
result >>>        ['Grapes', 'Watermelon']], dtype='<U10')
```

```python
# fruts are unchanged
fruits

result >>> array(['Strawberry', 'Mango', 'Grapes', 'Watermelon'], dtype='<U10')
```

```python
# Slicing an array returns a view of it
mini_basket = fruits[2:]
mini_basket

result >>> array(['Grapes', 'Watermelon'], dtype='<U10')
```

```python
fruits[3] = "Peach"
mini_basket

result >>> array(['Grapes', 'Peach'], dtype='<U10')
```

## Deep Copy

We now Create a deep copy of fruits 

```python
fruits = np.array(["Apple","Mango","Grapes","Watermelon"])
```

```python
basket = fruits.copy()
```

```python
basket is fruits

result >>> False
```

```python
basket [0] = "Strawberry"
basket

result >>> array(['Strawberry', 'Mango', 'Grapes', 'Watermelon'], dtype='<U10')
```

```python
fruits

result >>> array(['Apple', 'Mango', 'Grapes', 'Watermelon'], dtype='<U10')
```

```python
basket.shape = 2,2
```

```python
# Shape of basket: 
basket

result >>> array([['Strawberry', 'Mango'],
result >>>        ['Grapes', 'Watermelon']], dtype='<U10')
```

```python
# Shape of fruits: 
fruits

result >>> array(['Apple', 'Mango', 'Grapes', 'Watermelon'], dtype='<U10')
```

## Indexing - Integer Arrays

NumPy arrays can be indexed with slices, but also with boolean or integer arrays <b>(masks)</b>. It means passing an array of indices to access multiple array elements at once. This method is called fancy indexing. It creates copies not views.

```python
a = np.arange(12)**2
a

result >>> array([  0,   1,   4,   9,  16,  25,  36,  49,  64,  81, 100, 121])
```

Suppose we want to access three different elements. We could do it like this:

```python
a[2],a[6],a[8]

result >>> (4, 36, 64)
```

Alternatively, we can pass a single list or array of indices to obtain the same result:

```python
indx_1 = [2,6,8]
a[indx_1]

result >>> array([ 4, 36, 64])
```

When using fancy indexing, the shape of the result reflects the shape of the index arrays rather than the shape of the array being indexed

```python
indx_2 = np.array([[2,4],[8,10]])
a[indx_2]

result >>> array([[  4,  16],
result >>>        [ 64, 100]])
```

```python
food = np.array([["blueberry","strawberry","cherry","blackberry"],
                 ["pinenut","hazelnuts","cashewnut","coconut"],
                 ["mustard","paprika","nutmeg","clove"]])
```

```python
food

result >>> array([['blueberry', 'strawberry', 'cherry', 'blackberry'],
result >>>        ['pinenut', 'hazelnuts', 'cashewnut', 'coconut'],
result >>>        ['mustard', 'paprika', 'nutmeg', 'clove']], dtype='<U10')
```

```python
# We will now select the corner elements of this array

row = np.array([[0,0],[2,2]])
col = np.array([[0,3],[0,3]])

food[row,col]

result >>> array([['blueberry', 'blackberry'],
result >>>        ['mustard', 'clove']], dtype='<U10')
```

```python
# Notice that the first value in the result is food[0,0], next is food[0,3] , food[2,0] and lastly food[2,3]
food[2,0]

result >>> 'mustard'
```

**Modifying Values with Fancy Indexing**

```python
food[row,col] = "000000"
food

result >>> array([['000000', 'strawberry', 'cherry', '000000'],
result >>>        ['pinenut', 'hazelnuts', 'cashewnut', 'coconut'],
result >>>        ['000000', 'paprika', 'nutmeg', '000000']], dtype='<U10')
```

```python
a

result >>> array([  0,   1,   4,   9,  16,  25,  36,  49,  64,  81, 100, 121])
```

```python
a[indx_1] = 999
```

```python
a

result >>> array([  0,   1, 999,   9,  16,  25, 999,  49, 999,  81, 100, 121])
```

### Pandas Example

We will load the GDP per capita (in US dollars) data for all the countries in the world for the year 2016 into a variable using pandas. 

**GDP per capita (current US$)**

GDP per capita is gross domestic product divided by midyear population. GDP is the sum of gross value added by all resident producers in the economy plus any product taxes and minus any subsidies not included in the value of the products. It is calculated without making deductions for depreciation of fabricated assets or for depletion and degradation of natural resources. Data are in current U.S. dollars.

```python
gdp_16 = pd.read_csv("data/gdp_pc.csv")["2016"].values
```

```python
type(gdp_16)

result >>> numpy.ndarray
```

```python
gdp_16.shape

result >>> (264,)
```

```python
plt.plot(gdp_16)
plt.show()
```

<img src='data:image/png;base64,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' />

```python
np.median(gdp_16)

stderr >>> /Users/butuzov/Desktop/jupyter/lib/python3.7/site-packages/numpy/lib/function_base.py:3405: RuntimeWarning: Invalid value encountered in median
stderr >>>   r = func(a, **kwargs)
result >>> nan
```

```python
gdp_16

result >>> array([            nan,   1944.11700491,   6454.13537039,  11540.02556108,
result >>>                    nan,  16726.72218488,  72399.65347339,  19939.93077479,
result >>>          8832.76343477,             nan,  22661.48853616,  46012.32845157,
result >>>         50551.5531752 ,  17256.62697044,    777.75285166,  46428.67142458,
result >>>          2167.64282343,   1771.01528722,   3579.75679135,  19242.62261574,
result >>>                    nan,  22516.81648228,  12172.06646607,  18066.3006624 ,
result >>>          8461.53565294,             nan,   7234.19524335,  15123.85001428,
result >>>         18064.60442722,  77420.61217204,   8900.76455309,  16956.72477025,
result >>>           698.70665573,  44819.48360027,  26851.1685408 ,  63888.73238665,
result >>>                    nan,  23193.97411052,  15529.08410645,   3693.43690912,
result >>>          3609.37559695,    801.63011974,   5717.29037095,  14153.92792519,
result >>>          1521.85721808,   6551.31894829,  16609.73935418,  15370.63489842,
result >>>                    nan,             nan,             nan,  32707.872887  ,
result >>>         34749.21236327,  48860.52529211,             nan,  10947.71518781,
result >>>         49029.01483891,  15204.93246446,  15013.29637986,  13839.96863584,
result >>>          9175.64442097,  17025.39067931,  19516.08298373,  31366.86706838,
result >>>         11242.04751604,  11128.80247428,  42063.79406454,             nan,
result >>>         36304.85427275,  29743.33742802,   1734.46407452,  39610.86680734,
result >>>          4254.34870787,  43378.14602882,   9109.95196333,  41343.29253554,
result >>>                    nan,   3508.71519196,  18102.86221919,  42656.2166022 ,
result >>>         10004.5284119 ,   4292.44889992,             nan,   1966.37968826,
result >>>          1676.86477284,   1608.70437903,  26058.077443  ,  26778.50496589,
result >>>         14200.00719381,             nan,   7944.69135314,             nan,
result >>>          7836.3791763 ,  46864.95679091,  58617.97062639,   4736.83914255,
result >>>          2270.13030428,  23422.41797653,   1783.71512959,  26700.75608264,
result >>>         12933.94330355,  10588.17884122,   3627.9834637 ,   5201.03397093,
result >>>         11609.02662172,   2834.05481866,             nan,   6570.61624581,
result >>>                    nan,  71472.29596487,  19948.81948407,  17348.93653235,
result >>>         50745.68297711,  37258.22359457,  38380.17241178,   8821.31142953,
result >>>          9047.7689234 ,  42281.18818958,  25285.94773126,   3155.11426394,
result >>>          3552.09178948,   3736.96461203,   2108.53722771,  26382.78803193,
result >>>         36532.47268398,  74263.99862606,  14732.95458354,   6549.66799975,
result >>>         14308.75112714,    812.67394381,             nan,  12952.72516857,
result >>>         15210.95079266,   2622.06907623,   1701.37258736,             nan,
result >>>         12312.94005099,   6801.29184829,  10414.49383242,   2951.02189398,
result >>>         16689.09756884,  29862.31889041, 102389.43772763,  25587.38808847,
result >>>        105420.41423718,             nan,   7857.4877165 ,             nan,
result >>>          5332.28777311,   1506.23832521,  15347.99776172,  19515.40260527,
result >>>         17274.82307137,   4022.97863242,  11442.8331081 ,  14942.19764347,
result >>>          2125.71602653,  37928.34126866,   5721.22800363,  13113.97108338,
result >>>         17633.13298468,  12252.27719639,             nan,   1216.79278744,
result >>>          3852.52590437,  21102.5589232 ,   1168.82562558,  27682.6079403 ,
result >>>         56344.96375629,  10624.92711311,             nan,    986.20696062,
result >>>          5861.08965603,   5539.82671595,  50538.6065696 ,  58790.06140389,
result >>>          2477.9033451 ,  13966.49778022,  38565.40345261,  41885.92611335,
result >>>                    nan,  25029.76382064,   5235.47808502,  23008.66536715,
result >>>         13018.60861238,   7804.16780147,  16305.48534045,   4182.53934131,
result >>>         27383.254937  ,   3775.07502387,  37740.88825606,             nan,
result >>>         30658.63199533,   9567.33912818,             nan,   5700.08893902,
result >>>         44804.46017271,             nan, 127480.48251099,  23027.28926439,
result >>>         24788.67927912,   1912.90118206,   6062.85581457,  54416.61249319,
result >>>          4730.46186883,   2566.11917856,  87832.58651422,   2235.3146905 ,
result >>>          1476.21371969,   8616.81241755,  60932.93003722,             nan,
result >>>         14514.96055305,   3721.66067336,             nan,   3723.92091102,
result >>>         22092.44125736,   3237.3925436 ,  14966.7137949 ,  30460.38408107,
result >>>         32723.07195376,  48904.55437072,   8329.56058149,             nan,
result >>>         28383.89621333,             nan,             nan,   1990.72665731,
result >>>         13986.6283938 ,  20172.0612279 ,   1490.53623573,  16913.36607526,
result >>>          2979.31023902,  16875.9877489 ,  15126.46133663,   2140.35737389,
result >>>         13214.12254515,   5745.19122926,   6062.85581457,   3723.92091102,
result >>>         32854.71905846,  11595.51237921,  25247.2017505 ,   3651.02148404,
result >>>          2786.2722427 ,   1819.43435088,   8269.61463355,  16884.35346567,
result >>>         21619.60793375,  57638.15908799,   6512.68213067,  11456.78995224,
result >>>                    nan,             nan,             nan,   6295.59058534,
result >>>          3080.56575434,  16216.92698557,   6378.25656692,  10063.75847519,
result >>>          2507.47166249,  13196.81122696,   3933.06646044,   2027.08491653])
```

```python
# The complement operator (~) can be used to remove nan elements
gdp_16 = gdp_16[~np.isnan(gdp_16)]
```

```python
gdp_16

result >>> array([  1944.11700491,   6454.13537039,  11540.02556108,  16726.72218488,
result >>>         72399.65347339,  19939.93077479,   8832.76343477,  22661.48853616,
result >>>         46012.32845157,  50551.5531752 ,  17256.62697044,    777.75285166,
result >>>         46428.67142458,   2167.64282343,   1771.01528722,   3579.75679135,
result >>>         19242.62261574,  22516.81648228,  12172.06646607,  18066.3006624 ,
result >>>          8461.53565294,   7234.19524335,  15123.85001428,  18064.60442722,
result >>>         77420.61217204,   8900.76455309,  16956.72477025,    698.70665573,
result >>>         44819.48360027,  26851.1685408 ,  63888.73238665,  23193.97411052,
result >>>         15529.08410645,   3693.43690912,   3609.37559695,    801.63011974,
result >>>          5717.29037095,  14153.92792519,   1521.85721808,   6551.31894829,
result >>>         16609.73935418,  15370.63489842,  32707.872887  ,  34749.21236327,
result >>>         48860.52529211,  10947.71518781,  49029.01483891,  15204.93246446,
result >>>         15013.29637986,  13839.96863584,   9175.64442097,  17025.39067931,
result >>>         19516.08298373,  31366.86706838,  11242.04751604,  11128.80247428,
result >>>         42063.79406454,  36304.85427275,  29743.33742802,   1734.46407452,
result >>>         39610.86680734,   4254.34870787,  43378.14602882,   9109.95196333,
result >>>         41343.29253554,   3508.71519196,  18102.86221919,  42656.2166022 ,
result >>>         10004.5284119 ,   4292.44889992,   1966.37968826,   1676.86477284,
result >>>          1608.70437903,  26058.077443  ,  26778.50496589,  14200.00719381,
result >>>          7944.69135314,   7836.3791763 ,  46864.95679091,  58617.97062639,
result >>>          4736.83914255,   2270.13030428,  23422.41797653,   1783.71512959,
result >>>         26700.75608264,  12933.94330355,  10588.17884122,   3627.9834637 ,
result >>>          5201.03397093,  11609.02662172,   2834.05481866,   6570.61624581,
result >>>         71472.29596487,  19948.81948407,  17348.93653235,  50745.68297711,
result >>>         37258.22359457,  38380.17241178,   8821.31142953,   9047.7689234 ,
result >>>         42281.18818958,  25285.94773126,   3155.11426394,   3552.09178948,
result >>>          3736.96461203,   2108.53722771,  26382.78803193,  36532.47268398,
result >>>         74263.99862606,  14732.95458354,   6549.66799975,  14308.75112714,
result >>>           812.67394381,  12952.72516857,  15210.95079266,   2622.06907623,
result >>>          1701.37258736,  12312.94005099,   6801.29184829,  10414.49383242,
result >>>          2951.02189398,  16689.09756884,  29862.31889041, 102389.43772763,
result >>>         25587.38808847, 105420.41423718,   7857.4877165 ,   5332.28777311,
result >>>          1506.23832521,  15347.99776172,  19515.40260527,  17274.82307137,
result >>>          4022.97863242,  11442.8331081 ,  14942.19764347,   2125.71602653,
result >>>         37928.34126866,   5721.22800363,  13113.97108338,  17633.13298468,
result >>>         12252.27719639,   1216.79278744,   3852.52590437,  21102.5589232 ,
result >>>          1168.82562558,  27682.6079403 ,  56344.96375629,  10624.92711311,
result >>>           986.20696062,   5861.08965603,   5539.82671595,  50538.6065696 ,
result >>>         58790.06140389,   2477.9033451 ,  13966.49778022,  38565.40345261,
result >>>         41885.92611335,  25029.76382064,   5235.47808502,  23008.66536715,
result >>>         13018.60861238,   7804.16780147,  16305.48534045,   4182.53934131,
result >>>         27383.254937  ,   3775.07502387,  37740.88825606,  30658.63199533,
result >>>          9567.33912818,   5700.08893902,  44804.46017271, 127480.48251099,
result >>>         23027.28926439,  24788.67927912,   1912.90118206,   6062.85581457,
result >>>         54416.61249319,   4730.46186883,   2566.11917856,  87832.58651422,
result >>>          2235.3146905 ,   1476.21371969,   8616.81241755,  60932.93003722,
result >>>         14514.96055305,   3721.66067336,   3723.92091102,  22092.44125736,
result >>>          3237.3925436 ,  14966.7137949 ,  30460.38408107,  32723.07195376,
result >>>         48904.55437072,   8329.56058149,  28383.89621333,   1990.72665731,
result >>>         13986.6283938 ,  20172.0612279 ,   1490.53623573,  16913.36607526,
result >>>          2979.31023902,  16875.9877489 ,  15126.46133663,   2140.35737389,
result >>>         13214.12254515,   5745.19122926,   6062.85581457,   3723.92091102,
result >>>         32854.71905846,  11595.51237921,  25247.2017505 ,   3651.02148404,
result >>>          2786.2722427 ,   1819.43435088,   8269.61463355,  16884.35346567,
result >>>         21619.60793375,  57638.15908799,   6512.68213067,  11456.78995224,
result >>>          6295.59058534,   3080.56575434,  16216.92698557,   6378.25656692,
result >>>         10063.75847519,   2507.47166249,  13196.81122696,   3933.06646044,
result >>>          2027.08491653])
```

```python
gdp_16.shape

result >>> (229,)
```

```python
np.median(gdp_16)

result >>> 13113.971083378701
```

```python
np.mean(gdp_16)

result >>> 19135.51651250454
```

The unofficial threshold for a country with a developed economy is a GDP per capita of USD 12,000. Some economists prefer to see a per capita GDP of at least USD 25,000 to be comfortable declaring a country as developed, however. Many highly developed countries, including the United States, have high per capita GDPs of USD 40,000 or above.

```python
np.count_nonzero(gdp_16[gdp_16 > 40000])

result >>> 32
```

```python
# 10 Lowest gdp per capita
np.sort(gdp_16)[0:10]

result >>> array([ 698.70665573,  777.75285166,  801.63011974,  812.67394381,
result >>>         986.20696062, 1168.82562558, 1216.79278744, 1476.21371969,
result >>>        1490.53623573, 1506.23832521])
```

```python
# 10 Highest gdp per capita
np.sort(gdp_16)[-11:-1]
# np.sort(gdp_16)[219:]

result >>> array([ 58790.06140389,  60932.93003722,  63888.73238665,  71472.29596487,
result >>>         72399.65347339,  74263.99862606,  77420.61217204,  87832.58651422,
result >>>        102389.43772763, 105420.41423718])
```

**How many countries have gdp per capita higher than Ukraine?**

```python
ukraine = gdp_16[gdp_16 > 8269.6], gdp_16[gdp_16 <= 8269.6]
```

```python
# Hiegher GDP
np.count_nonzero(ukraine[0])

result >>> 145
```

```python
# lower gdp
np.count_nonzero(ukraine[1])

result >>> 84
```

```python
# Countries between Costa Rica and Finland
np.count_nonzero((gdp_16 > 16609.7) & (gdp_16 < 43378.1))

result >>> 65
```

## Indexing - Boolean Arrays

When we index arrays with arrays of (integer) indices we are providing the list of indices to pick. With boolean indices the approach is different; we explicitly choose which items in the array we want and which ones we don’t.

```python
a = np.arange(16).reshape(4,4)
```

```python
a

result >>> array([[ 0,  1,  2,  3],
result >>>        [ 4,  5,  6,  7],
result >>>        [ 8,  9, 10, 11],
result >>>        [12, 13, 14, 15]])
```

```python
indx_bool = a > 9
```

```python
indx_bool

result >>> array([[False, False, False, False],
result >>>        [False, False, False, False],
result >>>        [False, False,  True,  True],
result >>>        [ True,  True,  True,  True]])
```

```python
a[indx_bool]

result >>> array([10, 11, 12, 13, 14, 15])
```

```python
a[a > 9]

result >>> array([10, 11, 12, 13, 14, 15])
```

```python
# filtering
a * (a > 9)

result >>> array([[ 0,  0,  0,  0],
result >>>        [ 0,  0,  0,  0],
result >>>        [ 0,  0, 10, 11],
result >>>        [12, 13, 14, 15]])
```

```python
a < 6

result >>> array([[ True,  True,  True,  True],
result >>>        [ True,  True, False, False],
result >>>        [False, False, False, False],
result >>>        [False, False, False, False]])
```

```python
np.count_nonzero(a < 6)

result >>> 6
```

```python
np.sum(a < 6)

result >>> 6
```

```python
# How many values less than 6 in each row?
np.sum(a < 6, axis=1)

result >>> array([4, 2, 0, 0])
```

```python
# Are there any values greater than 8?
np.any(a > 8)

result >>> True
```

```python
# Are all values less than 10?
np.all(a < 10)

result >>> False
```

```python
# Are all values less than 100?
np.all(a < 100)

result >>> True
```

```python
# Are all values in each row less than 9?
np.all(a < 9, axis=1)

result >>> array([ True,  True, False, False])
```

### nan's

This trick doesn't really works with float matrixes with `nan` values. You can run into RuntimeError.

## Structured Arrays

**Structured arrays** or **record arrays** are useful when you perform computations, and at the same time you could keep closely related data together. Structured arrays provide efficient storage for compound, heterogeneous data.

NumPy also provides powerful capabilities to create arrays of records, as multiple data types live in one NumPy array. However, one principle in NumPy that still needs to be honored is that the data type in each field (think of this as a column in the records) needs to be homogeneous. 

```python
name  = ["Alice","Beth","Cathy","Dorothy"]
studentId  = [1,2,3,4]
score = [85.4,90.4,87.66,78.9]
```

There's nothing here that tells us that the three arrays are related; it would be more natural if we could use a single structure to store all of this data. 

**Define the np array with the names of the 'columns' and the data format for each**
* U10 represents a 10-character Unicode string
* i4 is short for int32 (i for int, 4 for 4 bytes)
* f8 is shorthand for float64

```python
student_data = np.zeros(4, dtype={'names':('name', 'studentId', 'score'),
                          'formats':('U10', 'i4', 'f8')})
```

```python
# np.zeros() for a string sets it to an empty string
student_data

result >>> array([('', 0, 0.), ('', 0, 0.), ('', 0, 0.), ('', 0, 0.)],
result >>>       dtype=[('name', '<U10'), ('studentId', '<i4'), ('score', '<f8')])
```

```python
student_data.dtype

result >>> dtype([('name', '<U10'), ('studentId', '<i4'), ('score', '<f8')])
```

Now that we've created an empty container array, we can fill the array with our lists of values

```python
student_data['name'] = name
student_data['studentId'] = studentId
student_data['score'] = score
student_data

result >>> array([('Alice', 1, 85.4 ), ('Beth', 2, 90.4 ), ('Cathy', 3, 87.66),
result >>>        ('Dorothy', 4, 78.9 )],
result >>>       dtype=[('name', '<U10'), ('studentId', '<i4'), ('score', '<f8')])
```

The handy thing with structured arrays is that you can now refer to values either by index or by name

```python
student_data['name']

result >>> array(['Alice', 'Beth', 'Cathy', 'Dorothy'], dtype='<U10')
```

```python
student_data['studentId']

result >>> array([1, 2, 3, 4], dtype=int32)
```

```python
student_data['score']

result >>> array([85.4 , 90.4 , 87.66, 78.9 ])
```

If you index student_data at position 1 you get a structure:

```python
student_data[1]

result >>> ('Beth', 2, 90.4)
```

```python
# Get the name attribute from the last row
student_data[-1]['name']

result >>> 'Dorothy'
```

```python
# Get names where score is above 85
student_data[student_data['score'] > 85]['name']

result >>> array(['Alice', 'Beth', 'Cathy'], dtype='<U10')
```

Note that if you'd like to do any operations that are any more complicated than these, you should probably consider the Pandas package with provides a powerful data structure called data frames.

## Broadcasting

* Broadcasting is a powerful mechanism that allows numpy to work with arrays of different shapes when performing arithmetic operations. 
* Broadcasting solves the problem of arithmetic between arrays of differing shapes by in effect replicating the smaller array along the last mismatched dimension.
* NumPy operations are usually done on pairs of arrays on an element-by-element basis. In the simplest case, the two arrays must have exactly the same shape, as in the following example:

```python
a = np.array([1,2,3,4,5])
b = np.array([10,10,10,10,10])
```

```python
a * b

result >>> array([10, 20, 30, 40, 50])
```

If the dimensions of two arrays are dissimilar, element-to-element operations are not possible. However, operations on arrays of non-similar shapes is still possible in NumPy, because of the <b>broadcasting </b> capability. The smaller array is broadcast to the size of the larger array so that they have compatible shapes.

```python
a * 10

result >>> array([10, 20, 30, 40, 50])
```

```python
# Scalar and n-Dimensional Array
np.ones((4,3)) * 10

result >>> array([[10., 10., 10.],
result >>>        [10., 10., 10.],
result >>>        [10., 10., 10.],
result >>>        [10., 10., 10.]])
```

```python
# One-Dimensional and n-Dimensional Arrays
heights  = [165,170,168,183,172,169]
weights  = [61,76,56,81,62,60]
student_bio = np.array([heights,weights])
```

```python
student_bio

result >>> array([[165, 170, 168, 183, 172, 169],
result >>>        [ 61,  76,  56,  81,  62,  60]])
```

```python
factor_1 = np.array([0.0328084,2.20462 ])
```

```python
factor_1.shape

result >>> (2,)
```

```python
student_bio.shape

result >>> (2, 6)
```

**General Broadcasting Rules**

When operating on two arrays, NumPy compares their shapes element-wise. The dimensions are considered in reverse order, starting with the trailing dimensions, and working its way forward. Two dimensions are compatible when

1. they are equal
2. one of them is of size 1

**Shape mismatch**
This fails because there is a mismatch in the trailing dimensions  
* student bio:  `2` x `6` 
* factor_1: `2`

The trailing dimensions here are 6 and 2, so there is a mismatch

```python
try:
    student_bio * factor_1
except ValueError as e:
    print("Error:", e)

stdout >>> Error: operands could not be broadcast together with shapes (2,6) (2,)
```

```python
factor_2 = np.array([[0.0328084],[2.20462 ]])
```

```python
factor_2.shape

result >>> (2, 1)
```

```python
student_bio * factor_2

result >>> array([[  5.413386 ,   5.577428 ,   5.5118112,   6.0039372,   5.6430448,
result >>>           5.5446196],
result >>>        [134.48182  , 167.55112  , 123.45872  , 178.57422  , 136.68644  ,
result >>>         132.2772   ]])
```

## Automatic Reshaping
> If you're specifying 2 out of 3 dimensions, the product of the two dimensions must be a divisor of the total number of elements in the array

```python
a = np.arange(30)
```

```python
a

result >>> array([ 0,  1,  2,  3,  4,  5,  6,  7,  8,  9, 10, 11, 12, 13, 14, 15, 16,
result >>>        17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29])
```

`-1` means "whatever is needed"

```python
a.shape = 2,-1,3  
a

result >>> array([[[ 0,  1,  2],
result >>>         [ 3,  4,  5],
result >>>         [ 6,  7,  8],
result >>>         [ 9, 10, 11],
result >>>         [12, 13, 14]],
result >>> 
result >>>        [[15, 16, 17],
result >>>         [18, 19, 20],
result >>>         [21, 22, 23],
result >>>         [24, 25, 26],
result >>>         [27, 28, 29]]])
```

```python
a.shape

result >>> (2, 5, 3)
```

```python
a.shape = 2,3,-1
a

result >>> array([[[ 0,  1,  2,  3,  4],
result >>>         [ 5,  6,  7,  8,  9],
result >>>         [10, 11, 12, 13, 14]],
result >>> 
result >>>        [[15, 16, 17, 18, 19],
result >>>         [20, 21, 22, 23, 24],
result >>>         [25, 26, 27, 28, 29]]])
```

```python
a.shape

result >>> (2, 3, 5)
```

## Vector Stacking

```python
x = np.array([["Germany","France"],["Berlin","Paris"]])
y = np.array([["Hungary","Austria"],["Budapest","Vienna"]])
```

```python
x.shape

result >>> (2, 2)
```

```python
y.shape

result >>> (2, 2)
```

```python
# Joining two arrays along axis 0
np.concatenate((x,y))

result >>> array([['Germany', 'France'],
result >>>        ['Berlin', 'Paris'],
result >>>        ['Hungary', 'Austria'],
result >>>        ['Budapest', 'Vienna']], dtype='<U8')
```

```python
# Joining two arrays along axis 1 - Column-wise
np.concatenate((x,y), axis = 1)

result >>> array([['Germany', 'France', 'Hungary', 'Austria'],
result >>>        ['Berlin', 'Paris', 'Budapest', 'Vienna']], dtype='<U8')
```

```python
a = np.array([1, 2, 3])
b = np.array([2, 3, 4])
```

```python
np.stack((a, b))

result >>> array([[1, 2, 3],
result >>>        [2, 3, 4]])
```

```python
studentId = np.array([1,2,3,4])
name   = np.array(["Alice","Beth","Cathy","Dorothy"])
scores  = np.array([65,78,90,81])
```

```python
np.stack((studentId, name, scores))

result >>> array([['1', '2', '3', '4'],
result >>>        ['Alice', 'Beth', 'Cathy', 'Dorothy'],
result >>>        ['65', '78', '90', '81']], dtype='<U21')
```

```python
np.stack((studentId, name, scores)).shape

result >>> (3, 4)
```

```python
np.stack((studentId, name, scores), axis =1)

result >>> array([['1', 'Alice', '65'],
result >>>        ['2', 'Beth', '78'],
result >>>        ['3', 'Cathy', '90'],
result >>>        ['4', 'Dorothy', '81']], dtype='<U21')
```

```python
np.stack((studentId, name, scores), axis =1).shape

result >>> (4, 3)
```

```python
# Stacks row wise 

np.vstack((studentId, name, scores))

result >>> array([['1', '2', '3', '4'],
result >>>        ['Alice', 'Beth', 'Cathy', 'Dorothy'],
result >>>        ['65', '78', '90', '81']], dtype='<U21')
```

```python
# Stacks column wise
np.hstack((studentId, name, scores))

result >>> array(['1', '2', '3', '4', 'Alice', 'Beth', 'Cathy', 'Dorothy', '65',
result >>>        '78', '90', '81'], dtype='<U21')
```

```python
np.hstack((studentId, name, scores)).shape

result >>> (12,)
```

## Histograms

```python
np.histogram([1, 2, 1] , bins = [0, 1, 2, 3])

result >>> (array([0, 2, 1]), array([0, 1, 2, 3]))
```

**normal distribution**

Mean = 2  
StdDev = 0.5  
Num of data points = 1000

```python
mu, sigma = 2 , 0.5
data = np.random.normal(mu, sigma, 10000)
```

```python
data

result >>> array([1.83407739, 1.4550252 , 2.00480208, ..., 1.43580333, 3.18492252,
result >>>        2.3338091 ])
```

```python
(n, bin_edges) = np.histogram( data , bins = 50 )
```

Here, bins is an integer, it defines the number of equal-width bins in the given range (10, by default). 

Numpy computes the occurrences of input data that fall within each bin, which in turns determines the area (not necessarily the height if the bins aren't of equal width) of each bar.

* `n` == number of samples in each bin 
* `bin_edges` = The right edges of the bins

NOTE: The number of bin_edges == len(histogram)+1. This is because bin_edges[0] is the left edge of the first bin, while all other values are the right edges of each bin

```python
n

result >>> array([  1,   1,   1,   3,   3,   7,  11,   7,  16,  18,  46,  36,  59,
result >>>         93, 140, 174, 205, 264, 301, 295, 425, 442, 522, 592, 590, 544,
result >>>        603, 601, 572, 512, 470, 429, 364, 346, 270, 242, 196, 162, 129,
result >>>         91,  60,  52,  40,  14,  18,  10,   9,   3,   6,   5])
```

```python
bin_edges

result >>> array([0.04808849, 0.12204368, 0.19599886, 0.26995405, 0.34390924,
result >>>        0.41786443, 0.49181961, 0.5657748 , 0.63972999, 0.71368518,
result >>>        0.78764036, 0.86159555, 0.93555074, 1.00950592, 1.08346111,
result >>>        1.1574163 , 1.23137149, 1.30532667, 1.37928186, 1.45323705,
result >>>        1.52719224, 1.60114742, 1.67510261, 1.7490578 , 1.82301298,
result >>>        1.89696817, 1.97092336, 2.04487855, 2.11883373, 2.19278892,
result >>>        2.26674411, 2.3406993 , 2.41465448, 2.48860967, 2.56256486,
result >>>        2.63652004, 2.71047523, 2.78443042, 2.85838561, 2.93234079,
result >>>        3.00629598, 3.08025117, 3.15420636, 3.22816154, 3.30211673,
result >>>        3.37607192, 3.4500271 , 3.52398229, 3.59793748, 3.67189267,
result >>>        3.74584785])
```

```python
plt.plot(bin_edges[ 1 : ], n)
plt.show()
```

<img src='data:image/png;base64,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' />

`matplotlib` also has a function to build histograms (called hist) that differs from the one in NumPy. 

The main difference is that `pylab.hist` plots the histogram automatically, while `numpy.histogram` only generates the data.

```python
plt.hist([1, 2, 1], bins = [0, 1, 2, 3])
plt.show()
```

<img src='data:image/png;base64,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' />

```python
plt.hist(data, bins = 50)
plt.show()
```

<img src='data:image/png;base64,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' />

```python
# Bin selection can be automated
plt.hist(data, bins = 'auto')  
plt.show()
```

<img src='data:image/png;base64,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' />

## Functions and Mehtods

```python
salaries = np.genfromtxt('data/salary.csv', delimiter=',')
```

```python
salaries

result >>> array([60000., 58000., 56967., ..., 54647., 25000., 70000.])
```

```python
# Returns the indices of the maximum values along an axis.
salaries[np.argmax(salaries)]

result >>> 850000.0
```

```python
# argmin
salaries[np.argmin(salaries)]

result >>> 11400.0
```

```python
np.argsort(salaries)

result >>> array([282, 969, 606, ..., 829, 389, 246])
```

```python
s = np.array([4,6,7,82,3,4])
s.shape = 2, -1
```

```python
s

result >>> array([[ 4,  6,  7],
result >>>        [82,  3,  4]])
```

```python
np.argsort(s)

result >>> array([[0, 1, 2],
result >>>        [1, 2, 0]])
```

**specifying sorting algorithm**

```python
np.argsort(salaries,kind='mergesort')

result >>> array([282, 969, 606, ..., 829, 389, 246])
```

```python
np.argsort(salaries,kind='quicksort')

result >>> array([282, 969, 606, ..., 829, 389, 246])
```

The functions `max` , `min` , `sort` gives the respective elements itself instead of indices.

The `where()` function returns the **indices** of elements in an input array where the given condition is satisfied.

```python
np.where(salaries > 100000)

result >>> (array([  22,   27,   33,   45,   48,   59,   77,   78,   87,   91,   94,
result >>>          102,  109,  117,  123,  151,  221,  236,  242,  243,  246,  271,
result >>>          296,  297,  303,  337,  343,  378,  385,  389,  402,  408,  432,
result >>>          436,  450,  454,  503,  504,  538,  560,  563,  580,  581,  602,
result >>>          607,  687,  701,  728,  736,  745,  769,  779,  802,  819,  822,
result >>>          829,  859,  913,  922,  946,  949,  977,  980,  998, 1019, 1033,
result >>>         1062, 1065, 1104, 1105, 1112, 1130]),)
```

The `extract()` function returns the elements satisfying any condition

```python
np.extract(salaries > np.mean(salaries), salaries)

result >>> array([ 60000.,  58000.,  56967.,  70000.,  75000.,  62000.,  56000.,
result >>>        261546.,  77000., 500000.,  77570.,  60000.,  65000.,  60000.,
result >>>        120000.,  81000.,  75000.,  75000.,  73000.,  75000., 210000.,
result >>>         90000., 105227.,  84000.,  98000.,  70000., 115000.,  72000.,
result >>>         57400.,  70000.,  71796., 130773., 103529.,  57000.,  70000.,
result >>>         75000.,  57000., 119875.,  60000., 132000., 100000., 141671.,
result >>>         87500.,  90000.,  65000., 109100.,  98080.,  87689., 550000.,
result >>>         56697.,  60000., 180000.,  92000.,  63000., 220000.,  66000.,
result >>>         80000.,  65460.,  65000., 109294.,  70000.,  57758.,  75000.,
result >>>         75000.,  62662.,  64000.,  90000.,  60000.,  60000.,  70000.,
result >>>         63000.,  80000.,  70000.,  60000.,  90000., 135380.,  65000.,
result >>>         70000., 103529.,  64000.,  70000., 120000., 185000.,  70000.,
result >>>        850000.,  61500.,  60000.,  63000., 181339.,  75665.,  58000.,
result >>>         57000., 100000.,  65000.,  62662.,  90000.,  70835.,  75000.,
result >>>        111600., 102000.,  60000.,  65000., 625000., 100000.,  78000.,
result >>>         75000.,  60000.,  64000.,  63000.,  59656.,  69000.,  80750.,
result >>>         84703., 110464., 102000.,  56732.,  63000.,  62000.,  60000.,
result >>>         70151.,  60000.,  56400.,  65000., 140000.,  75000.,  65000.,
result >>>        120000., 820000.,  80000., 250000.,  58000., 130000.,  85000.,
result >>>         60000.,  65000.,  71925., 136375.,  90000., 141000.,  90000.,
result >>>         65000.,  66000.,  72000.,  65000., 142800.,  68000., 116626.,
result >>>         65000.,  56000.,  57758.,  58000.,  60000.,  60000.,  97000.,
result >>>         67382.,  90000.,  85000.,  75000.,  57000.,  60000.,  76600.,
result >>>         60000., 110000., 130000.,  61398.,  60000.,  65000.,  70000.,
result >>>        100000.,  59000.,  57000.,  60000.,  76000.,  60000.,  88000.,
result >>>        444000.,  59000.,  70000.,  72000., 100000.,  70000., 116000.,
result >>>         65000., 110000.,  70000.,  82000.,  56000.,  60000.,  72000.,
result >>>         60000.,  70000., 113996., 160000.,  60000.,  81733.,  66000.,
result >>>         80000.,  66000., 150000.,  56000., 108000.,  62031., 100000.,
result >>>         60000.,  68000.,  60000.,  60000.,  60000.,  90000.,  61700.,
result >>>         75000.,  77000.,  96000.,  60000.,  86000.,  68000.,  63000.,
result >>>         65000.,  60000.,  90000.,  90000.,  80000.,  63000.,  80000.,
result >>>         70915., 120000.,  73500.,  60000.,  60000., 120000.,  56000.,
result >>>         84396.,  95000.,  94704.,  71000.,  70000.,  60000., 220000.,
result >>>        117701.,  57758.,  80000.,  65000.,  72000.,  70000., 120000.,
result >>>         60000.,  60000.,  63500.,  60000.,  59640.,  61027.,  70000.,
result >>>         61800., 120010.,  83500.,  75000.,  72000.,  84000., 110000.,
result >>>         60005.,  59500.,  70518.,  66000.,  58000.,  70000.,  99500.,
result >>>         66858., 103645.,  77000.,  56000.,  62000.,  60000.,  57400.,
result >>>         92735.,  58000.,  69000., 102000.,  68000., 108977.,  75653.,
result >>>         68000.,  59938., 680000.,  56967.,  57600.,  75000.,  60000.,
result >>>         60000.,  58000., 152568.,  98324.,  65000.,  63605.,  82000.,
result >>>         57000.,  62000.,  85000.,  60000.,  62500.,  68656.,  87280.,
result >>>         60000.,  57400.,  76000.,  63207.,  59000., 114800.,  60000.,
result >>>         72000., 150000.,  70000.,  60000.,  57000.,  70000.,  59938.,
result >>>        111841., 110000.,  63000.,  63207.,  59000.,  64000.,  58000.,
result >>>         75000.,  75000.,  70500.,  61000.,  60555., 108000.,  60000.,
result >>>         80000., 102000., 100000.,  64800.,  62000., 134449.,  75000.,
result >>>         70000.,  65000.,  94000.,  65000.,  60000., 120000.,  80000.,
result >>>         95000.,  65000., 101349.,  74000.,  70000.,  76289.,  95000.,
result >>>         78000.,  72000.,  92000.,  65000., 120000.,  72000., 250000.,
result >>>         68500.,  62000.,  81000.,  70000.,  84000.,  56000.,  87689.,
result >>>        175000., 261546.,  59000.,  56000., 120000.,  82000.,  70000.,
result >>>         71380.,  62000., 100000.,  80000., 103000.,  57000.,  76800.,
result >>>         75000.,  90000.,  70000.])
```

## Vectorization

```python
def addsubtract(a,b):
    "Return a-b if a > b, otherwise return a + b"
    if a > b:
        return a - b
    else:
        return a + b
```

```python
addsubtract(5,7)

result >>> 12
```

```python
vec_addsubtract = np.vectorize(addsubtract)
```

```python
vec_addsubtract([0,3,6,9],[1,3,5,7])

result >>> array([1, 6, 1, 2])
```

```python
# wrong shape
try:
    vec_addsubtract([0,3,6,9],[1,3,5,7,8])
except ValueError as e:
    print("Error:", e)

stdout >>> Error: operands could not be broadcast together with shapes (4,) (5,)
```

## Tips and tricks

* https://arogozhnikov.github.io/2015/09/29/NumpyTipsAndTricks1.html
* https://arogozhnikov.github.io/2015/09/30/NumpyTipsAndTricks2.html
* https://jayrambhia.com/notes/numpy-tricks
* http://anie.me/numpy-tricks/

### Computing order of elements in array

```python
data = np.random.random(10)
np.argsort(np.argsort(data))

result >>> array([9, 6, 8, 0, 3, 7, 2, 4, 1, 5])
```

```python
# scipy function which does the same, but it's more general and faster, so prefer using it:
from scipy.stats import rankdata
rankdata(data) - 1

result >>> array([9., 6., 8., 0., 3., 7., 2., 4., 1., 5.])
```

### IronTransform (flattener of distribution)

Sometimes you need to write monotonic tranformation, which turns one distribution into uniform.

This method is useful to compare distributions or to work with distributions with heavy tails or strange shape.

```python
class IronTransform:
    def fit(self, data, weights):
        weights = weights / weights.sum()
        sorter = np.argsort(data)
        self.x = data[sorter]
        self.y = np.cumsum(weights[sorter])
        return self
        
    def transform(self, data):
        return np.interp(data, self.x, self.y)
    
    
sig_pred = np.random.normal(size=10000) + 1
bck_pred = np.random.normal(size=10000) - 1 

plt.figure()
plt.hist(sig_pred, bins=30, alpha=0.5)
plt.hist(bck_pred, bins=30, alpha=0.5)
plt.show()
```

<img src='data:image/png;base64,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' />

```python
iron = IronTransform().fit(sig_pred, weights=np.ones(len(sig_pred)))

plt.figure()
plt.hist(iron.transform(sig_pred), bins=30, alpha=0.5)
plt.hist(iron.transform(bck_pred), bins=30, alpha=0.5)
plt.show()
```

<img src='data:image/png;base64,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' />

**flatnonzero**

```python
## flattennonzero
np.flatnonzero(np.arange(10).reshape(2,5))

result >>> array([1, 2, 3, 4, 5, 6, 7, 8, 9])
```

```python
np.flatnonzero(np.arange(10) % 2 == 0 )

result >>> array([0, 2, 4, 6, 8])
```