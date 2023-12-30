# Type System

```python
from typing import Dict, List, Optional, Union, Optional
```

```python
def r(input:List[str]) -> List[str]:
    return list(reversed(input))

my_list = [1,2,3,4]
r(my_list), r("foobar"), r(1)
```

```python
c: int | float = 3.4
c = 1
c = "foo"
```