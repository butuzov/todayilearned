# PyTest

This notes are result of reading book "[Python Testing with pytest](https://www.amazon.com/Python-Testing-pytest-Effective-Scalable/dp/1680502409)" and implementing it on practice.

Additional Resources:
 - https://pythontesting.net/framework/pytest/pytest-introduction/
 - https://docs.pytest.org/en/latest/index.html
 - http://hackebrot.github.io/pytest-tricks/

## TODO - Examples

* [ ] flask examples
* [ ] getting setarted examples
* [ ] configuration (`conftest.py`)

## TODO - Getting Started

## Writing Test Functions

### Assert Everything
https://wiki.python.org/moin/UsingAssertionsEffectively

```python
def test_asserts_ok():
    """ Just Don't: Tuple is True """
    name = "ssss"
    assert type(name) is int, ("name is not a int: {}".format(name))
    
def test_asserts_bad():
    """ Just Don't: Tuple is True """
    name = "ssss"
    assert (type(name) is int, ("name is not a int: {}".format(name))

```

### Assert Exception

```python
def test_approx():
    with pytest.raises(AssertionError):
        """expected exception"""
        assert .1 + .2 == .3

    with pytest.raises(AssertionError):
        """unexpected exception"""
        assert 4 / 0
    """ using approx """
    assert .1 + .2 == pytest.approx(.3, rel=1e-19) 
```

## Setups/Teardowns

```python
def setup_module(module):
    print("\nAction[{}] / Func[{}] / Argument [{}]".format(
        "setup_module", setup_module.__name__, module.__name__))


def teardown_module(module):
    print("\nAction[{}] / Func[{}] / Argument [{}]".format(
        "teardown_module", teardown_module.__name__, module.__name__))


def setup_function(func):
    print("\nAction[{}] / Func[{}] / Argument [{}]".format(
        "setup_function", setup_function.__name__, func.__name__))


def teardown_function(func):
    print("\nAction[{}] / Func[{}] / Argument [{}]".format(
        "teardown_function", teardown_function.__name__, func.__name__))
    
class TestClass:

    @classmethod
    def setup_class(cls):
        print("Class Setup", cls)

    @classmethod
    def teardown_class(cls):
        print("Class teardown", cls)

    def setup_method(self, method):
        print("Method: Setup, {}".format(method))

    def teardown_method(self, method):
        print("Method: Teardown, {}".format(method))
```

### Scopes

We can user setups and teadowns within scope context (`module`, `function`, `session`, `class`)

```python
import pytest

@pytest.fixture(scope="function", autouse=True)
def setup(request):
    """
    request.fixturename
    request.scope
    https://docs.pytest.org/en/latest/reference.html#request
    """
    print("Setup", request.scope)

    def fabrica():
        pass

    yield fabrica
    print("TearDown", request.scope)


@pytest.mark.usefixtures("setup")
def test_scopes1():
    assert True


@pytest.mark.usefixtures("setup")
def test_scopes2():
    assert True
```

### Using `addfinalizer` to add tearDown steps

---

```python
import pytest

@pytest.fixture()
def finalizer_demo(request):
    """ creating some heavy presetup """

    def teardown1():
        """ destroys pre setup - step 1 """
        print("teardown1")

    request.addfinalizer(teardown1)

    def teardown2():
        """ destroys pre setup - step 2 """
        print("teardown2")

    request.addfinalizer(teardown2)


def test_finalizer(finalizer_demo):
    assert True
```

## Fixtures

```python
## todo
```

## Marks

### Skips

```python

@pytest.mark.skip(reason="Lets just skip it")
def test_skipper_1():
    """skipping it, not going to run"""
    assert False


def test_skipper_2():
    """skipping it, not going to run"""
    pytest.skip("skipping it, not going to run")


@pytest.mark.skipif(sys.version_info < (3, 6),
                    reason="requires python3.6 or higher")
def test_skipper_3():
    """skipping it, not going to run"""
    print(f"Running on  Python {sys.version_info[0]}.{sys.version_info[1]}")


py36 = pytest.mark.skipif(sys.version_info < (3, 6),
                          reason="requires python3.6 or higher")


@py36
def test_skipper_4():
    """skipping it, not going to run"""
    print(f"Running on  Python {sys.version_info[0]}.{sys.version_info[1]}")
    
```

### X-Fail and X-Pass

Useful keys 

* `-rx`  - extra sumarry for : (x)failed
* `--lf` - last failed only or all
* `--ff` - first failed 


```python
import pytest, sys

# Expect to fail, xpass otherwise
@pytest.mark.xfail
def test_failure1():
    assert False

# Fail within test function
@pytest.mark.xfail
def test_failure2():
    pytest.fail(msg="Failed", pytrace=True)

# Fail if X Pass
@pytest.mark.xfail(strict=True)
def test_failure3():
    assert False

# Expect Failure if True 
@pytest.mark.xfail(sys.version_info < (3, 6), reason="python3.6 api changes")
def test_failure4():
    answer = 42
    assert f"The answer is {answer}" == "The answer is 42"

# Expect Failure if true (expression as string)
@pytest.mark.xfail('1==1', reason="dummy reason")
def test_failure5():
    assert False

# fail, don't execute
@pytest.mark.xfail(run=False)
def test_failure6():
    pass

# Fail with expection expected - Expections tuple
@pytest.mark.xfail(raises=(ZeroDivisionError,))
def test_failure7():
    4 / 0

# Failt with expection expected - Just expection
@pytest.mark.xfail(raises=IndexError)
def test_failure8():
    x = []
    x[1] = 1
```

## Builtin Fixtures

### Temporary Directory

```python
def test_tmpdir(tmpdir):
    CONTENT = "test"
    p = tmpdir.mkdir("sub").join("hello.txt")
    p.write(CONTENT)

    assert p.read() == CONTENT
    assert len(tmpdir.listdir()) == 1

# examples of some methods
def pprint(name, what):
    base = "/private/var/folders"
    print("{:<10}: {}".format(name, str(what).replace(base, "")))


def test_directory(tmpdir):
    print()

    pprint("dir base", tmpdir)
    pprint("dir not", tmpdir / "n")

    base = tmpdir.mkdir("n")
    pprint("dir (n)", base)
    base = tmpdir.mkdir("f").mkdir("1")
    pprint("dir (n1)", base)
    pprint("dir (n1:s)", base.strpath)

    aspathfile = base.join("hello.txt")
    pprint("file", aspathfile)
    pprint("file_str", aspathfile.strpath)

    aspathfile.write("lol")

    base = tmpdir.mkdir("s").join('inventory')
    pprint("dir (ensure)", base.ensure())
    pprint("dir (join:s)", base.strpath)

```

### Parametrize

```python
import pytest
from collections import OrderedDict

xfail = pytest.mark.xfail

tests_table = OrderedDict({
    'ints': (1, 2, 3),
    'floats': (1., 2., 3.),
    'sts': ("twenty", "two", "twentytwo"),
    'floats_failed':
    pytest.param(.1, .2, .3, marks=xfail(reason="Float point addition"))
})


@pytest.mark.parametrize(("a", "b", "expected"), tests_table.values())
def test_sample(a, b, expected):
    assert a + b == expected


@pytest.mark.parametrize("a, b, expected", tests_table.values(), ids=list(tests_table.keys()))
def test_sample2(a, b, expected):
    assert a + b == expected
```

## Plugins

### Functionality

| Plugin             | Description  
|--------------------|---------------
| pytest-repeat      | Run Tests More Than Once
| pytest-xdist       | Run Tests in Parallel
| pytest-timeout     | Put Time Limits on Your Tests
| pytest-instafail   | See Details of Failures and Errors as They Happen
| pytest-pycodestyle | Comply with Python’s Style Guide
| pytest-pep8        | Comply with Python’s Style Guide
| pytest-flake8      | Check for Style + Linting
| pytest-pylint      | Check for Style + Linting
| pytest-selenium    | Test with a Web Browser
| pytest-django      | django
| pytest-flask       | flask

----

```bash
# pytest-repeat: repeat all tests twice
pytest --count=2
# pytest-xdist: automatically detect the number of CPUs 
pytest -n auto
# pytest-timeout: give test .5 seconds to pass
pytest --timeout=0.5
# pytest-instafail: 
pytest --instafail --maxfail=2 
```

### Representations


| Plugin             | Description  
|--------------------|---------------
| pytest-sugar       | Instafail + Colors + Progress Bar
| pytest-html        | Generate HTML Reports for Test Sessions

-----


```bash
# pytest-html: report.html 
pytest --html=report.html
```

###  Todo
- pytest-runner
- pytest-cov

## Configuration

You can configure `pytest` in next files `pytest.ini` or `tox.ini` or `setup.cfg`:

### `pytest.ini`

```ini
[pytest]
addopts = -ra --verbose -q
testpaths = tests
```

### `tox.ini`

We can use section `[pytest]` for `pytest` settings.

```ini
[pytest]
addopts = -ra --verbose --cov --cov-report=html
testpaths = tests
```

### `setup.cfg`

We can use `[tool:pytest]` section to describe settings for pytest.

```ini
[tool:pytest]
addopts = -ra --verbose --cov --cov-report=html
testpaths = tests
```

## Using pytest with other tools

### `venv` & `pip`

Regular virtual environment usage:

```bash
> python -m venv .venv
> source .venv/bin/activate
(.venv) > pip install pytest
(.venv) > pytest
(.venv) > deactivate
> 
```