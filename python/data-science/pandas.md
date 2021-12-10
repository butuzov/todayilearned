----

![https://pandas.pydata.org/](pandas.png)

----

```python
import pandas as pd
import numpy as np

import matplotlib.pyplot as plt
%matplotlib inline

# Just So we can log 
# nice logging
from icecream import ic
import sys, re
def printError(e): print("Error: {}".format(e), file=sys.stderr)
    
def jupyter(*args): 
    print(*[re.sub(r",\s{1,}", ", ", i.replace(",\n", ", ")) for i in args], file=sys.stdout)
    
ic.configureOutput(prefix='ic> ', outputFunction=jupyter)

rnd = np.random.RandomState(42)
```

# Pandas

_**Talks**_:
    
* "[How to use pandas the wrong way](https://www.youtube.com/watch?v=4JwpDGrMsJE)", by Pietro Battiston  @EuroPython 2017 (
[code](https://pietrobattiston.it/python:pycon)
)
* "[Neat Analytics with Pandas Indexes](https://www.youtube.com/watch?v=WjEVJ4x2SXE&)", by Alexander Hendorf @EuroPython 2017 (
  [slides](https://www.slideshare.net/PoleSystematicParisRegion/neat-analytics-with-pandas-indexes-alexander-hendorf)
)
* "[Data Wrangling & Visualization with Pandas and Jupyter](https://github.com/alanderex/pydata-pandas-workshop)", by Alexander Hendorf @EuroPython 2018

*  "[Pandas for Data Analysis](https://www.youtube.com/watch?v=oGzU688xCUs)" by Daniel Chen, @SciPy 2017 (
  [code](https://github.com/chendaniely/scipy-2017-tutorial-pandas)
)

_**Docks:**_

* https://pandas.pydata.org/pandas-docs/stable/reference/

_**Books:**_

* [Python for Data Analysis](http://shop.oreilly.com/product/0636920023784.do)
* [Python Data Science](http://shop.oreilly.com/product/0636920034919.do)

## Getting Started

`pandas` allow us plot, explore, data. perform math on large sets of it.

### Versions and general overview

```python
# version
ic(pd.__version__)

stdout >>> ic> pd.__version__: '0.24.2'
result >>> '0.24.2'
```

```python
economics = pd.read_csv('data/economics.csv', index_col='date',parse_dates=True)
economics.head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
  </tbody>
</table>
</div>

### Options

https://pandas.pydata.org/pandas-docs/stable/user_guide/options.html#available-options

```python
ic(pd.describe_option("display.max_rows"))

with pd.option_context('display.max_rows', 2):
    ic(pd.get_option("display.max_rows"))
ic(pd.get_option("display.max_rows"))
print("----"*20)
ic(pd.get_option("display.max_rows"))
ic(pd.set_option("display.max_rows", 2))
ic(pd.get_option("display.max_rows"))
print("----"*20)
ic(pd.reset_option("display.max_rows"))
ic(pd.get_option("display.max_rows"))

stdout >>> display.max_rows : int
stdout >>>     If max_rows is exceeded, switch to truncate view. Depending on
stdout >>>     `large_repr`, objects are either centrally truncated or printed as
stdout >>>     a summary view. 'None' value means unlimited.
stdout >>> 
stdout >>>     In case python/IPython is running in a terminal and `large_repr`
stdout >>>     equals 'truncate' this can be set to 0 and pandas will auto-detect
stdout >>>     the height of the terminal and print a truncated object which fits
stdout >>>     the screen height. The IPython notebook, IPython qtconsole, or
stdout >>>     IDLE do not run in a terminal and hence it is not possible to do
stdout >>>     correct auto-detection.
stdout >>>     [default: 60] [currently: 60]
stdout >>> 
stdout >>> 
stdout >>> ic> pd.describe_option("display.max_rows"): None
stdout >>> ic> pd.get_option("display.max_rows"): 2
stdout >>> ic> pd.get_option("display.max_rows"): 60
stdout >>> --------------------------------------------------------------------------------
stdout >>> ic> pd.get_option("display.max_rows"): 60
stdout >>> ic> pd.set_option("display.max_rows", 2): None
stdout >>> ic> pd.get_option("display.max_rows"): 2
stdout >>> --------------------------------------------------------------------------------
stdout >>> ic> pd.reset_option("display.max_rows"): None
stdout >>> ic> pd.get_option("display.max_rows"): 60
result >>> 60
```

## DataTypes: Series and DataFrames

Types:
* `np.bool` (`bool`) - Stored as a single byte.
* `np.int` (`int`) - Defaulted to 64 bits, Unsigned ints is alaso available (`np.uint`)
* `np.float` (`float`) - Defaulted to 64 bits.
* `np.complex` (`complex`) - Rarely seen in DA
* `np.object` (`O`, `object`) - Typically strings but is a catch- all for columns with multiple different types or other Python objects (tuples, lists, dicts, and so on).
* `np.datetime64`, `pd.Timestamp` (`datetime64`) - Specific moment in time with nanosecond precision.
* `np.timedelta64`,`pd.Timedelta` (`timedelta64`) - An amount of time, from days to nanoseconds.
* `pd.Categorical` (`Categorical`) - Specific only to pandas. Useful for object columns with relatively few unique values.

```python
ic(economics.get_dtype_counts())

print("-"*60)

ic(economics.dtypes)

stdout >>> ic> economics.get_dtype_counts(): float64    3
stdout >>>                                   int64      2
stdout >>>                                   dtype: int64
stdout >>> ------------------------------------------------------------
stdout >>> ic> economics.dtypes: pce         float64
stdout >>>                       pop           int64
stdout >>>                       psavert     float64
stdout >>>                       uempmed     float64
stdout >>>                       unemploy      int64
stdout >>>                       dtype: object
result >>> pce         float64
result >>> pop           int64
result >>> psavert     float64
result >>> uempmed     float64
result >>> unemploy      int64
result >>> dtype: object
```

### `pd.Series`

`numpy` array with labels

![image.png](attachment:image.png)

```python
rnd_series = pd.Series(rnd.randint(0, 10, 6))

ic(type(rnd_series))
ic(rnd_series.shape)
ic(rnd_series.index)
ic(rnd_series.value_counts(normalize=True))
ic(rnd_series.hasnans)

stdout >>> ic> type(rnd_series): <class 'pandas.core.series.Series'>
stdout >>> ic> rnd_series.shape: (6,)
stdout >>> ic> rnd_series.index: RangeIndex(start=0, stop=6, step=1)
stdout >>> ic> rnd_series.value_counts(normalize=True): 6    0.333333
stdout >>>                                              7    0.166667
stdout >>>                                              4    0.166667
stdout >>>                                              3    0.166667
stdout >>>                                              9    0.166667
stdout >>>                                              dtype: float64
stdout >>> ic> rnd_series.hasnans: False
result >>> False
```

### `pd.DataFrame`

![image.png](attachment:image.png)

```python
economics = pd.read_csv('data/economics.csv', index_col='date',parse_dates=True)
economics.head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
  </tbody>
</table>
</div>

```python
# row
ic(economics.values[0])
ic(type(economics.values))
ic(economics.shape)
ic(economics.columns[0:4])
ic(economics.index[0])

stdout >>> ic> economics.values[0]: array([5.07400e+02, 1.98712e+05, 1.25000e+01, 4.50000e+00, 2.94400e+03])
stdout >>> ic> type(economics.values): <class 'numpy.ndarray'>
stdout >>> ic> economics.shape: (574, 5)
stdout >>> ic> economics.columns[0:4]: Index(['pce', 'pop', 'psavert', 'uempmed'], dtype='object')
stdout >>> ic> economics.index[0]: Timestamp('1967-07-01 00:00:00')
result >>> Timestamp('1967-07-01 00:00:00')
```

```python
economics.info()

stdout >>> <class 'pandas.core.frame.DataFrame'>
stdout >>> DatetimeIndex: 574 entries, 1967-07-01 to 2015-04-01
stdout >>> Data columns (total 5 columns):
stdout >>> pce         574 non-null float64
stdout >>> pop         574 non-null int64
stdout >>> psavert     574 non-null float64
stdout >>> uempmed     574 non-null float64
stdout >>> unemploy    574 non-null int64
stdout >>> dtypes: float64(3), int64(2)
stdout >>> memory usage: 26.9 KB
```

## Data: Generate Data (`numpy`)

`pandas` use `numpy` to generate indexes, and perform operations on matrixes

### Random data in Dataframe/Series

```python
# generating random data based on numpy generated array
pd.DataFrame(rnd.randint(0, 10, 6).reshape(-1, 6)).head()
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>0</th>
      <th>1</th>
      <th>2</th>
      <th>3</th>
      <th>4</th>
      <th>5</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>2</td>
      <td>6</td>
      <td>7</td>
      <td>4</td>
      <td>3</td>
      <td>7</td>
    </tr>
  </tbody>
</table>
</div>

## Data: Import and Export

* https://pandas.pydata.org/pandas-docs/stable/reference/io.html

### Import
We can import data from: `csv`, `excel`,  `excel`, `hdfs`, etc... 

```python
# Read a comma-separated values (csv) file into DataFrame.
economics = pd.read_csv('data/economics.csv', index_col='date', parse_dates=True)
economics.head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
    <tr>
      <th>1967-09-01</th>
      <td>516.3</td>
      <td>199113</td>
      <td>11.7</td>
      <td>4.6</td>
      <td>2958</td>
    </tr>
  </tbody>
</table>
</div>

### Export

```python
# excel export depends on openpyxl
!pip install openpyxl 2>&1 1>/dev/null
```

```python
# Write DataFrame to a comma-separated values (csv) file.
economics.to_csv(path_or_buf="data.csv", index=False)

# and actual saving dataframe to excel
economics.to_excel("data.xlsx")
```

```python
!unlink data.xlsx
!unlink data.csv
```

### Export to data structures

```python
dfdict = economics.to_dict()
ic(len(dfdict))
ic(dfdict.keys())

stdout >>> ic> len(dfdict): 5
stdout >>> ic> dfdict.keys(): dict_keys(['pce', 'pop', 'psavert', 'uempmed', 'unemploy'])
result >>> dict_keys(['pce', 'pop', 'psavert', 'uempmed', 'unemploy'])
```

## Data: Indexing

### `pd.Index`

```python
# manualy created index
indexes=pd.Index([0,1,2,3,4,5,6])

ic(type(indexes))
ic(indexes.size)
ic(indexes.shape) 
ic(indexes.dtype)

stdout >>> ic> type(indexes): <class 'pandas.core.indexes.numeric.Int64Index'>
stdout >>> ic> indexes.size: 7
stdout >>> ic> indexes.shape: (7,)
stdout >>> ic> indexes.dtype: dtype('int64')
result >>> dtype('int64')
```

Create index and sort index on existing dataframe

```python
economics.set_index('unemploy').sort_index().head(5)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
    </tr>
    <tr>
      <th>unemploy</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>2685</th>
      <td>577.2</td>
      <td>201621</td>
      <td>10.9</td>
      <td>4.4</td>
    </tr>
    <tr>
      <th>2686</th>
      <td>568.8</td>
      <td>201095</td>
      <td>10.4</td>
      <td>4.6</td>
    </tr>
    <tr>
      <th>2689</th>
      <td>572.3</td>
      <td>201290</td>
      <td>10.6</td>
      <td>4.8</td>
    </tr>
    <tr>
      <th>2692</th>
      <td>589.5</td>
      <td>201881</td>
      <td>9.4</td>
      <td>4.9</td>
    </tr>
    <tr>
      <th>2709</th>
      <td>544.6</td>
      <td>200208</td>
      <td>12.2</td>
      <td>4.6</td>
    </tr>
  </tbody>
</table>
</div>

and reset index

```python
economics.set_index('unemploy').reset_index().head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>unemploy</th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>2944</td>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
    </tr>
    <tr>
      <th>1</th>
      <td>2945</td>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
    </tr>
  </tbody>
</table>
</div>

### `pd.DatetimeIndex`

https://docs.scipy.org/doc/numpy/reference/arrays.datetime.html

```python
# we can use numpy generated arrys for indexes or data
np.arange('1993-01-01', '1993-01-20', dtype="datetime64[W]")

result >>> array(['1992-12-31', '1993-01-07'], dtype='datetime64[W]')
```

```python
# date generation
pd.date_range("2001-01-01", periods=3, freq="w")

result >>> DatetimeIndex(['2001-01-07', '2001-01-14', '2001-01-21'], dtype='datetime64[ns]', freq='W-SUN')
```

```python
# autoguessing date format
pd.to_datetime(['1/2/2018', 'Jan 04, 2018'])

result >>> DatetimeIndex(['2018-01-02', '2018-01-04'], dtype='datetime64[ns]', freq=None)
```

```python
# providing dateformat
pd.to_datetime(['2/1/2018', '6/1/2018'], format="%d/%m/%Y")

result >>> DatetimeIndex(['2018-01-02', '2018-01-06'], dtype='datetime64[ns]', freq=None)
```

```python
# bussines days 
pd.date_range("2018-01-02", periods=3, freq='B')

result >>> DatetimeIndex(['2018-01-02', '2018-01-03', '2018-01-04'], dtype='datetime64[ns]', freq='B')
```

Parsing dates

```python
economics = pd.read_csv('data/economics.csv',  parse_dates=True)
economics.head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>date</th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>1967-07-01</td>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1</th>
      <td>1967-08-01</td>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
    <tr>
      <th>2</th>
      <td>1967-09-01</td>
      <td>516.3</td>
      <td>199113</td>
      <td>11.7</td>
      <td>4.6</td>
      <td>2958</td>
    </tr>
  </tbody>
</table>
</div>

```python
pd.to_datetime(economics['date'], format='%Y-%m-%d').dt.strftime("%Y")[1:4]

result >>> 1    1967
result >>> 2    1967
result >>> 3    1967
result >>> Name: date, dtype: object
```

post load datetime transformation

```python
economics = pd.read_csv('data/economics.csv')
economics['date'] = pd.to_datetime(economics['date'])
economics.set_index('date',inplace=True)
economics.head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
    <tr>
      <th>1967-09-01</th>
      <td>516.3</td>
      <td>199113</td>
      <td>11.7</td>
      <td>4.6</td>
      <td>2958</td>
    </tr>
  </tbody>
</table>
</div>

### Labels

```python
tmp = pd.Series(['Oleg', 'Developer'], index=['person', 'who'])

#  loc allow to request data by named index
# iloc allow to request data by numeric index 
ic(tmp.loc['person'])
ic(tmp.iloc[1])

stdout >>> ic> tmp.loc['person']: 'Oleg'
stdout >>> ic> tmp.iloc[1]: 'Developer'
result >>> 'Developer'
```

```python
tmp = pd.Series(range(26), index=[x for x in 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'])

ic(tmp[3:9].values)
ic(tmp["D":"I"].values)
ic(tmp.iloc[3:9].values)

stdout >>> ic> tmp[3:9].values: array([3, 4, 5, 6, 7, 8])
stdout >>> ic> tmp["D":"I"].values: array([3, 4, 5, 6, 7, 8])
stdout >>> ic> tmp.iloc[3:9].values: array([3, 4, 5, 6, 7, 8])
result >>> array([3, 4, 5, 6, 7, 8])
```

re-labeling / re-indexing

```python
# reindexing
tmp.index = [x for x in 'GATTACAHIJKLMNOPQRSTUVWXYZ']
ic(tmp.loc['G'])

# Requesting non uniq values
try:
    tmp.loc['G':'A']
except KeyError as e:
    printError(e)

stderr >>> Error: "Cannot get right slice bound for non-unique label: 'A'"
stdout >>> ic> tmp.loc['G']: 0
```

### Index Based Access

```python
economics = pd.read_csv('data/economics.csv', index_col=["date"], parse_dates=True)
economics.head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
    <tr>
      <th>1967-09-01</th>
      <td>516.3</td>
      <td>199113</td>
      <td>11.7</td>
      <td>4.6</td>
      <td>2958</td>
    </tr>
  </tbody>
</table>
</div>

```python
# access top data using column name
economics["unemploy"].head(2)

result >>> date
result >>> 1967-07-01    2944
result >>> 1967-08-01    2945
result >>> Name: unemploy, dtype: int64
```

```python
# or using column as property
economics.unemploy.head(2)

result >>> date
result >>> 1967-07-01    2944
result >>> 1967-08-01    2945
result >>> Name: unemploy, dtype: int64
```

```python
economics.iloc[[0], :]
economics.iloc[[0,2,5,4], [0,1,2]]
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
    </tr>
    <tr>
      <th>1967-09-01</th>
      <td>516.3</td>
      <td>199113</td>
      <td>11.7</td>
    </tr>
    <tr>
      <th>1967-12-01</th>
      <td>525.8</td>
      <td>199657</td>
      <td>12.1</td>
    </tr>
    <tr>
      <th>1967-11-01</th>
      <td>518.1</td>
      <td>199498</td>
      <td>12.5</td>
    </tr>
  </tbody>
</table>
</div>

```python
economics[['pce', 'pop']].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
    </tr>
  </tbody>
</table>
</div>

`loc` based indexing

```python
# ranges and selections
economics.loc['1967-07-01':'1967-09-01', ['pop', 'pce']].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pop</th>
      <th>pce</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>198712</td>
      <td>507.4</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>198911</td>
      <td>510.5</td>
    </tr>
  </tbody>
</table>
</div>

```python
# ranges and ALL
economics.loc['1967-07-01':'1967-09-01', :].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
  </tbody>
</table>
</div>

```python
# ranges and ranges
economics.loc['1967-07-01':'1967-09-01', 'pce':'psavert'].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
    </tr>
  </tbody>
</table>
</div>

`iloc` based indexes

```python
# ranges and ALL
economics.iloc[0:5,:].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
  </tbody>
</table>
</div>

```python
# ranges and selection
economics.iloc[0:5,[0,3,4]].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
  </tbody>
</table>
</div>

```python
# ranges and ranges
economics.iloc[0:5, 2:4].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>psavert</th>
      <th>uempmed</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>12.5</td>
      <td>4.5</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>12.5</td>
      <td>4.7</td>
    </tr>
  </tbody>
</table>
</div>

```python
# combining different aproaches
pd.concat([
    economics.iloc[0:1, 2:4],
    economics.iloc[6:7, 2:4],
    economics.loc['1978-07-01':'1978-09-01', ['psavert','uempmed']],
]).head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>psavert</th>
      <th>uempmed</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>12.5</td>
      <td>4.5</td>
    </tr>
    <tr>
      <th>1968-01-01</th>
      <td>11.7</td>
      <td>5.1</td>
    </tr>
    <tr>
      <th>1978-07-01</th>
      <td>10.3</td>
      <td>5.8</td>
    </tr>
  </tbody>
</table>
</div>

```python
# combining preselected columns with loc.
economics[['pce', 'unemploy']].loc['1967-07-01':'1969-07-01'].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>2945</td>
    </tr>
  </tbody>
</table>
</div>

```python
# combining preselected columns with loc.
economics[['pce', 'unemploy']].iloc[10:14].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1968-05-01</th>
      <td>550.4</td>
      <td>2740</td>
    </tr>
    <tr>
      <th>1968-06-01</th>
      <td>556.8</td>
      <td>2938</td>
    </tr>
  </tbody>
</table>
</div>

### Conditional Indexes

```python
economics.loc[((economics.pce <= 630) & (economics.pce >= 600)), ['psavert', 'unemploy']].head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>psavert</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1969-05-01</th>
      <td>10.0</td>
      <td>2713</td>
    </tr>
    <tr>
      <th>1969-06-01</th>
      <td>10.9</td>
      <td>2816</td>
    </tr>
  </tbody>
</table>
</div>

### Labeled Indexes (2)

```python
size = 10
tmp = pd.DataFrame(rnd.randint(0, 10, size**2).reshape(size, -1), 
                  index=[f"R{x:02d}" for x in range(size)],
                  columns = [f"C{x:02d}" for x in range(size)])

tmp
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>C00</th>
      <th>C01</th>
      <th>C02</th>
      <th>C03</th>
      <th>C04</th>
      <th>C05</th>
      <th>C06</th>
      <th>C07</th>
      <th>C08</th>
      <th>C09</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>R00</th>
      <td>7</td>
      <td>2</td>
      <td>5</td>
      <td>4</td>
      <td>1</td>
      <td>7</td>
      <td>5</td>
      <td>1</td>
      <td>4</td>
      <td>0</td>
    </tr>
    <tr>
      <th>R01</th>
      <td>9</td>
      <td>5</td>
      <td>8</td>
      <td>0</td>
      <td>9</td>
      <td>2</td>
      <td>6</td>
      <td>3</td>
      <td>8</td>
      <td>2</td>
    </tr>
    <tr>
      <th>R02</th>
      <td>4</td>
      <td>2</td>
      <td>6</td>
      <td>4</td>
      <td>8</td>
      <td>6</td>
      <td>1</td>
      <td>3</td>
      <td>8</td>
      <td>1</td>
    </tr>
    <tr>
      <th>R03</th>
      <td>9</td>
      <td>8</td>
      <td>9</td>
      <td>4</td>
      <td>1</td>
      <td>3</td>
      <td>6</td>
      <td>7</td>
      <td>2</td>
      <td>0</td>
    </tr>
    <tr>
      <th>R04</th>
      <td>3</td>
      <td>1</td>
      <td>7</td>
      <td>3</td>
      <td>1</td>
      <td>5</td>
      <td>5</td>
      <td>9</td>
      <td>3</td>
      <td>5</td>
    </tr>
    <tr>
      <th>R05</th>
      <td>1</td>
      <td>9</td>
      <td>1</td>
      <td>9</td>
      <td>3</td>
      <td>7</td>
      <td>6</td>
      <td>8</td>
      <td>7</td>
      <td>4</td>
    </tr>
    <tr>
      <th>R06</th>
      <td>1</td>
      <td>4</td>
      <td>7</td>
      <td>9</td>
      <td>8</td>
      <td>8</td>
      <td>0</td>
      <td>8</td>
      <td>6</td>
      <td>8</td>
    </tr>
    <tr>
      <th>R07</th>
      <td>7</td>
      <td>0</td>
      <td>7</td>
      <td>7</td>
      <td>2</td>
      <td>0</td>
      <td>7</td>
      <td>2</td>
      <td>2</td>
      <td>0</td>
    </tr>
    <tr>
      <th>R08</th>
      <td>4</td>
      <td>9</td>
      <td>6</td>
      <td>9</td>
      <td>8</td>
      <td>6</td>
      <td>8</td>
      <td>7</td>
      <td>1</td>
      <td>0</td>
    </tr>
    <tr>
      <th>R09</th>
      <td>6</td>
      <td>6</td>
      <td>7</td>
      <td>4</td>
      <td>2</td>
      <td>7</td>
      <td>5</td>
      <td>2</td>
      <td>0</td>
      <td>2</td>
    </tr>
  </tbody>
</table>
</div>

```python
# rows not found.
tmp['C05':'C09']
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>C00</th>
      <th>C01</th>
      <th>C02</th>
      <th>C03</th>
      <th>C04</th>
      <th>C05</th>
      <th>C06</th>
      <th>C07</th>
      <th>C08</th>
      <th>C09</th>
    </tr>
  </thead>
  <tbody>
  </tbody>
</table>
</div>

```python
# but as column its ok!
tmp['C05']

result >>> R00    7
result >>> R01    2
result >>> R02    6
result >>> R03    3
result >>> R04    5
result >>> R05    7
result >>> R06    8
result >>> R07    0
result >>> R08    6
result >>> R09    7
result >>> Name: C05, dtype: int64
```

```python
tmp['R05':'R06']
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>C00</th>
      <th>C01</th>
      <th>C02</th>
      <th>C03</th>
      <th>C04</th>
      <th>C05</th>
      <th>C06</th>
      <th>C07</th>
      <th>C08</th>
      <th>C09</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>R05</th>
      <td>1</td>
      <td>9</td>
      <td>1</td>
      <td>9</td>
      <td>3</td>
      <td>7</td>
      <td>6</td>
      <td>8</td>
      <td>7</td>
      <td>4</td>
    </tr>
    <tr>
      <th>R06</th>
      <td>1</td>
      <td>4</td>
      <td>7</td>
      <td>9</td>
      <td>8</td>
      <td>8</td>
      <td>0</td>
      <td>8</td>
      <td>6</td>
      <td>8</td>
    </tr>
  </tbody>
</table>
</div>

### MultiIndex

```python
pdlen = 20

tmp = pd.DataFrame(
    {
        'city'     : [x for x in ['Paris', 'London', 'Berlin', 'Manchester', 'Kyiv']*10][:pdlen],
        'category' : rnd.randint(0, 7, pdlen),
        'price'    : rnd.randint(10, 300, pdlen),
        'rating'   : rnd.randint(0, 5, pdlen),
    }
)

tmp['country'] = tmp['city'].map({
    'Paris':'FR', 
    'London':'UK', 
    'Berlin':'DE', 
    'Manchester':'US', 
    'Kyiv':'UA',
})

tmp.head(5)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>city</th>
      <th>category</th>
      <th>price</th>
      <th>rating</th>
      <th>country</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>Paris</td>
      <td>4</td>
      <td>72</td>
      <td>4</td>
      <td>FR</td>
    </tr>
    <tr>
      <th>1</th>
      <td>London</td>
      <td>6</td>
      <td>240</td>
      <td>0</td>
      <td>UK</td>
    </tr>
    <tr>
      <th>2</th>
      <td>Berlin</td>
      <td>5</td>
      <td>250</td>
      <td>4</td>
      <td>DE</td>
    </tr>
    <tr>
      <th>3</th>
      <td>Manchester</td>
      <td>2</td>
      <td>61</td>
      <td>3</td>
      <td>US</td>
    </tr>
    <tr>
      <th>4</th>
      <td>Kyiv</td>
      <td>0</td>
      <td>105</td>
      <td>3</td>
      <td>UA</td>
    </tr>
  </tbody>
</table>
</div>

```python
tmp = tmp.groupby(['country', 'city', 'category']).mean()
tmp
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th></th>
      <th></th>
      <th>price</th>
      <th>rating</th>
    </tr>
    <tr>
      <th>country</th>
      <th>city</th>
      <th>category</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th rowspan="4" valign="top">DE</th>
      <th rowspan="4" valign="top">Berlin</th>
      <th>0</th>
      <td>22.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th>3</th>
      <td>293.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th>5</th>
      <td>250.0</td>
      <td>4.0</td>
    </tr>
    <tr>
      <th>6</th>
      <td>246.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th rowspan="3" valign="top">FR</th>
      <th rowspan="3" valign="top">Paris</th>
      <th>1</th>
      <td>252.0</td>
      <td>2.0</td>
    </tr>
    <tr>
      <th>4</th>
      <td>151.5</td>
      <td>3.5</td>
    </tr>
    <tr>
      <th>6</th>
      <td>38.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th rowspan="4" valign="top">UA</th>
      <th rowspan="4" valign="top">Kyiv</th>
      <th>0</th>
      <td>105.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th>2</th>
      <td>179.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th>5</th>
      <td>180.0</td>
      <td>1.0</td>
    </tr>
    <tr>
      <th>6</th>
      <td>196.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th rowspan="3" valign="top">UK</th>
      <th rowspan="3" valign="top">London</th>
      <th>1</th>
      <td>167.5</td>
      <td>1.5</td>
    </tr>
    <tr>
      <th>2</th>
      <td>45.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th>6</th>
      <td>240.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th rowspan="3" valign="top">US</th>
      <th rowspan="3" valign="top">Manchester</th>
      <th>2</th>
      <td>61.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th>4</th>
      <td>75.0</td>
      <td>4.0</td>
    </tr>
    <tr>
      <th>6</th>
      <td>160.5</td>
      <td>1.0</td>
    </tr>
  </tbody>
</table>
</div>

```python
# show all indexes and levels
ic(tmp.index)
ic(tmp.index.levels)
print("-"*70)
ic(tmp.index.names)
ic(tmp.index.values)
print("-"*70)
ic(tmp.index.get_level_values(2))
print("-"*70)
ic(tmp.loc["UA"])
print("-"*70)
ic(tmp.loc["UA", "Kyiv"].max())
ic(tmp.loc["UA", "Kyiv", 0:4])

stdout >>> ic> tmp.index: MultiIndex(levels=[['DE', 'FR', 'UA', 'UK', 'US'], ['Berlin', 'Kyiv', 'London', 'Manchester', 'Paris'], [0, 1, 2, 3, 4, 5, 6]], codes=[[0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 4, 4], [0, 0, 0, 0, 4, 4, 4, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3], [0, 3, 5, 6, 1, 4, 6, 0, 2, 5, 6, 1, 2, 6, 2, 4, 6]], names=['country', 'city', 'category'])
stdout >>> ic> tmp.index.levels: FrozenList([['DE', 'FR', 'UA', 'UK', 'US'], ['Berlin', 'Kyiv', 'London', 'Manchester', 'Paris'], [0, 1, 2, 3, 4, 5, 6]])
stdout >>> ----------------------------------------------------------------------
stdout >>> ic> tmp.index.names: FrozenList(['country', 'city', 'category'])
stdout >>> ic> tmp.index.values: array([('DE', 'Berlin', 0), ('DE', 'Berlin', 3), ('DE', 'Berlin', 5), ('DE', 'Berlin', 6), ('FR', 'Paris', 1), ('FR', 'Paris', 4), ('FR', 'Paris', 6), ('UA', 'Kyiv', 0), ('UA', 'Kyiv', 2), ('UA', 'Kyiv', 5), ('UA', 'Kyiv', 6), ('UK', 'London', 1), ('UK', 'London', 2), ('UK', 'London', 6), ('US', 'Manchester', 2), ('US', 'Manchester', 4), ('US', 'Manchester', 6)], dtype=object)
stdout >>> ----------------------------------------------------------------------
stdout >>> ic> tmp.index.get_level_values(2): Int64Index([0, 3, 5, 6, 1, 4, 6, 0, 2, 5, 6, 1, 2, 6, 2, 4, 6], dtype='int64', name='category')
stdout >>> ----------------------------------------------------------------------
stdout >>> ic> tmp.loc["UA"]:                price  rating
stdout >>>                    city category               
stdout >>>                    Kyiv 0         105.0     3.0
stdout >>>                         2         179.0     0.0
stdout >>>                         5         180.0     1.0
stdout >>>                         6         196.0     0.0
stdout >>> ----------------------------------------------------------------------
stdout >>> ic> tmp.loc["UA", "Kyiv"].max(): price     196.0
stdout >>>                                  rating      3.0
stdout >>>                                  dtype: float64
stdout >>> ic> tmp.loc["UA", "Kyiv", 0:4]:                        price  rating
stdout >>>                                 country city category               
stdout >>>                                 UA      Kyiv 0         105.0     3.0
stdout >>>                                              2         179.0     0.0
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th></th>
      <th></th>
      <th>price</th>
      <th>rating</th>
    </tr>
    <tr>
      <th>country</th>
      <th>city</th>
      <th>category</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th rowspan="2" valign="top">UA</th>
      <th rowspan="2" valign="top">Kyiv</th>
      <th>0</th>
      <td>105.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th>2</th>
      <td>179.0</td>
      <td>0.0</td>
    </tr>
  </tbody>
</table>
</div>

```python
# without inplace=True it will return new dataframe
tmp.rename(index={'UA':'ЮА'}, columns={'price':'Precio'}, inplace=True)
tmp
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th></th>
      <th></th>
      <th>Precio</th>
      <th>rating</th>
    </tr>
    <tr>
      <th>country</th>
      <th>city</th>
      <th>category</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th rowspan="4" valign="top">DE</th>
      <th rowspan="4" valign="top">Berlin</th>
      <th>0</th>
      <td>22.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th>3</th>
      <td>293.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th>5</th>
      <td>250.0</td>
      <td>4.0</td>
    </tr>
    <tr>
      <th>6</th>
      <td>246.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th rowspan="3" valign="top">FR</th>
      <th rowspan="3" valign="top">Paris</th>
      <th>1</th>
      <td>252.0</td>
      <td>2.0</td>
    </tr>
    <tr>
      <th>4</th>
      <td>151.5</td>
      <td>3.5</td>
    </tr>
    <tr>
      <th>6</th>
      <td>38.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th rowspan="4" valign="top">ЮА</th>
      <th rowspan="4" valign="top">Kyiv</th>
      <th>0</th>
      <td>105.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th>2</th>
      <td>179.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th>5</th>
      <td>180.0</td>
      <td>1.0</td>
    </tr>
    <tr>
      <th>6</th>
      <td>196.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th rowspan="3" valign="top">UK</th>
      <th rowspan="3" valign="top">London</th>
      <th>1</th>
      <td>167.5</td>
      <td>1.5</td>
    </tr>
    <tr>
      <th>2</th>
      <td>45.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th>6</th>
      <td>240.0</td>
      <td>0.0</td>
    </tr>
    <tr>
      <th rowspan="3" valign="top">US</th>
      <th rowspan="3" valign="top">Manchester</th>
      <th>2</th>
      <td>61.0</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th>4</th>
      <td>75.0</td>
      <td>4.0</td>
    </tr>
    <tr>
      <th>6</th>
      <td>160.5</td>
      <td>1.0</td>
    </tr>
  </tbody>
</table>
</div>

### `pd.Index` as Sets

```python
# intersection
ic(pd.Index([1,2,3,4]) & pd.Index([3,4,5,6]))

# union
ic(pd.Index([1,2,3,4]) | pd.Index([3,4,5,6]))

# symetric difference
ic(pd.Index([1,2,3,4]) ^ pd.Index([3,4,5,6]))

stdout >>> ic> pd.Index([1,2,3,4]) & pd.Index([3,4,5,6]): Int64Index([3, 4], dtype='int64')
stdout >>> ic> pd.Index([1,2,3,4]) | pd.Index([3,4,5,6]): Int64Index([1, 2, 3, 4, 5, 6], dtype='int64')
stdout >>> ic> pd.Index([1,2,3,4]) ^ pd.Index([3,4,5,6]): Int64Index([1, 2, 5, 6], dtype='int64')
result >>> Int64Index([1, 2, 5, 6], dtype='int64')
```

### Index Aligment

```python
# source data area and population
data_area = { 'Alaska': 1723337,  'Texas': 695662, 'California': 423967}
data_ppls = { 'California': 38332521, 'Texas': 26448193, 'New York': 19651127}

area = pd.Series(data_area, name='area')
ppls = pd.Series(data_ppls, name='ppls')

# Index aligment
ic(ppls/area)

stdout >>> ic> ppls/area: Alaska              NaN
stdout >>>                California    90.413926
stdout >>>                New York            NaN
stdout >>>                Texas         38.018740
stdout >>>                dtype: float64
result >>> Alaska              NaN
result >>> California    90.413926
result >>> New York            NaN
result >>> Texas         38.018740
result >>> dtype: float64
```

One more example

```python
A = pd.Series([2, 4, 6], index=[0, 1, 2]) 
B = pd.Series([1, 3, 5], index=[1, 2, 3]) 

ic(A+B)

stdout >>> ic> A+B: 0    NaN
stdout >>>          1    5.0
stdout >>>          2    9.0
stdout >>>          3    NaN
stdout >>>          dtype: float64
result >>> 0    NaN
result >>> 1    5.0
result >>> 2    9.0
result >>> 3    NaN
result >>> dtype: float64
```

using functions methods (`.add()` instead `+` ) we can fill nan values

```python
# using function we can fill nan values
ic(A.add(B, fill_value=100))

stdout >>> ic> A.add(B, fill_value=100): 0    102.0
stdout >>>                               1      5.0
stdout >>>                               2      9.0
stdout >>>                               3    105.0
stdout >>>                               dtype: float64
result >>> 0    102.0
result >>> 1      5.0
result >>> 2      9.0
result >>> 3    105.0
result >>> dtype: float64
```

```python
economics = pd.read_csv('data/economics.csv', index_col='date',parse_dates=True)

ic(economics.unemploy.idxmin())
ic(economics.unemploy.min())
ic(np.argmin(economics.unemploy))
print("-"*60)
ic(economics.unemploy.idxmax())
ic(economics.unemploy.max())
ic(np.argmax(economics.unemploy))

stderr >>> /Users/butuzov/venvs/jupyter/lib/python3.7/site-packages/numpy/core/fromnumeric.py:56: FutureWarning: 
stderr >>> The current behaviour of 'Series.argmin' is deprecated, use 'idxmin'
stderr >>> instead.
stderr >>> The behavior of 'argmin' will be corrected to return the positional
stderr >>> minimum in the future. For now, use 'series.values.argmin' or
stderr >>> 'np.argmin(np.array(values))' to get the position of the minimum
stderr >>> row.
stderr >>>   return getattr(obj, method)(*args, **kwds)
stderr >>> /Users/butuzov/venvs/jupyter/lib/python3.7/site-packages/numpy/core/fromnumeric.py:56: FutureWarning: 
stderr >>> The current behaviour of 'Series.argmax' is deprecated, use 'idxmax'
stderr >>> instead.
stderr >>> The behavior of 'argmax' will be corrected to return the positional
stderr >>> maximum in the future. For now, use 'series.values.argmax' or
stderr >>> 'np.argmax(np.array(values))' to get the position of the maximum
stderr >>> row.
stderr >>>   return getattr(obj, method)(*args, **kwds)
stdout >>> ic> economics.unemploy.idxmin(): Timestamp('1968-12-01 00:00:00')
stdout >>> ic> economics.unemploy.min(): 2685
stdout >>> ic> np.argmin(economics.unemploy): Timestamp('1968-12-01 00:00:00')
stdout >>> ------------------------------------------------------------
stdout >>> ic> economics.unemploy.idxmax(): Timestamp('2009-10-01 00:00:00')
stdout >>> ic> economics.unemploy.max(): 15352
stdout >>> ic> np.argmax(economics.unemploy): Timestamp('2009-10-01 00:00:00')
result >>> Timestamp('2009-10-01 00:00:00')
```

## Data: Handling Data

### Handling Missing Data

```python
titanic = pd.read_csv('data/titanic.csv', keep_default_na=False)
titanic['age'][4:7].head()

result >>> 4    35.0
result >>> 5        
result >>> 6    54.0
result >>> Name: age, dtype: object
```

```python
# Additional strings to recognize as NA/NaN. 
# If dict passed, specific per-column NA values. 
# By default the following values are interpreted as 
# NaN: ‘’, ‘#N/A’, ‘#N/A N/A’, ‘#NA’, ‘-1.#IND’, ‘-1.#QNAN’, ‘-NaN’, 
# ‘-nan’, ‘1.#IND’, ‘1.#QNAN’, ‘N/A’, ‘NA’, ‘NULL’, ‘NaN’, ‘n/a’, ‘nan’, ‘null’.
titanic = pd.read_csv('data/titanic.csv', na_values=[35.0])
titanic['age'][4:7].head()

result >>> 4     NaN
result >>> 5     NaN
result >>> 6    54.0
result >>> Name: age, dtype: float64
```

```python
titanic = pd.read_csv('data/titanic.csv', keep_default_na=True)
titanic['age'][4:7].head()

result >>> 4    35.0
result >>> 5     NaN
result >>> 6    54.0
result >>> Name: age, dtype: float64
```

```python
ic(titanic['age'].hasnans)
ic(titanic['age'].isnull().sum())
ic(titanic['age'].dropna().hasnans)
ic(titanic['age'].dropna().isnull().sum())
ic(titanic['age'].notnull().sum())
ic(np.count_nonzero(titanic['age'].isnull()))

stdout >>> ic> titanic['age'].hasnans: True
stdout >>> ic> titanic['age'].isnull().sum(): 177
stdout >>> ic> titanic['age'].dropna().hasnans: False
stdout >>> ic> titanic['age'].dropna().isnull().sum(): 0
stdout >>> ic> titanic['age'].notnull().sum(): 714
stdout >>> ic> np.count_nonzero(titanic['age'].isnull()): 177
result >>> 177
```

```python
titanic.age.value_counts(dropna=True).plot.pie()
```

<img src='data:image/png;base64,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' />

getting rid of `nan`s

```python
ages = pd.DataFrame()

ages['fillna(0)']      = titanic['age'].fillna(0)
ages['fillna - ffill'] = titanic['age'].fillna(method='ffill')
ages['fillna - bfill'] = titanic['age'].fillna(method='bfill')
ages['dropna']         = titanic['age'].dropna(inplace=False )
ages['interpolate']    = titanic['age'].interpolate()
ages['min']            = titanic['age'].fillna(titanic['age'].min())
ages['mean']           = titanic['age'].fillna(titanic['age'].mean())
ages[4:7]
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>fillna(0)</th>
      <th>fillna - ffill</th>
      <th>fillna - bfill</th>
      <th>dropna</th>
      <th>interpolate</th>
      <th>min</th>
      <th>mean</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>4</th>
      <td>35.0</td>
      <td>35.0</td>
      <td>35.0</td>
      <td>35.0</td>
      <td>35.0</td>
      <td>35.00</td>
      <td>35.000000</td>
    </tr>
    <tr>
      <th>5</th>
      <td>0.0</td>
      <td>35.0</td>
      <td>54.0</td>
      <td>NaN</td>
      <td>44.5</td>
      <td>0.42</td>
      <td>29.699118</td>
    </tr>
    <tr>
      <th>6</th>
      <td>54.0</td>
      <td>54.0</td>
      <td>54.0</td>
      <td>54.0</td>
      <td>54.0</td>
      <td>54.00</td>
      <td>54.000000</td>
    </tr>
  </tbody>
</table>
</div>

### Melting/Mergin/Concating Data

```python
titanic = pd.read_csv('data/titanic.csv', keep_default_na=True)
titanic.head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>survived</th>
      <th>pclass</th>
      <th>sex</th>
      <th>age</th>
      <th>sibsp</th>
      <th>parch</th>
      <th>fare</th>
      <th>embarked</th>
      <th>class</th>
      <th>who</th>
      <th>adult_male</th>
      <th>deck</th>
      <th>embark_town</th>
      <th>alive</th>
      <th>alone</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>0</td>
      <td>3</td>
      <td>male</td>
      <td>22.0</td>
      <td>1</td>
      <td>0</td>
      <td>7.2500</td>
      <td>S</td>
      <td>Third</td>
      <td>man</td>
      <td>True</td>
      <td>NaN</td>
      <td>Southampton</td>
      <td>no</td>
      <td>False</td>
    </tr>
    <tr>
      <th>1</th>
      <td>1</td>
      <td>1</td>
      <td>female</td>
      <td>38.0</td>
      <td>1</td>
      <td>0</td>
      <td>71.2833</td>
      <td>C</td>
      <td>First</td>
      <td>woman</td>
      <td>False</td>
      <td>C</td>
      <td>Cherbourg</td>
      <td>yes</td>
      <td>False</td>
    </tr>
    <tr>
      <th>2</th>
      <td>1</td>
      <td>3</td>
      <td>female</td>
      <td>26.0</td>
      <td>0</td>
      <td>0</td>
      <td>7.9250</td>
      <td>S</td>
      <td>Third</td>
      <td>woman</td>
      <td>False</td>
      <td>NaN</td>
      <td>Southampton</td>
      <td>yes</td>
      <td>True</td>
    </tr>
  </tbody>
</table>
</div>

Renaming columns and indexes

```python
titanic.rename(columns={'sex':'gender'}, inplace=True)
titanic.head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>survived</th>
      <th>pclass</th>
      <th>gender</th>
      <th>age</th>
      <th>sibsp</th>
      <th>parch</th>
      <th>fare</th>
      <th>embarked</th>
      <th>class</th>
      <th>who</th>
      <th>adult_male</th>
      <th>deck</th>
      <th>embark_town</th>
      <th>alive</th>
      <th>alone</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>0</td>
      <td>3</td>
      <td>male</td>
      <td>22.0</td>
      <td>1</td>
      <td>0</td>
      <td>7.2500</td>
      <td>S</td>
      <td>Third</td>
      <td>man</td>
      <td>True</td>
      <td>NaN</td>
      <td>Southampton</td>
      <td>no</td>
      <td>False</td>
    </tr>
    <tr>
      <th>1</th>
      <td>1</td>
      <td>1</td>
      <td>female</td>
      <td>38.0</td>
      <td>1</td>
      <td>0</td>
      <td>71.2833</td>
      <td>C</td>
      <td>First</td>
      <td>woman</td>
      <td>False</td>
      <td>C</td>
      <td>Cherbourg</td>
      <td>yes</td>
      <td>False</td>
    </tr>
    <tr>
      <th>2</th>
      <td>1</td>
      <td>3</td>
      <td>female</td>
      <td>26.0</td>
      <td>0</td>
      <td>0</td>
      <td>7.9250</td>
      <td>S</td>
      <td>Third</td>
      <td>woman</td>
      <td>False</td>
      <td>NaN</td>
      <td>Southampton</td>
      <td>yes</td>
      <td>True</td>
    </tr>
  </tbody>
</table>
</div>

droping columns

```python
titanic.drop('embarked', axis=1, errors="ignore", inplace=True)
titanic.drop('who', axis=1, errors="ignore", inplace=True)
titanic.drop(['adult_male', 'alone'], axis=1, errors="ignore", inplace=True) 
titanic.drop(['parch', 'deck'], axis=1, errors="ignore", inplace=True)
titanic.head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>survived</th>
      <th>pclass</th>
      <th>gender</th>
      <th>age</th>
      <th>sibsp</th>
      <th>fare</th>
      <th>class</th>
      <th>embark_town</th>
      <th>alive</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>0</td>
      <td>3</td>
      <td>male</td>
      <td>22.0</td>
      <td>1</td>
      <td>7.2500</td>
      <td>Third</td>
      <td>Southampton</td>
      <td>no</td>
    </tr>
    <tr>
      <th>1</th>
      <td>1</td>
      <td>1</td>
      <td>female</td>
      <td>38.0</td>
      <td>1</td>
      <td>71.2833</td>
      <td>First</td>
      <td>Cherbourg</td>
      <td>yes</td>
    </tr>
    <tr>
      <th>2</th>
      <td>1</td>
      <td>3</td>
      <td>female</td>
      <td>26.0</td>
      <td>0</td>
      <td>7.9250</td>
      <td>Third</td>
      <td>Southampton</td>
      <td>yes</td>
    </tr>
  </tbody>
</table>
</div>

```python
# fares 
titanic = titanic.groupby(['class', 'embark_town', 'gender'])['age'].mean().reset_index()
titanic
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>class</th>
      <th>embark_town</th>
      <th>gender</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>First</td>
      <td>Cherbourg</td>
      <td>female</td>
      <td>36.052632</td>
    </tr>
    <tr>
      <th>1</th>
      <td>First</td>
      <td>Cherbourg</td>
      <td>male</td>
      <td>40.111111</td>
    </tr>
    <tr>
      <th>2</th>
      <td>First</td>
      <td>Queenstown</td>
      <td>female</td>
      <td>33.000000</td>
    </tr>
    <tr>
      <th>3</th>
      <td>First</td>
      <td>Queenstown</td>
      <td>male</td>
      <td>44.000000</td>
    </tr>
    <tr>
      <th>4</th>
      <td>First</td>
      <td>Southampton</td>
      <td>female</td>
      <td>32.704545</td>
    </tr>
    <tr>
      <th>5</th>
      <td>First</td>
      <td>Southampton</td>
      <td>male</td>
      <td>41.897188</td>
    </tr>
    <tr>
      <th>6</th>
      <td>Second</td>
      <td>Cherbourg</td>
      <td>female</td>
      <td>19.142857</td>
    </tr>
    <tr>
      <th>7</th>
      <td>Second</td>
      <td>Cherbourg</td>
      <td>male</td>
      <td>25.937500</td>
    </tr>
    <tr>
      <th>8</th>
      <td>Second</td>
      <td>Queenstown</td>
      <td>female</td>
      <td>30.000000</td>
    </tr>
    <tr>
      <th>9</th>
      <td>Second</td>
      <td>Queenstown</td>
      <td>male</td>
      <td>57.000000</td>
    </tr>
    <tr>
      <th>10</th>
      <td>Second</td>
      <td>Southampton</td>
      <td>female</td>
      <td>29.719697</td>
    </tr>
    <tr>
      <th>11</th>
      <td>Second</td>
      <td>Southampton</td>
      <td>male</td>
      <td>30.875889</td>
    </tr>
    <tr>
      <th>12</th>
      <td>Third</td>
      <td>Cherbourg</td>
      <td>female</td>
      <td>14.062500</td>
    </tr>
    <tr>
      <th>13</th>
      <td>Third</td>
      <td>Cherbourg</td>
      <td>male</td>
      <td>25.016800</td>
    </tr>
    <tr>
      <th>14</th>
      <td>Third</td>
      <td>Queenstown</td>
      <td>female</td>
      <td>22.850000</td>
    </tr>
    <tr>
      <th>15</th>
      <td>Third</td>
      <td>Queenstown</td>
      <td>male</td>
      <td>28.142857</td>
    </tr>
    <tr>
      <th>16</th>
      <td>Third</td>
      <td>Southampton</td>
      <td>female</td>
      <td>23.223684</td>
    </tr>
    <tr>
      <th>17</th>
      <td>Third</td>
      <td>Southampton</td>
      <td>male</td>
      <td>26.574766</td>
    </tr>
  </tbody>
</table>
</div>

pivoting data

```python
titanic.pivot_table(index=['class','gender'],columns='embark_town',values='age').reset_index().head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th>embark_town</th>
      <th>class</th>
      <th>gender</th>
      <th>Cherbourg</th>
      <th>Queenstown</th>
      <th>Southampton</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>First</td>
      <td>female</td>
      <td>36.052632</td>
      <td>33.0</td>
      <td>32.704545</td>
    </tr>
    <tr>
      <th>1</th>
      <td>First</td>
      <td>male</td>
      <td>40.111111</td>
      <td>44.0</td>
      <td>41.897188</td>
    </tr>
    <tr>
      <th>2</th>
      <td>Second</td>
      <td>female</td>
      <td>19.142857</td>
      <td>30.0</td>
      <td>29.719697</td>
    </tr>
  </tbody>
</table>
</div>

melting data

```python
pd.melt(titanic, id_vars=['class', 'gender']).head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>class</th>
      <th>gender</th>
      <th>variable</th>
      <th>value</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>First</td>
      <td>female</td>
      <td>embark_town</td>
      <td>Cherbourg</td>
    </tr>
    <tr>
      <th>1</th>
      <td>First</td>
      <td>male</td>
      <td>embark_town</td>
      <td>Cherbourg</td>
    </tr>
    <tr>
      <th>2</th>
      <td>First</td>
      <td>female</td>
      <td>embark_town</td>
      <td>Queenstown</td>
    </tr>
  </tbody>
</table>
</div>

```python
# todo - cover 
# unstack()
# stack()
```

### Common Math operations over the DataFrames/Series

```python
economics = pd.read_csv('data/economics.csv', index_col='date',parse_dates=True)
```

```python
economics.info()

stdout >>> <class 'pandas.core.frame.DataFrame'>
stdout >>> DatetimeIndex: 574 entries, 1967-07-01 to 2015-04-01
stdout >>> Data columns (total 5 columns):
stdout >>> pce         574 non-null float64
stdout >>> pop         574 non-null int64
stdout >>> psavert     574 non-null float64
stdout >>> uempmed     574 non-null float64
stdout >>> unemploy    574 non-null int64
stdout >>> dtypes: float64(3), int64(2)
stdout >>> memory usage: 26.9 KB
```

```python
economics.describe()
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>count</th>
      <td>574.000000</td>
      <td>574.000000</td>
      <td>574.000000</td>
      <td>574.000000</td>
      <td>574.000000</td>
    </tr>
    <tr>
      <th>mean</th>
      <td>4843.510453</td>
      <td>257189.381533</td>
      <td>7.936585</td>
      <td>8.610105</td>
      <td>7771.557491</td>
    </tr>
    <tr>
      <th>std</th>
      <td>3579.287206</td>
      <td>36730.801593</td>
      <td>3.124394</td>
      <td>4.108112</td>
      <td>2641.960571</td>
    </tr>
    <tr>
      <th>min</th>
      <td>507.400000</td>
      <td>198712.000000</td>
      <td>1.900000</td>
      <td>4.000000</td>
      <td>2685.000000</td>
    </tr>
    <tr>
      <th>25%</th>
      <td>1582.225000</td>
      <td>224896.000000</td>
      <td>5.500000</td>
      <td>6.000000</td>
      <td>6284.000000</td>
    </tr>
    <tr>
      <th>50%</th>
      <td>3953.550000</td>
      <td>253060.000000</td>
      <td>7.700000</td>
      <td>7.500000</td>
      <td>7494.000000</td>
    </tr>
    <tr>
      <th>75%</th>
      <td>7667.325000</td>
      <td>290290.750000</td>
      <td>10.500000</td>
      <td>9.100000</td>
      <td>8691.000000</td>
    </tr>
    <tr>
      <th>max</th>
      <td>12161.500000</td>
      <td>320887.000000</td>
      <td>17.000000</td>
      <td>25.200000</td>
      <td>15352.000000</td>
    </tr>
  </tbody>
</table>
</div>

```python
economics.pce.mean()

result >>> 4843.510452961673
```

```python
economics.pce = economics.pce * 2
economics.head()
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>1014.8</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>1021.0</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
    <tr>
      <th>1967-09-01</th>
      <td>1032.6</td>
      <td>199113</td>
      <td>11.7</td>
      <td>4.6</td>
      <td>2958</td>
    </tr>
    <tr>
      <th>1967-10-01</th>
      <td>1025.8</td>
      <td>199311</td>
      <td>12.5</td>
      <td>4.9</td>
      <td>3143</td>
    </tr>
    <tr>
      <th>1967-11-01</th>
      <td>1036.2</td>
      <td>199498</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>3066</td>
    </tr>
  </tbody>
</table>
</div>

```python
cs = economics.pce.cumsum()
ic(cs.head(1))
ic(cs.tail(1))

stdout >>> ic> cs.head(1): date
stdout >>>                 1967-07-01    1014.8
stdout >>>                 Name: pce, dtype: float64
stdout >>> ic> cs.tail(1): date
stdout >>>                 2015-04-01    5560350.0
stdout >>>                 Name: pce, dtype: float64
result >>> date
result >>> 2015-04-01    5560350.0
result >>> Name: pce, dtype: float64
```

### Math operations over the DataFrames/Series

Using aplly over column data

```python
titanic = pd.read_csv('data/titanic.csv', keep_default_na=True)
titanic.head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>survived</th>
      <th>pclass</th>
      <th>sex</th>
      <th>age</th>
      <th>sibsp</th>
      <th>parch</th>
      <th>fare</th>
      <th>embarked</th>
      <th>class</th>
      <th>who</th>
      <th>adult_male</th>
      <th>deck</th>
      <th>embark_town</th>
      <th>alive</th>
      <th>alone</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>0</td>
      <td>3</td>
      <td>male</td>
      <td>22.0</td>
      <td>1</td>
      <td>0</td>
      <td>7.2500</td>
      <td>S</td>
      <td>Third</td>
      <td>man</td>
      <td>True</td>
      <td>NaN</td>
      <td>Southampton</td>
      <td>no</td>
      <td>False</td>
    </tr>
    <tr>
      <th>1</th>
      <td>1</td>
      <td>1</td>
      <td>female</td>
      <td>38.0</td>
      <td>1</td>
      <td>0</td>
      <td>71.2833</td>
      <td>C</td>
      <td>First</td>
      <td>woman</td>
      <td>False</td>
      <td>C</td>
      <td>Cherbourg</td>
      <td>yes</td>
      <td>False</td>
    </tr>
    <tr>
      <th>2</th>
      <td>1</td>
      <td>3</td>
      <td>female</td>
      <td>26.0</td>
      <td>0</td>
      <td>0</td>
      <td>7.9250</td>
      <td>S</td>
      <td>Third</td>
      <td>woman</td>
      <td>False</td>
      <td>NaN</td>
      <td>Southampton</td>
      <td>yes</td>
      <td>True</td>
    </tr>
  </tbody>
</table>
</div>

```python
def my_mean(vector):
    return vector * 5

titanic.rename(columns={'fare':'1912_£'}, inplace=True)
titanic['1912_$'] = titanic['1912_£'].apply(my_mean)

titanic.loc[:, ['class', '1912_£', '1912_$']].head(5)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>class</th>
      <th>1912_£</th>
      <th>1912_$</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>Third</td>
      <td>7.2500</td>
      <td>36.2500</td>
    </tr>
    <tr>
      <th>1</th>
      <td>First</td>
      <td>71.2833</td>
      <td>356.4165</td>
    </tr>
    <tr>
      <th>2</th>
      <td>Third</td>
      <td>7.9250</td>
      <td>39.6250</td>
    </tr>
    <tr>
      <th>3</th>
      <td>First</td>
      <td>53.1000</td>
      <td>265.5000</td>
    </tr>
    <tr>
      <th>4</th>
      <td>Third</td>
      <td>8.0500</td>
      <td>40.2500</td>
    </tr>
  </tbody>
</table>
</div>

### Grouped Calculations

```python
df = pd.read_csv('data/economics.csv', parse_dates=['date'])
ic(df['date'].dtypes)

stdout >>> ic> df['date'].dtypes: dtype('<M8[ns]')
result >>> dtype('<M8[ns]')
```

```python
df.groupby('date')['pce', 'unemploy'].mean().head(10)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>2945</td>
    </tr>
    <tr>
      <th>1967-09-01</th>
      <td>516.3</td>
      <td>2958</td>
    </tr>
    <tr>
      <th>1967-10-01</th>
      <td>512.9</td>
      <td>3143</td>
    </tr>
    <tr>
      <th>1967-11-01</th>
      <td>518.1</td>
      <td>3066</td>
    </tr>
    <tr>
      <th>1967-12-01</th>
      <td>525.8</td>
      <td>3018</td>
    </tr>
    <tr>
      <th>1968-01-01</th>
      <td>531.5</td>
      <td>2878</td>
    </tr>
    <tr>
      <th>1968-02-01</th>
      <td>534.2</td>
      <td>3001</td>
    </tr>
    <tr>
      <th>1968-03-01</th>
      <td>544.9</td>
      <td>2877</td>
    </tr>
    <tr>
      <th>1968-04-01</th>
      <td>544.6</td>
      <td>2709</td>
    </tr>
  </tbody>
</table>
</div>

```python
df.groupby(by=df['date'].dt.strftime('%Y'))['unemploy', 'uempmed'].mean().reset_index().head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>date</th>
      <th>unemploy</th>
      <th>uempmed</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>1967</td>
      <td>3012.333333</td>
      <td>4.700000</td>
    </tr>
    <tr>
      <th>1</th>
      <td>1968</td>
      <td>2797.416667</td>
      <td>4.500000</td>
    </tr>
    <tr>
      <th>2</th>
      <td>1969</td>
      <td>2830.166667</td>
      <td>4.441667</td>
    </tr>
  </tbody>
</table>
</div>

```python
df.groupby(by=df['date'].dt.strftime('%Y'))['unemploy', 'uempmed'].first().reset_index().head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>date</th>
      <th>unemploy</th>
      <th>uempmed</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>1967</td>
      <td>2944</td>
      <td>4.5</td>
    </tr>
    <tr>
      <th>1</th>
      <td>1968</td>
      <td>2878</td>
      <td>5.1</td>
    </tr>
    <tr>
      <th>2</th>
      <td>1969</td>
      <td>2718</td>
      <td>4.4</td>
    </tr>
  </tbody>
</table>
</div>

```python
df.groupby(by=df['date'].dt.strftime('%Y'))['unemploy', 'uempmed'].last().reset_index().head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>date</th>
      <th>unemploy</th>
      <th>uempmed</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>1967</td>
      <td>3018</td>
      <td>4.8</td>
    </tr>
    <tr>
      <th>1</th>
      <td>1968</td>
      <td>2685</td>
      <td>4.4</td>
    </tr>
    <tr>
      <th>2</th>
      <td>1969</td>
      <td>2884</td>
      <td>4.6</td>
    </tr>
  </tbody>
</table>
</div>

### Joins

```python
size = 6
df_square = pd.DataFrame(rnd.randint(0, 10, size**2).reshape(size, -1), 
                  index=[f"R{x:02d}" for x in range(size)],
                  columns = [f"C{x:02d}" for x in range(size)])

df_square
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>C00</th>
      <th>C01</th>
      <th>C02</th>
      <th>C03</th>
      <th>C04</th>
      <th>C05</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>R00</th>
      <td>2</td>
      <td>6</td>
      <td>5</td>
      <td>7</td>
      <td>8</td>
      <td>4</td>
    </tr>
    <tr>
      <th>R01</th>
      <td>0</td>
      <td>2</td>
      <td>9</td>
      <td>7</td>
      <td>5</td>
      <td>7</td>
    </tr>
    <tr>
      <th>R02</th>
      <td>8</td>
      <td>3</td>
      <td>0</td>
      <td>0</td>
      <td>9</td>
      <td>3</td>
    </tr>
    <tr>
      <th>R03</th>
      <td>6</td>
      <td>1</td>
      <td>2</td>
      <td>0</td>
      <td>4</td>
      <td>0</td>
    </tr>
    <tr>
      <th>R04</th>
      <td>7</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>1</td>
      <td>5</td>
    </tr>
    <tr>
      <th>R05</th>
      <td>6</td>
      <td>4</td>
      <td>0</td>
      <td>0</td>
      <td>2</td>
      <td>1</td>
    </tr>
  </tbody>
</table>
</div>

```python
df_line = pd.Series(rnd.randint(0, 10, size), index=[f"R{x+3:02d}" for x in range(size)])
df_line.name="C06"
df_line

result >>> R03    4
result >>> R04    9
result >>> R05    5
result >>> R06    6
result >>> R07    3
result >>> R08    6
result >>> Name: C06, dtype: int64
```

```python
df_square['C06'] = df_line
df_square
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>C00</th>
      <th>C01</th>
      <th>C02</th>
      <th>C03</th>
      <th>C04</th>
      <th>C05</th>
      <th>C06</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>R00</th>
      <td>2</td>
      <td>6</td>
      <td>5</td>
      <td>7</td>
      <td>8</td>
      <td>4</td>
      <td>NaN</td>
    </tr>
    <tr>
      <th>R01</th>
      <td>0</td>
      <td>2</td>
      <td>9</td>
      <td>7</td>
      <td>5</td>
      <td>7</td>
      <td>NaN</td>
    </tr>
    <tr>
      <th>R02</th>
      <td>8</td>
      <td>3</td>
      <td>0</td>
      <td>0</td>
      <td>9</td>
      <td>3</td>
      <td>NaN</td>
    </tr>
    <tr>
      <th>R03</th>
      <td>6</td>
      <td>1</td>
      <td>2</td>
      <td>0</td>
      <td>4</td>
      <td>0</td>
      <td>4.0</td>
    </tr>
    <tr>
      <th>R04</th>
      <td>7</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>1</td>
      <td>5</td>
      <td>9.0</td>
    </tr>
    <tr>
      <th>R05</th>
      <td>6</td>
      <td>4</td>
      <td>0</td>
      <td>0</td>
      <td>2</td>
      <td>1</td>
      <td>5.0</td>
    </tr>
  </tbody>
</table>
</div>

clean up data

```python
df_square.drop('C06', axis=1, errors="ignore", inplace=True)
```

```python
df_square.join(pd.DataFrame(df_line), how="inner")
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>C00</th>
      <th>C01</th>
      <th>C02</th>
      <th>C03</th>
      <th>C04</th>
      <th>C05</th>
      <th>C06</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>R03</th>
      <td>6</td>
      <td>1</td>
      <td>2</td>
      <td>0</td>
      <td>4</td>
      <td>0</td>
      <td>4</td>
    </tr>
    <tr>
      <th>R04</th>
      <td>7</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>1</td>
      <td>5</td>
      <td>9</td>
    </tr>
    <tr>
      <th>R05</th>
      <td>6</td>
      <td>4</td>
      <td>0</td>
      <td>0</td>
      <td>2</td>
      <td>1</td>
      <td>5</td>
    </tr>
  </tbody>
</table>
</div>

```python
df_square.join(pd.DataFrame(df_line), how="outer")
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>C00</th>
      <th>C01</th>
      <th>C02</th>
      <th>C03</th>
      <th>C04</th>
      <th>C05</th>
      <th>C06</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>R00</th>
      <td>2.0</td>
      <td>6.0</td>
      <td>5.0</td>
      <td>7.0</td>
      <td>8.0</td>
      <td>4.0</td>
      <td>NaN</td>
    </tr>
    <tr>
      <th>R01</th>
      <td>0.0</td>
      <td>2.0</td>
      <td>9.0</td>
      <td>7.0</td>
      <td>5.0</td>
      <td>7.0</td>
      <td>NaN</td>
    </tr>
    <tr>
      <th>R02</th>
      <td>8.0</td>
      <td>3.0</td>
      <td>0.0</td>
      <td>0.0</td>
      <td>9.0</td>
      <td>3.0</td>
      <td>NaN</td>
    </tr>
    <tr>
      <th>R03</th>
      <td>6.0</td>
      <td>1.0</td>
      <td>2.0</td>
      <td>0.0</td>
      <td>4.0</td>
      <td>0.0</td>
      <td>4.0</td>
    </tr>
    <tr>
      <th>R04</th>
      <td>7.0</td>
      <td>0.0</td>
      <td>0.0</td>
      <td>1.0</td>
      <td>1.0</td>
      <td>5.0</td>
      <td>9.0</td>
    </tr>
    <tr>
      <th>R05</th>
      <td>6.0</td>
      <td>4.0</td>
      <td>0.0</td>
      <td>0.0</td>
      <td>2.0</td>
      <td>1.0</td>
      <td>5.0</td>
    </tr>
    <tr>
      <th>R06</th>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>6.0</td>
    </tr>
    <tr>
      <th>R07</th>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>3.0</td>
    </tr>
    <tr>
      <th>R08</th>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>6.0</td>
    </tr>
  </tbody>
</table>
</div>

```python
df_square.join(pd.DataFrame(df_line), how="left")
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>C00</th>
      <th>C01</th>
      <th>C02</th>
      <th>C03</th>
      <th>C04</th>
      <th>C05</th>
      <th>C06</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>R00</th>
      <td>2</td>
      <td>6</td>
      <td>5</td>
      <td>7</td>
      <td>8</td>
      <td>4</td>
      <td>NaN</td>
    </tr>
    <tr>
      <th>R01</th>
      <td>0</td>
      <td>2</td>
      <td>9</td>
      <td>7</td>
      <td>5</td>
      <td>7</td>
      <td>NaN</td>
    </tr>
    <tr>
      <th>R02</th>
      <td>8</td>
      <td>3</td>
      <td>0</td>
      <td>0</td>
      <td>9</td>
      <td>3</td>
      <td>NaN</td>
    </tr>
    <tr>
      <th>R03</th>
      <td>6</td>
      <td>1</td>
      <td>2</td>
      <td>0</td>
      <td>4</td>
      <td>0</td>
      <td>4.0</td>
    </tr>
    <tr>
      <th>R04</th>
      <td>7</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>1</td>
      <td>5</td>
      <td>9.0</td>
    </tr>
    <tr>
      <th>R05</th>
      <td>6</td>
      <td>4</td>
      <td>0</td>
      <td>0</td>
      <td>2</td>
      <td>1</td>
      <td>5.0</td>
    </tr>
  </tbody>
</table>
</div>

```python
df_square.join(pd.DataFrame(df_line), how="right")
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>C00</th>
      <th>C01</th>
      <th>C02</th>
      <th>C03</th>
      <th>C04</th>
      <th>C05</th>
      <th>C06</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>R03</th>
      <td>6.0</td>
      <td>1.0</td>
      <td>2.0</td>
      <td>0.0</td>
      <td>4.0</td>
      <td>0.0</td>
      <td>4</td>
    </tr>
    <tr>
      <th>R04</th>
      <td>7.0</td>
      <td>0.0</td>
      <td>0.0</td>
      <td>1.0</td>
      <td>1.0</td>
      <td>5.0</td>
      <td>9</td>
    </tr>
    <tr>
      <th>R05</th>
      <td>6.0</td>
      <td>4.0</td>
      <td>0.0</td>
      <td>0.0</td>
      <td>2.0</td>
      <td>1.0</td>
      <td>5</td>
    </tr>
    <tr>
      <th>R06</th>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>6</td>
    </tr>
    <tr>
      <th>R07</th>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>3</td>
    </tr>
    <tr>
      <th>R08</th>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>NaN</td>
      <td>6</td>
    </tr>
  </tbody>
</table>
</div>

### Sorting values and indexes

```python
economics = pd.read_csv('data/economics.csv', index_col='date',parse_dates=True)
economics.sort_values(by='unemploy', ascending=False).head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>2009-10-01</th>
      <td>9924.6</td>
      <td>308189</td>
      <td>5.4</td>
      <td>18.9</td>
      <td>15352</td>
    </tr>
    <tr>
      <th>2010-04-01</th>
      <td>10106.9</td>
      <td>309191</td>
      <td>5.6</td>
      <td>22.1</td>
      <td>15325</td>
    </tr>
    <tr>
      <th>2009-11-01</th>
      <td>9946.1</td>
      <td>308418</td>
      <td>5.7</td>
      <td>19.8</td>
      <td>15219</td>
    </tr>
  </tbody>
</table>
</div>

```python
economics.sort_index(ascending=False).head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>2015-04-01</th>
      <td>12158.9</td>
      <td>320887</td>
      <td>5.6</td>
      <td>11.7</td>
      <td>8549</td>
    </tr>
    <tr>
      <th>2015-03-01</th>
      <td>12161.5</td>
      <td>320707</td>
      <td>5.2</td>
      <td>12.2</td>
      <td>8575</td>
    </tr>
    <tr>
      <th>2015-02-01</th>
      <td>12095.9</td>
      <td>320534</td>
      <td>5.7</td>
      <td>13.1</td>
      <td>8705</td>
    </tr>
  </tbody>
</table>
</div>

## Time Series Analisis

```python
economics = pd.read_csv('data/economics.csv', index_col='date',parse_dates=True)
economics.head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
  </tbody>
</table>
</div>

```python
economics.index

result >>> DatetimeIndex(['1967-07-01', '1967-08-01', '1967-09-01', '1967-10-01',
result >>>                '1967-11-01', '1967-12-01', '1968-01-01', '1968-02-01',
result >>>                '1968-03-01', '1968-04-01',
result >>>                ...
result >>>                '2014-07-01', '2014-08-01', '2014-09-01', '2014-10-01',
result >>>                '2014-11-01', '2014-12-01', '2015-01-01', '2015-02-01',
result >>>                '2015-03-01', '2015-04-01'],
result >>>               dtype='datetime64[ns]', name='date', length=574, freq=None)
```

**TIME SERIES OFFSET ALIASES**


ALIAS  | DESCRIPTION
-------|-------------------------------------------------------
B      | business day frequency
C      | custom business day frequency (experimental)
D      | calendar day frequency
W      | weekly frequency
M      | month end frequency
SM     | semi-month end frequency (15th and end of month)
BM     | business month end frequency
CBM    | custom business month end frequency
MS     | month start frequency
SMS    | semi-month start frequency (1st and 15th)
BMS    | business month start frequency
CBMS   | custom business month start frequency
Q      | quarter end frequency
BQ     | business quarter endfrequency
QS     | quarter start frequency
BQS    | business quarter start frequency
A      | year end frequency
BA     | business year end frequency
AS     | year start frequency
BAS    | business year start frequency
BH     | business hour frequency
H      | hourly frequency
T, min | minutely frequency
S      | secondly frequency
L, ms  | milliseconds
U, us  | microseconds
N      | nanoseconds

### Time Resampling

```python
# yearly
economics.resample(rule='A').mean().head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-12-31</th>
      <td>515.166667</td>
      <td>199200.333333</td>
      <td>12.300000</td>
      <td>4.700000</td>
      <td>3012.333333</td>
    </tr>
    <tr>
      <th>1968-12-31</th>
      <td>557.458333</td>
      <td>200663.750000</td>
      <td>11.216667</td>
      <td>4.500000</td>
      <td>2797.416667</td>
    </tr>
    <tr>
      <th>1969-12-31</th>
      <td>604.483333</td>
      <td>202648.666667</td>
      <td>10.741667</td>
      <td>4.441667</td>
      <td>2830.166667</td>
    </tr>
  </tbody>
</table>
</div>

```python
economics.resample(rule='A').first().head(3)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-12-31</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1968-12-31</th>
      <td>531.5</td>
      <td>199808</td>
      <td>11.7</td>
      <td>5.1</td>
      <td>2878</td>
    </tr>
    <tr>
      <th>1969-12-31</th>
      <td>584.2</td>
      <td>201760</td>
      <td>10.0</td>
      <td>4.4</td>
      <td>2718</td>
    </tr>
  </tbody>
</table>
</div>

```python
# custom resampling function

def total(entry):
    if len(entry):
        return entry.sum()

economics.unemploy.resample(rule='A').apply(total).head(2)

result >>> date
result >>> 1967-12-31    18074
result >>> 1968-12-31    33569
result >>> Freq: A-DEC, Name: unemploy, dtype: int64
```

### Time/Data Shifting

http://pandas.pydata.org/pandas-docs/stable/generated/pandas.DataFrame.shift.html

```python
economics.unemploy.shift(1).head(2)

result >>> date
result >>> 1967-07-01       NaN
result >>> 1967-08-01    2944.0
result >>> Name: unemploy, dtype: float64
```

```python
# Shift everything forward one month
economics.unemploy.shift(periods=2, freq='M').head()

result >>> date
result >>> 1967-08-31    2944
result >>> 1967-09-30    2945
result >>> 1967-10-31    2958
result >>> 1967-11-30    3143
result >>> 1967-12-31    3066
result >>> Freq: M, Name: unemploy, dtype: int64
```

### Rolling and Expanding

A common process with time series is to create data based off of a rolling mean. The idea is to divide the data into "windows" of time, and then calculate an aggregate function for each window. In this way we obtain a simple moving average. 

```python
economics.unemploy.plot(figsize=(12,6)).autoscale(axis='x',tight=True)
economics.unemploy.rolling(window=14).mean().plot()
# economics.unemploy.rolling(window=30).max().plot()
# economics.unemploy.rolling(window=30).min().plot()
```

<img src='data:image/png;base64,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' />

Instead of calculating values for a rolling window of dates, what if you wanted to take into account everything from the start of the time series up to each point in time? For example, instead of considering the average over the last 7 days, we would consider all prior data in our expanding set of averages.

```python
economics.unemploy.expanding(min_periods=30).mean().plot(figsize=(12,6))
```

<img src='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAs8AAAF3CAYAAABaA6YyAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAALEgAACxIB0t1+/AAAADl0RVh0U29mdHdhcmUAbWF0cGxvdGxpYiB2ZXJzaW9uIDMuMC4zLCBodHRwOi8vbWF0cGxvdGxpYi5vcmcvnQurowAAIABJREFUeJzs3XeYXVW9//H3mp6ZyaRO6qRXEtLI0ESpShUBBURBQbliv17Lzwv32jteu17xIogo0hVBQboICAQSEiCV9N4nmSTTy/r9MTs4YGImyczsKe/X85zn7LPO3nu+e3E455N11tk7xBiRJEmSdGAZaRcgSZIkdRaGZ0mSJKmFDM+SJElSCxmeJUmSpBYyPEuSJEktZHiWJEmSWsjwLEmSJLVQi8JzCOHTIYQFIYT5IYTbQgh5IYRRIYRZIYRlIYQ7Qgg5ybq5yeNlyfMjm+3nmqR9SQjhjLY5JEmSJKltHDA8hxCGAv8OlMYYjwQygUuAa4EfxhjHAjuAK5NNrgR2JO0/TNYjhDAp2W4ycCbw8xBCZusejiRJktR2WjptIwvoEULIAvKBjcCpwN3J8zcD5yfL5yWPSZ4/LYQQkvbbY4w1McaVwDLgmMM/BEmSJKl9ZB1ohRjj+hDC94A1QBXwMDAH2BljrE9WWwcMTZaHAmuTbetDCOVAv6T9uWa7br7NPvXv3z+OHDmyxQcjSZIkHYo5c+ZsizEWH2i9A4bnEEIfmkaNRwE7gbtomnbRJkIIVwFXAQwfPpzZs2e31Z+SJEmSAAghrG7Jei2ZtvFWYGWMcWuMsQ74A3AC0DuZxgFQAqxPltcDw5IisoBewPbm7fvY5jUxxutjjKUxxtLi4gOGf0mSJKndtCQ8rwGOCyHkJ3OXTwMWAn8FLkzWuRy4N1m+L3lM8vzjMcaYtF+SnI1jFDAOeL51DkOSJElqey2Z8zwrhHA38CJQD8wFrgfuB24PIXwjabsx2eRG4LchhGVAGU1n2CDGuCCEcCdNwbse+HiMsaGVj0eSJElqM6FpULhjKi0tjc55liRJUlsLIcyJMZYeaD2vMChJkiS1kOFZkiRJaiHDsyRJktRChmdJkiSphQzPkiRJUgsZniVJkqQWMjxLkiRJLXTAi6RIkiRJXU19QyNlFbVs21PLzsraFm9neJYkSVKXU1XbwLIte1i2dTcbdlazYWcVG8ur2VhezeZd1eyorOVQrhVoeJYkSVKntru6jrlrdjJn9Q4WbtzFq5t3s6as8nXhuE9+NoN79WBo7zxmDO9NcWEu/Qtz6F+YS5+CHI6/tmV/y/AsSZKkTmXLrmqeWb6dOat3MHv1DpZs2kVjhIwAo/oXcOSQXlwwYygTBvZk3MBCSvrkk5ed2Sp/2/AsSZKkDm/Zlj08vHATDy/YzLy1OwEoyMlkxvA+fPLUcZSO7MP0Yb3pmZfdpnUYniVJktQhrdtRyZ9e2si989azeNNuAKaW9OJzp4/n5AkDmDioJ1mZ7XvyOMOzJEmSOoza+kYeWrCJ381azXMrygCYMbw3Xz53EmdMHsSQ3j1Src/wLEmSpNStLavktufXcOfstWzbU0tJnx589m3jOW/6UIb3y0+7vNcYniVJkpSKhsbI44u38LtZq/nbq1sJwGlHDOTSY4dz4rhiMjJC2iX+E8OzJEmS2tX2PTXcOmsNtz2/hg3l1QzomcsnTx3HJUcPS31axoEYniVJktQuFm/axU1Pr+KeeeuprW/kLeP686VzJ3PaEQPIbucf/h0qw7MkSZLa1Nw1O/jp48t4fPEW8rIzuHBmCR88YSRjB/RMu7SDZniWJElSm5i9qowfP7aUp5Zuo3d+Np9923guO24EfQpy0i7tkBmeJUmS1KqeW7Gdnzy2lGeWb6dfQQ5XnzWRy44bQWFu54+enf8IJEmS1CHMWrGd7z/8Ks+vKqO4Zy5fOOcI3nvscPJzuk7k7DpHIkmSpFQs2riL7z64mL8u2crAoly+cu4kLjlmOHnZmWmX1uoMz5IkSTok2/bU8K0HFnHP3PX0zM3i6rMmcsWbRnbJ0LyX4VmSJEkHpbExcsfstXz7gUVU1TVw1Ymj+dhJY+mVn512aW3O8CxJkqQWW7JpN/99zyvMXr2DY0f15ZsXTGHsgMK0y2o3hmdJkiQdUFVtAz95fCm/fHIFPfOy+J8Lp3LhzBJC6HiX0G5LhmdJkiT9S39dsoUv3TuftWVVXDizhP86+wj6duJzNR8Ow7MkSZL2acuuar7654Xc//JGRhcXcNuHjuP4Mf3SLitVhmdJkiS9TkNj5NZZq/nug0uoaWjkM28bz4dPGk1uVtc9i0ZLGZ4lSZL0mg07q/jMnfN4bkUZJ4ztxzfOn8Ko/gVpl9VhGJ4lSZIEwAOvbOSaP7xCXUMj333XVC4q7X4/CDwQw7MkSVI3t6emnq/et4C75qxj2rDe/Pjd0xnpaPM+GZ4lSZK6sXlrd/Kp2+eypqyST5wylk+9dRzZmRlpl9VhGZ4lSZK6oYbGyHVPLOOHjy5lUFEet3/oOI4d3b3PpNEShmdJkqRuZv3OKj59xzyeX1nGudOG8I3zj6RXj65/ae3WYHiWJEnqRv726lY+dftc6hsiP7h4GhfMGOqPAg+C4VmSJKkbaGyM/Oyvy/jho68yYWBPrrtspqegOwSGZ0mSpC5uZ2Utn75jHn9dspULZgzlWxdMoUeOFzw5FIZnSZKkLmz++nI+csscNu+q5uvnH8llxw53msZhMDxLkiR1UXfNXst//3E+/QpyuPPDxzNjeJ+0S+r0DM+SJEldTF1DI9+8fxG/fmYVbxrTj5++Zwb9CnPTLqtLMDxLkiR1IWUVtXz8dy/y7IrtXPnmUVxz1kSyvOhJqzE8S5IkdRELN+ziqt/OZsvuGr5/0TTeNbMk7ZK6HMOzJElSF3DvvPVc/ftX6NUjm7s+fDzThvVOu6QuyfAsSZLUidXWN/LN+xdy87OrOWZkX3526QwG9MxLu6wuy/AsSZLUSW0sr+Jjv3uRuWt28qG3jOLzZ04k2/nNbcrwLEmS1Ak9s2wbn7xtLtV1Dfz80qM4e8rgtEvqFgzPkiRJnUhDY+QXf1vO9x9ewpjiQn7xvpmMKS5Mu6xuw/AsSZLUSazcVsHn7nqJOat3cO60IXznnVMoyDXOtSd7W5IkqYOrrK3nl0+u5Lq/LSMnM4MfvXs6500f4mW2U3DAGeUhhAkhhHnNbrtCCP8RQugbQngkhLA0ue+TrB9CCD8JISwLIbwcQjiq2b4uT9ZfGkK4vC0PTJIkqbPbVV3Hb55dxanf+xs/fPRVTp04gEc+cxLnzxhqcE7JAUeeY4xLgOkAIYRMYD1wD3A18FiM8TshhKuTx/8JnAWMS27HAtcBx4YQ+gJfBkqBCMwJIdwXY9zR6kclSZLUSe2oqOWpZdt4YvEW/jJ/E1V1DcwY3pufvncGR4/sm3Z53d7BTts4DVgeY1wdQjgPODlpvxl4gqbwfB7wmxhjBJ4LIfQOIQxO1n0kxlgGEEJ4BDgTuO1wD0KSJKkziTGyq6qebRU1bN9Ty8pte1i0cTdz1+zg5fXlxAhFeVmcN30I7z12OFNLvOBJR3Gw4fkS/hF2B8YYNybLm4CByfJQYG2zbdYlbftrlyRJahUxRnZW1lFeVcfu6np2V9exq7qOXdX17K6uZ091PRW19eypaVqurW9s2o7YbB/J/ev2+4/lECArI5CREcjKCGSGfyxnZARibLpwSW1DI7X1DdQ1xKbH9Y3srqmnrKKGsopa6hqa/wXIz8lk0uAiPnXaOE4cX8y0kt5kZjg1o6NpcXgOIeQA7wCueeNzMcYYQoj/vNXBCyFcBVwFMHz48NbYpSRJ6gIaGiNbdlezsbyazeVN95t2VbOpvOm2cVcVm3fVvBaI9ycvO4PC3CwKc7PIycog0BRQ9zWFuPm84r1LjTHSGCP1jZHGxn/cN8RIQ2MEArlZGeRkZZCTmdwny0N75zF1aC/6FebQtyCH/oW59C3IYVjffEb0zSfDsNzhHczI81nAizHGzcnjzSGEwTHGjcm0jC1J+3pgWLPtSpK29fxjmsfe9ife+EdijNcD1wOUlpa2SiCXJKm+oZFd1fVU1zXQ0NgUfvbe7x0ZrGtoGi1saIzkZmWSk5VBbnIryM2iKC+bvOwMf6jVxnZV17FiawUrtu5h+dY9yXIFK7dX/FMwzsnKYHCvPAYW5XHU8D4MKspjQFEevXtk0zMvi6K993lN9wW5WV6BT4flYMLze3j9/OT7gMuB7yT39zZr/0QI4XaafjBYngTsh4Bv7T0rB3A6+xjFliRpf2rrGymvqnvttquq6Sv58qo6yivrXvfc3tvu6nrKq+rYU1PfKjXkZGZQ1CMJYz2y6V+QQ3HPXAb0zKW4KI8BPXMZWJTHoKI8invm+rX7AWzYWcULq8qYs3oHizftZsXWCrbtqXnt+cyMwIi++YwuLuCkCcWM6JfP4F55DCrqwaBeefTJz/YfM2pXLQrPIYQC4G3Ah5s1fwe4M4RwJbAauDhpfwA4G1gGVAIfAIgxloUQvg68kKz3tb0/HpQkdV/lVXWs3FbBym17WFtWxY7K2teCcdN9/WtBuKqu4V/uq0d2Jr16ZL92K+mTT68e2RT1yHqtrUd2JhnJPNXMZI5qTmYgOzOD7OQr9syM8Noc1Zr6BqrrGqmsbWiqp7rudfVtLK/mpXXlbK+oed28WGgKfsWFuQzslcfAnrkMSkZIBxblMaR3HiP6FTC4KK/bfFXf2BhZumUPL6wq44VVZcxetYP1O6sAKMjJ5IjBRZw6sZjRxYWM7l/A6OJChvfNJyfLkWJ1HCG+8f/0DqS0tDTOnj077TIkSa2gtr6ReWt3Mnt1WdNX8NsqWLWtgu0Vta9bryAnMwm82a8Lwq/d8pvui/L+eZ00Q1Z9QyNlFbVs3lXD5l3N5uLuqmbza7cayqvqXrddTmYGJX17MKJvPiP6FTC8bz4j+jXdSvrkk5edmdIRHb6yilpeWreTl9eWM2/tDuas3sGu6qZvAIp75nLMyL6UjuzD0SP7MnFQT7KcTqEUhRDmxBhLD7SeVxiUJLWJxsbI/A3l/H3Zdp5Zvo3Zq3a8NnI8sCiXkf0KOH3yQEb2K2BU/wJGFxcwrG8+uVmdMyxmZWYwIJlvO4Ve+12vqraBzbuqWbejitVlFawpq2TN9kpWb6/khVU7/ml6yaCiPIb3y0+mLhQyuriAMcWFjOiX32Hm7sYY2bK7hoUbdrFw4y4WbtjFy+t3srasaVQ5BBhTXMjZUwZTOrIvR4/sw/C++U63UKfkyLMkqdVs21PDE0u28uSrW3lq6VZ2VDaNso4fWMibxvTn+DH9OG5UP3rlZ6dcaccUY6SsopbVzQL16rKKpuWySrbu/sdc4KyMwPB++YwpLkxuBYwZUMiY/oVt2r/1DY2s3FbxWkjee9/8G4ThffOZPKSIacN6M62kN0cOLaJnnv/N1bE58ixJahcbdlbx4PxNPLRgEy+sKqMxQv/CXE6ZMIATxxdzwtj+FPfMTbvMTiGEQL/CXPoV5nLU8D7/9Pzu5CwUy5OzUCzf0rT8xJItrztncP/CHEY3C9Uj+zWN6g/pnUdhblaLRnxjjGzeVcPSLbtZtmUPr27ezcINu1i8aTc1yRkvcjIzGD+okNOOGMCkwUVMGtKLiYN7UmRQVhfmyLMk6aBV1zXw0IJN3Dl7Lc8s306MMGFgT844chCnTxrIpMFF3eZHcB1BfUMj63ZUsWzLHlZs+0eoXr51z2uj/3vlZGXQvyCHfsn5hXtkZ5KV/GCyMUbKq+ooq6hl5dYKdjebQtI7P7spIA8uYtKQptuY4sIOM3VEOlyOPEuSWlWMkZfXlXPn7LXc99IGdlfXU9KnB586bRzvmDaE0cWFaZfYbWVlZjCyfwEj+xfwjwv+NimrqG2aV11WycadVZRV1LJtTy3bk8tC7z3vdV1jI4FA7/xseufncMFRvRk3oJCxA3oydkAh/QtznKMsYXiWJB1AbX0jf3ppAzc8vZJFG3eRl53BWUcO5qLSEo4b1c8R5g6ub0HTleymD+uddilSl2B4liTtU3lVHbfOWsOvn1nJ5l01TBjYk29dMIW3TxvsnFZJ3ZbhWZL0Otv31PCrv6/kN8+sZndNPW8e25/vXjiNE8f192t7Sd2e4VmSBEBFTT03PLWSXz61goraes46chAfO3ksRw7d/zmLJam7MTxLUjdXU9/AHS+s5SePLWXbnlrOmDyQz50+gXEDe6ZdmiR1OIZnSeqmqusauHP2Wq57Yjkby6s5ZlRfrn//xH2eX1iS1MTwLEndTHllHbfMWs1Nf1/Jtj21lI7ow7XvmspbnNMsSQdkeJakbmLDzip+9fRKbnt+DRW1DZw4vpiPnDSa40f3MzRLUgsZniWpC6uua+CRhZu5Z+56/vbqVgDOnTqYq04cw6QhRSlXJ0mdj+FZkrqYxsbIrJVl3DN3HX95ZRO7a+oZVJTHv71lFJcdO4JhffPTLlGSOi3DsyR1Ecu27OYPL67n3nkbWL+zioKcTM6aMph3zhjKsaP7kemVACXpsBmeJakT21RezQOvbOSeuet5ZX05GQHeMq6Yz585gdMnDaJHTmbaJUpSl2J4lqROorExsn5nFQs27OLZ5dv4+/LtLNuyB4DJQ4r4wjlH8I7pQxjQMy/lSiWp6zI8S+r26hsa2bqnhrKKWnZU1FFWWcuOilr21NRTU9dATX0jNfWNB9xPRghkZkBuViYFuVkU5GbSIzuTEAIxRiI03UeSZYhEGpMHEahriFTW1FNZ10BlTT0VtQ1U1tazfmc1SzfvprK2AYAe2ZkcM6ovF5eWcMqEAV7QRJLaieFZUrdRXlnH/A3lLNywi1XbK1hTVsmaskrW76iivjHud7vcrAxysjLI+Benc9sbiusaG6mtb+Rf7K5FcjIzyM/NpCAni/ycTIp75nJx6TAmDOrJ+IE9mTK0FzlZGYf3RyRJB83wLKnLKq+qY9aK7TyzfDvPLN/Gq5v3vPZcn/xshvfNZ8rQXpwzZTBD+/SgX0EufQty6FuQTe/8HApzs8jNyjjocyDHGKmua2RPTT1VyUhxCHtvgbD3MYGMACTLIUB2RgY9cjINxpLUQRmeJXUpFTX1PLpoM/fN28CTS7dS1xDJy87g6JF9OW/6UKaW9GLykF70LchpsxpCCPTIyfTHepLUBRmeJXV6NfUN/G3JVu57aQOPLdpCVV0Dg4ryuOJNI3nrEQOZPrw3uVkGWUnS4TM8S+q0Fm/axc3PrOb+lzewq7qePvnZvPOoobxj2hCOHtmXDM9rLElqZYZnSZ1KY2Pk8cVb+NXfV/LM8u3kZWdw1pGDecf0Ibx5bH+yM50rLElqO4ZnSZ3Cnpp67pq9lpufWcWq7ZUM7pXHf545kfccM4ze+W03f1mSpOYMz5I6rBgjL6zawQOvbOT3c9axu6aeo4b35nNnTOCMyYMcZZYktTvDs6QOZ0dFLb9/cR23zlrDim0V5GRmcOaRg/jgm0cxfVjvtMuTJHVjhmdJHcayLbu5/skV/HHeBmrrGykd0YdPnDqW0ycPojDXtytJUvr8NJKUuhfX7OAXTyzn4YWbycvO4OLSEi47bgQTBxWlXZokSa9jeJaUihgjTyzZynV/W87zK8vonZ/Nv582jsuPH0G/wty0y5MkaZ8Mz5LaVV1DI39+eQP/97cVLN60myG98vji2ydxydHDKHBqhiSpg/OTSlK7qKlv4O4567juieWs21HF+IGFfP+iabxj+hDPmiFJ6jQMz5LaVHVdA3e8sJZf/G05G8urmT6sN185dzKnThzgFQAlSZ2O4VlSm6isrefWWWv4vydXsHV3DceM7Mt3L5zKm8f2JwRDsySpczI8S2pVdQ2N3P7CWn7y2FK27q7hTWP68dP3zOC40f3SLk2SpMNmeJbUKhobI/e/spHvP7yEVdsrOXpkH35+6VEcPbJv2qVJktRqDM+SDttTS7dy7YOLmb9+FxMG9uRXV5RyyoQBTs+QJHU5hmdJh2xjeRVf/ON8Hl20haG9e/CDi6dx3vShZPpDQElSF2V4lnTQGhsjtz6/hu/8ZTH1jY1cc9ZErjhhJLlZmWmXJklSmzI8Szoom8qr+dTtc5m1sowTxvbj2xdMZXi//LTLkiSpXRieJbXY00u38anb51Jd18B33zWVi0pLnNcsSepWDM+SWuTXf1/JV/+8kLHFhVx32UzGDihMuyRJktqd4VnSvxRj5IePLuUnjy3lbZMG8qN3T6cg17cOSVL35CegpH/pm/cv4oanV3LRzBK+/c4pZGVmpF2SJEmpMTxL2q9fPb2SG55eyeXHj+Ar75js/GZJUrfnEJKkfXp4wSa+fv9Czpg8kC+fa3CWJAkMz5L24eV1O/nU7fOYWtKbH717Bhle9ESSJMDwLOkNNpVXc+XNs+lXmMMN7y+lR44XPpEkaS/Ds6TXNDRGPn3HPCpq6rnpiqMp7pmbdkmSJHUo/mBQ0mt+8bflPLtiO9+9cCrjBvZMuxxJkjqcFo08hxB6hxDuDiEsDiEsCiEcH0LoG0J4JISwNLnvk6wbQgg/CSEsCyG8HEI4qtl+Lk/WXxpCuLytDkrSwZu7Zgc/eORV3j51MBfNLEm7HEmSOqSWTtv4MfBgjHEiMA1YBFwNPBZjHAc8ljwGOAsYl9yuAq4DCCH0Bb4MHAscA3x5b+CWlK6q2gY+c+dLDCrK45sXTPHMGpIk7ccBw3MIoRdwInAjQIyxNsa4EzgPuDlZ7Wbg/GT5POA3sclzQO8QwmDgDOCRGGNZjHEH8AhwZqsejaRD8t2HFrNyWwX/c9FUevXITrscSZI6rJaMPI8CtgI3hRDmhhBuCCEUAANjjBuTdTYBA5PlocDaZtuvS9r21y4pRc8u385Nf1/FFW8ayZvG9E+7HEmSOrSWhOcs4CjguhjjDKCCf0zRACDGGIHYGgWFEK4KIcwOIczeunVra+xS0n7sqann/939EiP75fP5MyekXY4kSR1eS8LzOmBdjHFW8vhumsL05mQ6Bsn9luT59cCwZtuXJG37a3+dGOP1McbSGGNpcXHxwRyLpIP0rQcWsX5nFd+7aBr5OZ58R5KkAzlgeI4xbgLWhhD2DkudBiwE7gP2njHjcuDeZPk+4P3JWTeOA8qT6R0PAaeHEPokPxQ8PWmTlILnVmzn1llr+NBbRlM6sm/a5UiS1Cm0dKjpk8DvQgg5wArgAzQF7ztDCFcCq4GLk3UfAM4GlgGVybrEGMtCCF8HXkjW+1qMsaxVjkLSQamtb+QLf5xPSZ8efPqt49MuR5KkTqNF4TnGOA8o3cdTp+1j3Qh8fD/7+RXwq4MpUFLr++VTK1i2ZQ83XXG0l9+WJOkgeHluqZtZW1bJTx9fypmTB3HKxAFplyNJUqdieJa6ma/+aQEZIfClcyelXYokSZ2O4VnqRh5ZuJlHF23hP946jiG9e6RdjiRJnY7hWeomKmvr+cp9Cxg/sJAPnDAq7XIkSeqUPLGr1E389PFlrN9ZxZ0fPp7sTP/dLEnSofATVOoGXlq7k+ufXMGFM0s4ZpTndJYk6VAZnqUurrqugc/e9RLFhbl88e3+SFCSpMPhtA2pi/v+w0tYtmUPv/ngMfTqkZ12OZIkdWqOPEtd2Auryrjh6ZVceuxwThxfnHY5kiR1eoZnqYuqrK3nc3e9REmfHvzX2UekXY4kSV2C0zakLupL9y5gTVklt3/oOApy/V9dkqTW4Miz1AXdM3cdd89ZxydPGcuxo/ulXY4kSV2G4VnqYlZuq+AL98znmJF9+ffTxqVdjiRJXYrhWepCqmob+Ogtc8jOyuBHl0wny4uhSJLUqpwIKXURMUb+655XWLJ5NzddcTRDevdIuyRJkroch6WkLuKW51Zzz9z1fPqt4zl5woC0y5EkqUsyPEtdwNNLt/HVPy3k1IkD+MQpY9MuR5KkLsvwLHVySzbt5qO3zGFMcSE/umQ6GRkh7ZIkSeqyDM9SJ7Z5VzUf/PUL5OVk8qsPHE1RnpffliSpLRmepU6qrKKWy26Yxc7KWn51+dEM9QeCkiS1Oc+2IXVCu6vruOKm51ldVsnNHziGKSW90i5JkqRuwZFnqZPZuruGy26YxcINu7ju0qM4foxXEJQkqb048ix1Isu27OGKm55n254arrtsJqcdMTDtkiRJ6lYMz1In8ZdXNvL5379MblYGd1x1PNOG9U67JEmSuh3Ds9TB1dQ38M37F/GbZ1czbVhvfvaeGQzrm592WZIkdUuGZ6kD27yrmo/cMoe5a3byb28exefPnEhOlj9VkCQpLYZnqYOas7qMj9zyIhU19Vx36VGcNWVw2iVJktTtGZ6lDujWWWv48n3zGdK7B7dceSwTBvVMuyRJkoThWepQausb+cqfFnDrrDWcNL6Yn1wyg175XjVQkqSOwvAsdRCbd1Xzsd+9yJzVO/joyWP43OkTyMwIaZclSZKaMTxLHcDsVWV89HdN85v/971Hcc5U5zdLktQRGZ6llN3+/Bq+8Mf5lPRxfrMkSR2d4VlKSWNj5H8eXsJ1TyznxPHF/PQ9M+jVw/nNkiR1ZIZnKQW19Y189q6X+NNLG3jvscP52jsmk5Xp+ZslSeroDM9SO6uoqecjt8zhqaXb+PyZE/joSWMIwR8GSpLUGRiepXa0s7KWK256gZfX7eS7F07l4tJhaZckSZIOguFZaidbdlfzvhueZ+X2Cq67bCZnTB6UdkmSJOkgGZ6ldrBuRyWX3TCLLbtruOmKozlhbP+0S5IkSYfA8Cy1seVb9/C+G2axp6ae3155LDNH9Em7JEmSdIgMz1Ibqa1v5I4X1vDtvyymR3Ymt111HJOH9Eq7LEmSdBgMz1Irq61v5PcvruNnjy9j/c4q3jSmH9+7aBpDevdIuzRJknSYDM9SK1lbVskf567n9hfWsn5nFdOH9eZb75zCieP6eyo6SZK6CMOzdBgqaur5y/xN3D1nLc+tKAPg+NH9DM2SJHVRhmfpIO2qruPvS7fx6KIt/GX+RiprGxjZL5/Pvm08Fxw1lJI++WmXKEmS2ojhWWqBVzfv5tFFm3liyVZeXL2D+sZIz7wfnrNnAAAd7klEQVQszp06hItKS5g5oo+jzJIkdQOGZ2k/tuyu5r55G7hn7noWbNgFwKTBRXzoxNGcMmEAM4b3JjszI+UqJUlSezI8S81U1Tbw8MJN/OHF9Ty9bBsNjZGpJb348rmTOHvKYAYW5aVdoiRJSpHhWd1eY2PkuRXb+cPc9Tw4fxN7auoZ2rsHHzlpNBfMGMrYAT3TLlGSJHUQhmd1W1t313D782u47fk1bCivpjA3i7OnDOKCGSUcO6ovGRnOYZYkSa9neFa3EmPkxTU7+c2zq3jglY3UNUTeMq4/15x9BG+bNJC87My0S5QkSR2Y4VndxjPLtvHdh5Ywb+1OeuZmcemxI3jf8SMYU1yYdmmSJKmTMDyry1uwoZxrH1zCk69uZUivPL5+3mTeeVQJBbm+/CVJ0sFpUXoIIawCdgMNQH2MsTSE0Be4AxgJrAIujjHuCE0nu/0xcDZQCVwRY3wx2c/lwBeS3X4jxnhz6x2K9Hoby6u49i+L+eO8DfTqkc1/n30E7zt+hFMzJEnSITuYobdTYozbmj2+GngsxvidEMLVyeP/BM4CxiW3Y4HrgGOTsP1loBSIwJwQwn0xxh2tcBzS6zyycDP/7+6XqKpt4KMnj+EjJ42hV4/stMuSJEmd3OF8b30ecHKyfDPwBE3h+TzgNzHGCDwXQugdQhicrPtIjLEMIITwCHAmcNth1CC9Tk19A99+YDG/fmYVRw4t4qfvOYpR/QvSLkuSJHURLQ3PEXg4hBCB/4sxXg8MjDFuTJ7fBAxMlocCa5ttuy5p21+71Co2llfxod/MZv76XXzghJFcfdZEcrOcoiFJklpPS8Pzm2OM60MIA4BHQgiLmz8ZY4xJsD5sIYSrgKsAhg8f3hq7VDcwf305H/z1C1TWNnD9+2Zy+uRBaZckSZK6oIyWrBRjXJ/cbwHuAY4BNifTMUjutySrrweGNdu8JGnbX/sb/9b1McbSGGNpcXHxwR2NuqVHFm7mol88S3ZmBnd/9HiDsyRJajMHDM8hhIIQQs+9y8DpwHzgPuDyZLXLgXuT5fuA94cmxwHlyfSOh4DTQwh9Qgh9kv081KpHo24lxsgNT63gqt/OZvzAQu75+JuYOKgo7bIkSVIX1pJpGwOBe5rOQEcWcGuM8cEQwgvAnSGEK4HVwMXJ+g/QdJq6ZTSdqu4DADHGshDC14EXkvW+tvfHg9LBijHytT8v5Ka/r+LMyYP44bun0yPH+c2SJKlthaaTYnRMpaWlcfbs2WmXoQ4mxshX7lvAzc+u5gMnjOSL50wiIyOkXZYkSerEQghzYoylB1rPS6ypU4kx8tU/LeTmZ1fzb28exX+fcwTJtyKSJEltrkU/GJQ6ihueWsmvn1nFB08wOEuSpPZneFan8fjizXzrL4s4e8ogvmBwliRJKTA8q1N4dfNu/v22eUweUsT3L5ruHGdJkpQKw7M6vLKKWq68+QXyczL55ftLPauGJElKjT8YVIdWW9/IR26Zw5ZdNdzx4eMZ3KtH2iVJkqRuzPCsDivGyBf/OJ/nV5bx40umM31Y77RLkiRJ3ZzTNtRh3fj0Su6YvZZPnjqW86YPTbscSZIkw7M6pr8u3sK3HljEWUcO4tNvHZ92OZIkSYDhWR3Q8q17+Pfb5nLE4CK+f/E0z6whSZI6DMOzOpSKmno+8ts5ZGdlcP37S8nPcVq+JEnqOAzP6jBijHz+9y+zfOsefvqeGQzt7Zk1JElSx2J4Vodx49Mruf/ljfy/MyZywtj+aZcjSZL0TwzP6hCeW7Gdb/9lMWdMHshHThqddjmSJEn7ZHhW6jaVV/OJW19kRL98vnfRNELwB4KSJKljMjwrVbX1jXzsd3OorG3g/y6bSc+87LRLkiRJ2i9PZaBUffP+hby4Zic/e+8Mxg3smXY5kiRJ/5Ijz0rNPXPXcfOzq/m3N4/i7VOHpF2OJEnSARmelYplW/ZwzR9e4dhRfbn6rIlplyNJktQihme1u9r6Rj59xzx6ZGfy0/fMICvTl6EkSeocnPOsdveTx5byyvpyfnHZUQwoyku7HEmSpBZzyE/t6sU1O/j5E8u4cGYJZx45OO1yJEmSDorhWe2mvqGRL9wzn4FFeXz53ElplyNJknTQDM9qN7+btYaFG3fxpbdP8nzOkiSpUzI8q11s3V3D9x5ewlvG9efMIwelXY4kSdIhMTyrXVz74GKq6xr46jsme/ltSZLUaRme1eYWbCjn9y+u44MnjGJ0cWHa5UiSJB0yw7Pa3Hf+sphePbL52Clj0y5FkiTpsBie1aaefHUrTy3dxidOGUuvHv5IUJIkdW6GZ7WZxsbIt/+ymJI+PXjf8SPSLkeSJOmwGZ7VZv44bz2LNu7i/50xgdyszLTLkSRJOmyGZ7WJ6roGvvfQEqYM7cW5U4ekXY4kSVKrMDyrTdw1ey0byqu5+qyJZGR4ajpJktQ1GJ7V6uobGrn+qRXMGN6bN43pl3Y5kiRJrcbwrFb3wPxNrC2r4iMnjfGCKJIkqUsxPKtVxRi57onljCku4G1HDEy7HEmSpFZleFarenLpNhZt3MWHTxzjXGdJktTlGJ7Vqn7xxHIGFuVy3gzPsCFJkroew7NazUtrd/Lsiu1c+eZRntdZkiR1SYZntZpf/G05PfOyeM8xw9MuRZIkqU0YntUqVmzdw4MLNvG+40bQMy877XIkSZLahOFZreKXT60gOzODD5wwKu1SJEmS2ozhWYdt/c4qfj9nPRfNLKG4Z27a5UiSJLUZw7MO2//+dRmRyMdOGZt2KZIkSW3K8KzDsm5HJXfNXsvFpcMY2rtH2uVIkiS1KcOzDsv//nUZgcDHHXWWJEndgOFZh2xtWSV3zV7Hu48exhBHnSVJUjdgeNYh+9njy8jICHzslDFplyJJktQuDM86JGu2V3L3i+t47zHDGdzLUWdJktQ9GJ510GKMfOHe+eRkZvDRkx11liRJ3YfhWQftlllrePLVrVxz9kQGFuWlXY4kSVK7aXF4DiFkhhDmhhD+nDweFUKYFUJYFkK4I4SQk7TnJo+XJc+PbLaPa5L2JSGEM1r7YNT2XlhVxtf+tICTxhdz2bEj0i5HkiSpXR3MyPOngEXNHl8L/DDGOBbYAVyZtF8J7Ejaf5isRwhhEnAJMBk4E/h5CCHz8MpXe/rrki1c8avnKemTz08umUFGRki7JEmSpHbVovAcQigBzgFuSB4H4FTg7mSVm4Hzk+Xzksckz5+WrH8ecHuMsSbGuBJYBhzTGgehthVj5JbnVvNvN89mRL8CbvvQcfTKz067LEmSpHaX1cL1fgR8HuiZPO4H7Iwx1ieP1wFDk+WhwFqAGGN9CKE8WX8o8FyzfTbfRh3Uhp1VfPGP83ls8RZOmVDMT997FIW5LX3ZSJIkdS0HTEEhhLcDW2KMc0IIJ7d1QSGEq4CrAIYPH97Wf077sbOyll8/s4pfPrmCxgj/ffYRfOCEkWRl+htTSZLUfbVkCPEE4B0hhLOBPKAI+DHQO4SQlYw+lwDrk/XXA8OAdSGELKAXsL1Z+17Nt3lNjPF64HqA0tLSeCgHpUO3dXcNNzy9glueXU1FbQNvmzSQL719EsP65qddmiRJUuoOGJ5jjNcA1wAkI8+fizFeGkK4C7gQuB24HLg32eS+5PGzyfOPxxhjCOE+4NYQwg+AIcA44PnWPRwdqg07q7j+yRXc9vwaahsaOWfKYD5+yliOGFyUdmmSJEkdxuFMXv1P4PYQwjeAucCNSfuNwG9DCMuAMprOsEGMcUEI4U5gIVAPfDzG2HAYf1+HKcbIgg27uOW51fz+xXXECBfMGMpHTx7D6OLCtMuTJEnqcEKMHXdmRGlpaZw9e3baZXQp2/fU8NTSbSzYUM7ji7ewfGsFOVkZvLt0GB8+aTQlfZyeIUmSup8QwpwYY+mB1vO0Cd3Aiq17uP/ljTy+ZAvz1u4kRsjJymDGsN5c+ebRnHXkIPoU5KRdpiRJUodneO6idlXXcf/LG7l7zjrmrN4BwLSSXvzHaeM5deIAJg0pItOLnEiSJB0Uw3MXEmPkhVU7uHXWah5csInqukbGDijkmrMmcv6MoQwsyku7REmSpE7N8NwFxBi5/5WN/OSxpby6eQ8987K4aOYwLpxZwtSSXjRd4FGSJEmHy/DcyS3ZtJsv3TufWSvLmDioJ9e+awrnThtCfo7/aSVJklqbCasTe3D+Jv7jjrn0yM7kWxdM4d1HD3MesyRJUhsyPHdStzy3mi/eO59pJb355ftLKe6Zm3ZJkiRJXZ7huZOJMfKzx5fx/Ude5dSJA/jf9x5Fj5zMtMuSJEnqFgzPnUiMkf95aAk/f2I575wxlGsvnEp2ZkbaZUmSJHUbhudO5EePLuXnTyznPccM55vnH0mG85slSZLalcOWncTPHl/Kjx9bykUzSwzOkiRJKTE8dwLXP7mc7z38Ku+cMZTvvGuqwVmSJCklhucO7v6XN/KtBxZzztTB/M9F0zwVnSRJUooMzx3YS2t38pk751E6og8/uNjgLEmSlDbDcwe1qbyaf/vNbAYU5fJ/75tJbpano5MkSUqb4bkDqq5r4MO3zKGypp4bLz+afoVeAEWSJKkj8FR1HUyMkS/dO5+X1u7kF5fNZPzAnmmXJEmSpIQjzx3M7S+s5c7Z6/jkqWM588hBaZcjSZKkZgzPHcjaskq+8eeFnDC2H59+6/i0y5EkSdIbGJ47iMbGyOfvfpkQAtd6LmdJkqQOyfDcQdwyazXPrtjOf59zBCV98tMuR5IkSftgeO4AVm2r4NsPLObE8cVccvSwtMuRJEnSfhieU7Z3ukZWZuDad00hBKdrSJIkdVSG55TdPWcdz68q44vnTGJwrx5plyNJkqR/wfCcop2VtXznwcWUjujDhTNL0i5HkiRJB2B4TtF3H1pCeVUdXz//SM+uIUmS1AkYnlPy0tqd3Pb8Gt5//AiOGFyUdjmSJElqAcNzChoaI1+8dz79C3P59Nu8GIokSVJnYXhOwe0vrOHldeV84ZwjKMrLTrscSZIktZDhuZ2VV9bxvYeWcMyovrxj2pC0y5EkSdJBMDy3sx8++irlVXV85dzJntNZkiSpkzE8t6Olm3fz2+dW855jhjNpiD8SlCRJ6mwMz+3o6/cvoiAnk8+ePiHtUiRJknQIDM/t5Jll23jy1a38+2nj6FuQk3Y5kiRJOgSG53YQY+Tah5YwuFcelx03Iu1yJEmSdIgMz+3g4YWbeWntTj512jjysjPTLkeSJEmHyPDcxhoaI997aAmj+xdw4cyStMuRJEnSYTA8t7F7561n6ZY9fOb08WRl2t2SJEmdmWmuDdXWN/LDR19l8pAizj5ycNrlSJIk6TAZntvQ719cx9qyKj53xgQyMrwgiiRJUmdneG4jjY2RG55aweQhRZw8vjjtciRJktQKDM9t5IlXt7B8awUfestoL8MtSZLURRie28gvn1zJ4F55nDPVuc6SJEldheG5DcxfX86zK7ZzxZtGku0ZNiRJkroMk10buPHplRTkZHLJMcPTLkWSJEmtyPDcyjaWV/Gnlzbw7qOH06tHdtrlSJIkqRUZnlvZr59ZRWOMfOCEkWmXIkmSpFZmeG5FZRW1/O65NZw1ZTDD+uanXY4kSZJameG5Ff38r8uorK3n028dl3YpkiRJagOG51ayYWcVv3luNe88qoSxA3qmXY4kSZLagOG5lfzksaUQ4T8cdZYkSeqyDhieQwh5IYTnQwgvhRAWhBC+mrSPCiHMCiEsCyHcEULISdpzk8fLkudHNtvXNUn7khDCGW11UO1t6ebd3DVnHZceN5ySPs51liRJ6qpaMvJcA5waY5wGTAfODCEcB1wL/DDGOBbYAVyZrH8lsCNp/2GyHiGEScAlwGTgTODnIYTM1jyYNMQY+cIf51OYm8UnThmbdjmSJElqQwcMz7HJnuRhdnKLwKnA3Un7zcD5yfJ5yWOS508LIYSk/fYYY02McSWwDDimVY4iRffMXc+slWVcfdZE+hXmpl2OJEmS2lCL5jyHEDJDCPOALcAjwHJgZ4yxPlllHTA0WR4KrAVIni8H+jVv38c2nVJ5ZR3femAR04f15t2lw9IuR5IkSW2sReE5xtgQY5wOlNA0WjyxrQoKIVwVQpgdQpi9devWtvozreJ7Dy+hrKKWb5x/JBkZIe1yJEmS1MYO6mwbMcadwF+B44HeIYSs5KkSYH2yvB4YBpA83wvY3rx9H9s0/xvXxxhLY4ylxcXFB1Neu3pp7U5umbWa9x8/kiOH9kq7HEmSJLWDlpxtoziE0DtZ7gG8DVhEU4i+MFntcuDeZPm+5DHJ84/HGGPSfklyNo5RwDjg+dY6kPZUXdfAf/7+ZfoX5vKZ08enXY4kSZLaSdaBV2EwcHNyZowM4M4Y459DCAuB20MI3wDmAjcm698I/DaEsAwoo+kMG8QYF4QQ7gQWAvXAx2OMDa17OO3jfx5awuJNu/nVFaUU5WWnXY4kSZLayQHDc4zxZWDGPtpXsI+zZcQYq4GL9rOvbwLfPPgyO44nX93KjU+v5P3Hj+DUiQPTLkeSJEntyCsMHoRN5dV89q6XGDegkP86+4i0y5EkSVI7a8m0DQG7quu48uYXqKpt4JYrjyUvu9Nf30WSJEkHyZHnFthUXs2lv5zFkk27+el7ZzBhUM+0S5IkSVIKHHk+gN8+u4pvPbCYxhi5/v0zOWXCgLRLkiRJUkoMzwdQ1CObk8YX819nH8HwfvlplyNJkqQUGZ4P4LzpQzlveqe+irgkSZJaiXOeJUmSpBYyPEuSJEktZHiWJEmSWsjwLEmSJLWQ4VmSJElqIcOzJEmS1EKGZ0mSJKmFDM+SJElSCxmeJUmSpBYyPEuSJEktZHiWJEmSWsjwLEmSJLWQ4VmSJElqoRBjTLuG/QohbAVW7+fpXkD5Iex2OLDmELY71L/XWbZr7345nG27et90lu3sl32zX/bNftk3+2Xf7Jd9s1/2rbX6ZUSMsfiAW8UYO+UNuP4Qt9vazn+vs2zXrv1i33SJ7ewX+8V+sV/sF/ulI2zXrv3Smadt/OkQt9vZzn+vs2zX3v1yONt29b7pLNvZL/tmv+yb/bJv9su+2S/7Zr/sW7v2S4eettEWQgizY4yladfR0dgv+2ff7Jv9sm/2y77ZL/tmv+yb/bJv9su+tXe/dOaR50N1fdoFdFD2y/7ZN/tmv+yb/bJv9su+2S/7Zr/sm/2yb+3aL91u5FmSJEk6VN1x5FmSJEk6JF0iPIcQfhVC2BJCmN+sbVoI4dkQwishhD+FEIqaPTc1eW5B8nxe0v7uEMLLSfu1aRxLazqYfgkhXBpCmNfs1hhCmP6G/d3XfF+dVWv1Szd/vWSHEG5O2heFEK55w74yQwhzQwh/bu/jaG2t1S8hhE+FEOYnr5f/SONYWtNB9ktOCOGmpP2lEMLJ+9hfd3x/2W+/dMH3l2EhhL+GEBYmx/SppL1vCOGREMLS5L5P0h5CCD8JISxL+uGoN+yvKISwLoTwszSOp7W0Zr+EEK5N3mPmhxDendYxtZZD6JuJyf9nNSGEz+1jf633uXQop+joaDfgROAoYH6ztheAk5LlDwJfT5azgJeBacnjfkBmcr8GKE7abwZOS/vY2qtf3rDdFGD5G9reCdzafF+d9dYa/dLdXy/Ae4Hbk+V8YBUwstl2n0leL39O+7g6Qr8ARwLzk7Ys4FFgbNrH1o798nHgpmR5ADAHyGi2Xbd8f9lfv3TR95fBwFHJck/gVWAS8F3g6qT9auDaZPls4C9AAI4DZr1hfz9OXjM/S/vYOkK/AOcAjyTvLwXJa64o7eNr574ZABwNfBP43D7212qfS11i5DnG+CRQ9obm8cCTyfIjwLuS5dOBl2OMLyXbbo8xNgCjgaUxxq3Jeo8226ZTOsh+ae49wO17H4QQCml60X2jDcpsd63UL9399RKBghBCFtADqAV2AYQQSmh6I7+hrWtuD63UL0fQ9CFXGWOsB/5GU2DstA6yXyYBjyfbbaHptFKl0O3fX/bXL13x/WVjjPHFZHk3sAgYCpxH0z8OSO7PT5bPA34TmzwH9A4hDAYIIcwEBgIPt+MhtIlW7JdJwJMxxvoYYwVNg4RntuOhtLqD7ZsY45YY4wtA3Rv31dqfS10iPO/HApo6GOAiYFiyPB6IIYSHQggvhhA+n7QvAyaEEEYmH3znN9umK9lfvzT3buC2Zo+/DnwfqGzb0lJ1sP3S3V8vdwMVwEaaRsi+F2PcGxh+BHweaGzHOtvbwfbLfOAtIYR+IYR8mkaPutPr5SXgHSGErBDCKGBms+e68/vL/vqlS7+/hBBGAjOAWcDAGOPG5KlNNIViaApJa5tttg4YGkLIoOn18k9fy3d2h9MvNL2Wzgwh5IcQ+gOn0P1eM/9Kq34udeXw/EHgYyGEOTQN99cm7VnAm4FLk/sLQginxRh3AB8F7gCeounr1ob2Lrod7K9fAAghHAtUxhjnJ4+nA2NijPe0e6Xt66D6xdcLx9B0vEOAUcBnQwijQwhvB7bEGOekUm37Oah+iTEuAq6laaTsQWAe3ev18iuaPuRn0/Qh9gzQ4PvLvvulK7+/JN80/B74jxjjrubPxabv1g90CrCPAQ/EGNe1UYmpONx+iTE+DDxA02voNuBZfM3s3b7VP5eyWmtHHU2McTFNUzQIIYynabgemt6onowxbkuee4Cm+WmPxRj/RHK1mRDCVXSRF15z/6Jf9rqE1486Hw+UhhBW0fR6GRBCeCLGeHLbV9t+DqFf6Oavl/cCD8YY64AtIYS/0/R18wyaRtLOBvKAohDCLTHGy9q/+rZzCP2yIsZ4I3Bjss23aHov6lL21y/JVJVP710vhPAMTfMXT6Ibv7/8i37pku8vIYRsmkLQ72KMf0iaN4cQBscYNybTD7Yk7et5/chpSdJ2PE3f4nwMKARyQgh7YoxXt89RtL5W6hdijN+kab4vIYRbSV5LndlB9s3+nEArfy512ZHnEMKA5D4D+ALwi+Sph4ApyVcbWTS9eS98wzZ9aPrXbZeYs9ncv+iXvW0X02y+c4zxuhjjkBjjSJpG6l/tah9scPD98oZtuuPrZQ1wavJcAU0/XFkcY7wmxliSvF4uAR7vasEZDr5f3rDNcP7xA7kuZX/9krzfFiTLbwPqY4wLu/v7y/765Q3bdIn3lxBCoOkfj4tijD9o9tR9wOXJ8uXAvc3a3x+aHAeUJ3NgL40xDk9eM5+jaf5vZw7OrdIvoelMEv2SfU4FptLJ54QfQt/sU1t8LnWJkecQwm3AyUD/EMI64MtAYQjh48kqfwBugqav20MIP6Dpl6iRpq9/7k/W+3EIYVqy/LUYY6f+V9vB9EviRGBtjHFFuxbazlqxX7rz6+V/gZtCCAto+tX3TTHGl9u55HbRiv3y++TDrQ74+P9v735ebArjOI6/P2pIGdmQvZ2UKYsppWSpZEpWTJHIP2BDykKxl7VZ+lVKU5qFDVlIfkyyspAsRKlBzXK+Fucspok6U9epc+f9Wp7nubfveRa3T899zvlW1VJf9/A/rHNddgELSVZodslmey63NyNcl7H6faHZ/ZsF3id51167DNwE7ic5B3ym2aiA5gjCUZrz38vA2X7L7c2o1mUCeN7kTX4Bp9t/NoZsXWuTZDfNEajtwEqaV4LuXXvUYxTsMChJkiR1NLbHNiRJkqRRMzxLkiRJHRmeJUmSpI4Mz5IkSVJHhmdJkiSpI8OzJA1QkmtJ/tmiOMlMkr191iRJG4HhWZLG0wxgeJakEfM9z5I0EEmu0HTU+g58AV4DP4ELwGaaxgmzwBQw3479BE60X3Eb2EnTXOF82zZakrQOhmdJGoAkB4A5YJqmO+wbmjbPd6rqRzvnOvCtqm4lmQPmq+phO/YUuFhVH5NMAzeq6kj/dyJJwzYW7bklaQM4BDyqqmWAJI/b6/va0LwD2AYsrP1gkm3AQeBB274XYMt/r1iSxpDhWZKGbQ6YqarFJGeAw3+ZswlYqqqpHuuSpLHkA4OSNAzPgJkkW5NMAsfa65PA1yQTwKlV83+3Y1TVL+BTkpMAaezvr3RJGh+GZ0kagKp6A9wDFoEnwKt26CrwEngBrH4A8C5wKcnbJHtogvW5JIvAB+B4X7VL0jjxgUFJkiSpI3eeJUmSpI4Mz5IkSVJHhmdJkiSpI8OzJEmS1JHhWZIkSerI8CxJkiR1ZHiWJEmSOjI8S5IkSR39ASUhiaWFL5RCAAAAAElFTkSuQmCC' />

<div class="alert alert-info">Read more at <a href="TimeSeriesAnalysis.ipynb">TimeSeriesAnalysis.ipynb</a>
</div>   

## Visualization

* https://pandas.pydata.org/pandas-docs/stable/reference/api/pandas.DataFrame.plot.html
* https://matplotlib.org/

<div class="alert alert-info">Read more at <a href="visualization.ipynb">visualization.ipynb</a>
</div>   

```python
import matplotlib.pyplot as plt
%matplotlib inline
```

```python
economics = pd.read_csv('data/economics.csv', index_col='date',parse_dates=True)
economics.head(2)
```

<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>pce</th>
      <th>pop</th>
      <th>psavert</th>
      <th>uempmed</th>
      <th>unemploy</th>
    </tr>
    <tr>
      <th>date</th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>1967-07-01</th>
      <td>507.4</td>
      <td>198712</td>
      <td>12.5</td>
      <td>4.5</td>
      <td>2944</td>
    </tr>
    <tr>
      <th>1967-08-01</th>
      <td>510.5</td>
      <td>198911</td>
      <td>12.5</td>
      <td>4.7</td>
      <td>2945</td>
    </tr>
  </tbody>
</table>
</div>

```python
# ploting all data for a range of x values (date index)
_ = economics.unemploy.plot(
    figsize=(12,4),  
    xlim=['1970', '1980'],
)
```

<img src='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAt8AAAELCAYAAAABTepgAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAALEgAACxIB0t1+/AAAADl0RVh0U29mdHdhcmUAbWF0cGxvdGxpYiB2ZXJzaW9uIDMuMC4zLCBodHRwOi8vbWF0cGxvdGxpYi5vcmcvnQurowAAIABJREFUeJzt3Xd8leX9//HXJ3vvBMhiJYAgghDBPbC1rhZHHdW2trXl26q1e2i//XZZO36ddmhtxdGqOKoVra2lLhRBDUtkB0IgIZBNdk6Sc/3+ODcQkZlxThLez8fjPHJy3eu6Lw4n73Od675uc84hIiIiIiIDLyzUFRAREREROV4ofIuIiIiIBInCt4iIiIhIkCh8i4iIiIgEicK3iIiIiEiQKHyLiIiIiATJEcO3mc03syoze/eA8i+a2QYzW2tmP+9RfpuZlZjZRjP7UI/yC72yEjP7do/ysWb2plf+mJlF9dfJiYiIiIgMJkfT8/0AcGHPAjM7D5gLTHPOTQF+4ZVPBq4Fpnjb/NHMws0sHPgDcBEwGfiYty7Az4BfO+cKgHrgxr6elIiIiIjIYHTE8O2cWwzUHVD8BeCnzrkOb50qr3wusMA51+GcKwVKgFneo8Q5t9U55wMWAHPNzIA5wJPe9g8Cl/XxnEREREREBqXejvmeAJzlDRd51cxO8cpzgB091iv3yg5Vng40OOe6DigXERERERl2IvqwXRpwKnAK8LiZjeu3Wh2Cmc0D5gHEx8fPnDRp0kAfUkRERESOY8uXL69xzmX21/56G77Lgaeccw54y8z8QAZQAeT1WC/XK+MQ5bVAiplFeL3fPdd/H+fcvcC9AEVFRa64uLiX1RcREREROTIzK+vP/fV22Mk/gPMAzGwCEAXUAAuBa80s2szGAoXAW8DbQKE3s0kUgYsyF3rh/WXgo95+bwCe6e3JiIiIiIgMZkfs+TazR4FzgQwzKwe+B8wH5nvTD/qAG7wgvdbMHgfWAV3Azc65bm8/twAvAOHAfOfcWu8Q3wIWmNkdwErgvn48PxERERGRQcMCmXno0bATERERERloZrbcOVfUX/vTHS5FRERERIJE4VtEREREJEgUvkVEREREgkThW0REREQkSBS+RURERESCROFbRERERCRIFL5FRERERIJE4VtEREREJEgUvkVEREREgkThW0REREQkSBS+RURERESCROFbRERERCRIFL5FRERERIJE4VtEREREJEgUvkVEREREgkThW0REREQkSBS+RURERESCROFbRERERCRIFL5FRERERIJE4VtEREREJEgUvkVEREREgkThW0REREQkSI4Yvs1svplVmdm7B1n2NTNzZpbh/W5mdpeZlZjZO2Y2o8e6N5jZZu9xQ4/ymWa2xtvmLjOz/jo5EREREZHB5Gh6vh8ALjyw0MzygAuA7T2KLwIKvcc84G5v3TTge8BsYBbwPTNL9ba5G/hcj+3edywRERERkeHgiOHbObcYqDvIol8D3wRcj7K5wEMuYBmQYmajgA8Bi5xzdc65emARcKG3LMk5t8w554CHgMv6dkoiIiIiIoNTr8Z8m9lcoMI5t/qARTnAjh6/l3tlhysvP0i5iIiIiMiwE3GsG5hZHHA7gSEnQWVm8wgMZyE/Pz/YhxcRERER6ZPe9HyPB8YCq81sG5ALrDCzkUAFkNdj3Vyv7HDluQcpPyjn3L3OuSLnXFFmZmYvqi4iIiIiEjrHHL6dc2ucc1nOuTHOuTEEhorMcM7tAhYCn/RmPTkV2OOcqwReAC4ws1TvQssLgBe8ZY1mdqo3y8kngWf66dxERERERAaVo5lq8FFgKTDRzMrN7MbDrP48sBUoAf4M3ATgnKsDfgS87T1+6JXhrfMXb5stwL96dyoiIiIiIoObBSYZGXqKiopccXFxqKshIiIiIsOYmS13zhX11/50h0sRERERkSBR+BYRERERCRKFbxERERGRIFH4FhEREREJEoVvEREREZEgUfgWEREREQkShW8RERERkSBR+BYRERERCRKFbxERERGRIFH4FhEREREJEoVvEREREZEgUfgWEREREQkShW8RERERkSBR+BYRERERCRKFbxERERGRIFH4FhEREREJEoVvEREREZEgUfgWEREREQkShW8RERERkSBR+BYRERERCRKFbxERERGRIFH4FhEREREJEoVvEREREZEgOWL4NrP5ZlZlZu/2KPt/ZrbBzN4xs6fNLKXHstvMrMTMNprZh3qUX+iVlZjZt3uUjzWzN73yx8wsqj9PUERERERksDianu8HgAsPKFsEnOicOwnYBNwGYGaTgWuBKd42fzSzcDMLB/4AXARMBj7mrQvwM+DXzrkCoB64sU9nJCIiIiIySB0xfDvnFgN1B5T9xznX5f26DMj1ns8FFjjnOpxzpUAJMMt7lDjntjrnfMACYK6ZGTAHeNLb/kHgsj6ek4iIiIjIoNQfY74/A/zLe54D7OixrNwrO1R5OtDQI8jvLRcRERERGXb6FL7N7DtAF/Bw/1TniMebZ2bFZlZcXV0djEOKiIiIiPSbXodvM/sUcClwvXPOecUVQF6P1XK9skOV1wIpZhZxQPlBOefudc4VOeeKMjMze1t1EREREZGQ6FX4NrMLgW8CH3HOtfZYtBC41syizWwsUAi8BbwNFHozm0QRuChzoRfaXwY+6m1/A/BM705FRERERGRwO5qpBh8FlgITzazczG4Efg8kAovMbJWZ3QPgnFsLPA6sA/4N3Oyc6/bGdN8CvACsBx731gX4FvBVMyshMAb8vn49QxERERGRQcL2jxgZWoqKilxxcXGoqyEiIiIiw5iZLXfOFfXX/nSHSxERERGRIFH4FhEREREJEoVvEREREZEgUfgWEREREQkShW8RERERkSBR+BYRERERCRKFbxERERGRIFH4FhEREREJEoVvEREREZEgUfgWEREREQkShW8RERERkSBR+BYRERERCRKFbxERERGRIFH4FhEREREJEoVvEREREZEgUfgWEREREQkShW8RERERkSBR+BYRERERCRKFbxERERGRIFH4FhEREREJEoVvEREREZEgUfgWEREREQmSI4ZvM5tvZlVm9m6PsjQzW2Rmm72fqV65mdldZlZiZu+Y2Ywe29zgrb/ZzG7oUT7TzNZ429xlZtbfJykiIiIiMhgcTc/3A8CFB5R9G3jROVcIvOj9DnARUOg95gF3QyCsA98DZgOzgO/tDezeOp/rsd2BxxIRERERGRaOGL6dc4uBugOK5wIPes8fBC7rUf6QC1gGpJjZKOBDwCLnXJ1zrh5YBFzoLUtyzi1zzjngoR77EhEREREZVno75nuEc67Se74LGOE9zwF29Fiv3Cs7XHn5QcpFRERERIadPl9w6fVYu36oyxGZ2TwzKzaz4urq6mAcUkRERESk3/Q2fO/2hozg/azyyiuAvB7r5XplhyvPPUj5QTnn7nXOFTnnijIzM3tZdRERERGR0Oht+F4I7J2x5AbgmR7ln/RmPTkV2OMNT3kBuMDMUr0LLS8AXvCWNZrZqd4sJ5/ssS8RERERkWEl4kgrmNmjwLlAhpmV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' />

```python
# or we can specify what column we want to plot
ax = economics['unemploy'].plot(
    figsize=(14,4), 
    xlim=['1970', '1980'],
    ylim=[2200, 9000],
    title="Unemployment for period 1970-1980"
)
_ = ax.set(ylabel="Unemployment", xlabel="Years")

# this will make data to fit, no matter what we set in x
# ax.autoscale(axis='x',tight=True)
```

<img src='data:image/png;base64,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iCSpmFUgZtbQzKaa2Wwzm29mdwfLu5jZl2a2zMxeMrMGwfKsYH5Z8HznGm3dHixfbGanxipmERERERGRcMSya2c3cJK79wX6AaeZ2RDg98BD7t4NKAauCta/CigOlj8UrIeZ9QZGAEcApwGPmVlGDOMWERERERGplZgVVh5SGszWDx4OnAS8GiwfB5wTTA8P5gmeP9nMLFg+3t13u/tKYBkwKFZxi4iIiIiI1FZML0YyswwzmwVsAiYAy4Gt7l4ZrLIO6BhMdwTWAgTPbwPa1Fy+l21qvta1ZjbNzKZt3rw5FrsjIiIiIiKyVzEtrNy9yt37ATmEepl6xfC1Rrt7vrvnZ2dnx+plREREREREviMuw+e5+1bgY+AYoKWZZQZP5QAFwXQB0AkgeL4FUFhz+V62ERERERERSbhYjgqYbWYtg+lGwCnAQkIF1vnBaqOAN4Ppt4J5guc/cncPlo8IRg3sAnQHpsYqbhERERERkdrKPPAqYTsEGBeM4FcPeNnd3zGzBcB4M7sXmAk8Faz/FPCsmS0DigiNBIi7zzezl4EFQCVwg7tXxTBuERERERGRWrFQp1B6yc/P92nTpiU6DBERERERSWJmNt3d86PRVlyusRIREREREUlnKqxEREREREQipMJKREREREQkQiqsREREREREIqTCSkREREREJEIqrERERERERCKkwkpERERERCRCKqxEREREREQipMJKREREREQkQiqsREREREREIqTCSkREREREJEIqrERERERERCKkwkpERERERCRCKqxEREREREQipMJKREREREQkQiqsREREREREIqTCSkREREREJEIqrERERERERCKkwkpERERERCRCKqxEREREREQiZO6e6Biizsx2AvOj2GQLYFsSthXt9nKBNVFqC5J7X6MdWzRzp7yFR3kLj/IWnmTOW7TbS+bPBuUt8W1Fuz3lLTzKW/iOcPdGUWnJ3dPuAWyOcnujk7GtGMSWtHlLgX+HqOVOeUua2JQ35S0l8xaDf4ek/WxQ3hLflvKWHO0pb8mRu3Q9FXBrlNt7O0nbinZ7yZy3aLcX7diimTvlLTzKW3iUt/Akc96i3V4yfzYob4lvK9rtKW/hUd7CF7XcpeupgNPcPT/RcaQa5S18yl14lLfwKG/hUd7Co7yFR3kLj/IWHuUtfNHMXbr2WI1OdAApSnkLn3IX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' />

```python
# x as slicing
economics['unemploy']['1970':'1979'].plot(figsize=(14,4), ls='--',c='r').autoscale(axis='x',tight=True)
```

<img src='data:image/png;base64,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' />

```python
# or providing argument
_ = economics['unemploy'].plot(figsize=(14,4), xlim=['1970','1980'], ylim=[2200, 9000])
```

<img src='data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAA0gAAAEKCAYAAAAo6HWVAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAALEgAACxIB0t1+/AAAADl0RVh0U29mdHdhcmUAbWF0cGxvdGxpYiB2ZXJzaW9uIDMuMC4zLCBodHRwOi8vbWF0cGxvdGxpYi5vcmcvnQurowAAIABJREFUeJzs3Xd8leX9//HXlT3IJhAISSCEvSFsRNyKA/cGxEHdtvZn1S5ttVWrbb9qRUUQceFArdSFE0F22HskZEIGWWTn5Jzr90eOFCojYMLJeD8fjzzOOfe5znV/bor0vHNd93UZay0iIiIiIiICXp4uQEREREREpLlQQBIREREREXFTQBIREREREXFTQBIREREREXFTQBIREREREXFTQBIREREREXFrUEAyxtxnjNlsjNlijPml+1ikMeYrY8wu92OE+7gxxjxnjNltjNlojBl6SD9T3e13GWOmNs0liYiIiIiInJzjBiRjTH/gNmAEMAi4yBiTBDwEfGOt7QF8434NcAHQw/0zHXjR3U8k8Agw0t3XIz+GKhERERERkeagISNIfYCV1tpKa20d8D1wOTAJmOtuMxe41P18EvC6rbcCCDfGdALOA76y1hZZa4uBr4DzG/FaREREREREfhafBrTZDPzFGBMFVAETgRSgo7V2n7tNLtDR/TwWyDrk89nuY0c7fhhjzHTqR54IDg4e1rt37wZfjIiIiIiItD1r1qzZb62Nboy+jhuQrLXbjDFPAV8CFcB6wPk/bawxxjZGQdbamcBMgOTkZJuSktIY3YqIiIiISCtljMlorL4atEiDtXa2tXaYtXY8UAzsBPLcU+dwP+a7m+cAcYd8vIv72NGOi4iIiIiINAsNXcWug/sxnvr7j94GFgA/rkQ3FfjY/XwBMMW9mt0ooNQ9FW8hcK4xJsK9OMO57mMiIiIiIiLNQkPuQQL4wH0PkgO4y1pbYox5EnjPGHMLkAFc7W77GfX3Ke0GKoFpANbaImPMY8Bqd7s/W2uLGuk6REREREREfjZjbaPcOtQkdA+SiIiIiIgcjzFmjbU2uTH6atAUOxERERERkbZAAUlERERERMRNAUlERERERMRNAUlERERERMRNAUlERERERMRNAUlERERERMRNAUlERERERMRNAUlERERERMRNAUlERERERMRNAUlERERE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' />

### X-Ticks

https://matplotlib.org/api/dates_api.html

### Date Formatting
Formatting follows the Python datetime [**strftime**](http://strftime.org/) codes. 

The following examples are based on `datetime.datetime(2001, 2, 3, 16, 5, 6)` :

CODE | MEANING                                              | EXAMPLE
-----| -----------------------------------------------------|--------------
`%Y` | Year with century as a decimal number                | 2001
`%y` | Year without century as a zero-padded decimal number | 01
`%m` | Month as a zero-padded decimal number                | 02
`%B` | Month as locale’s full name                          | February
`%b` | Month as locale’s abbreviated name                   | Feb
`%d` | Day of the month as a zero-padded decimal number     | 03
`%A` | Weekday as locale’s full name                        | Saturday
`%a` | Weekday as locale’s abbreviated name                 | Sat
`%H` | Hour (24-hour clock) as a zero-padded decimal number | 16
`%I` | Hour (12-hour clock) as a zero-padded decimal number | 04
`%p` | Locale’s equivalent of either AM or PM               | PM
`%M` | Minute as a zero-padded decimal number               | 05
`%S` | Second as a zero-padded decimal number               | 06

----

CODE  | MEANING                                | EXAMPLE
------| ---------------------------------------|--------------
`%#m` | Month as a decimal number. (Windows)   | 2 
`%-m` | Month as a decimal number. (Mac/Linux) | 2 
`%#x` | Long date                              | Saturday, February 03, 2001 
`%#c` | Long date and time                     | Saturday, February 03, 2001 16:05:06 
</table>  
    

```python
from datetime import datetime
datetime(2001, 2, 3, 16, 5, 6).strftime("%A, %B %d, %Y  %I:%M:%S %p")

result >>> 'Saturday, February 03, 2001  04:05:06 PM'
```

```python
from matplotlib import dates
```

```python
# TODO - find better example so i can use month or weekend formatter.
# using Date Locator's
```

## Miscellaneous

### Loops 

```python
# series 
s = pd.Series(range(10000))
# list
l = list(s)
```

Panda's Time Series

```python
%timeit -n 1 -r 1 [s.iloc[i] for i in range(len(s))]

stdout >>> 153 ms ± 0 ns per loop (mean ± std. dev. of 1 run, 1 loop each)
```

Python's list

```python
%timeit -n 1 -r 1 [l[i] for i in range(len(l))]

stdout >>> 1.05 ms ± 0 ns per loop (mean ± std. dev. of 1 run, 1 loop each)
```

Pandas - accesing lst element

```python
%timeit -n 1 -r 1 s.iloc[-1:]

stdout >>> 233 µs ± 0 ns per loop (mean ± std. dev. of 1 run, 1 loop each)
```

List - accesing lst element

```python
%timeit -n 1 -r 1 l[-1:]

stdout >>> 1.32 µs ± 0 ns per loop (mean ± std. dev. of 1 run, 1 loop each)
```