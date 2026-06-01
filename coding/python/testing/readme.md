<!-- menu: Testing -->

# Testing with `pytest`

```shell
pip install pytest
```

## Future Reading

- https://docs.pytest.org/en/stable/
- [Python Testing with pytest: Simple, Rapid, Effective, and Scalable](https://www.amazon.com/Python-Testing-pytest-Effective-Scalable/dp/1680502409)

## Testing With PyTest

```python
# content of test_sample.py
def inc(x):
    return x + 1


def test_answer():
    assert inc(3) == 5
```

## Continuous integration

### github

```yaml
#
name: Build & Test

on:
  push:
    tags:
      - v[0-9]+.[0-9]+.[0-9]+
    branches:
      - main
  pull_request:

jobs:
  Build:
    runs-on: ubuntu-24.04
    strategy:
      fail-fast: false
      matrix:
        python-version: ["3.10", "3.11", "3.12", "3.13", "3.14"]

    steps:
      - uses: actions/checkout@v5
      - name: Set up Python ${{ matrix.python-version }}
        uses: actions/setup-python@v6
        with:
          python-version: ${{ matrix.python-version }}
      - uses: astral-sh/setup-uv@v7
        with:
          enable-cache: true
          cache-dependency-glob: pyproject.toml
      - run: uv pip install '.[tests]' --system
```

## Fixtures

https://docs.pytest.org/en/7.1.x/how-to/fixtures.html

```python
@pytest.fixture
def fruit_bowl():
    return ["apple", "banana"]

def test_fruit_salad(fruit_bowl):
    fruit_salad = FruitSalad(*fruit_bowl)
    ...
    asert ...
```

## Plugins

### [`pytest-crayons`](https://github.com/okken/pytest-crayons)

```python
def test_magenta(magenta):
    magenta("this should be magenta")
```

### `pytest-httpserver`

```python
def test_json_client(httpserver: HTTPServer):
    httpserver.expect_request("/foobar").respond_with_json({"foo": "bar"})
    assert requests.get(httpserver.url_for("/foobar")).json() == {"foo": "bar"}
```

### `pytest-cov`
