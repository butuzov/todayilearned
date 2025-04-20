# `uv`

## Checks

- https://docs.astral.sh/uv

## Usage

```shell
# Init Package (or --lib)
uv init . --package

# Add Dev dependencies
uv add --dev ruff ipykernel pytest pytest-cov mypy

# Add project dependencies
uv add python-fasthtml pandas duckdb

# oops! remvoe it!
uv remove duckdb

# Lock
uv lock --upgrade-package <package>==<version>

# Syncing
## Install the versions recorded in uv.lock
uv sync

## Or update the packages and uv.lock
## Note: It’s your responsibility to ensure everything works before committing
uv sync --upgrade

# Running
uv run python script.py
uv run myapp.main
# tooling..
uv run pytest
uv run mypy

# will run jupyter in the current project... without adding jupyter and its dependencies to the project
uv run --with jupyter jupyter notebook

# Building
uv build
```

## Managing Python itself

```shell
# list of pythons available
uv python list

# Install this version
uv python install cpython-3.10

# Shell
uvx -p cpython-3.10 python

# running package with python specific version
uvx --with pendulum -p 3.13t python
```

## Starting App

```python
# myapp/__init__.py
import argparse

def start_app():
    parser = argparse.ArgumentParser(description="Start the app")
    parser.add_argument("--port", type=int, help="Port number")
    args = parser.parse_args()
    print(f"App started at port {args.port}")
```

```toml
# pyproject.toml
[project.scripts]
start = "myapp:start_app"
```

```shell
uv run start --help
uv run start --port 1234
# > App started at port 1234
```

## Jupyter

1. `uv tool run jupyter lab`
2. `!uv pip install my-dependcies`

## Scripts

Self Containing (dependencies)

```python
#!/usr/bin/env -S uv run

# /// script
# dependencies = [ 'requests', ]
# ///

import asyncio

import requests

async def main():
    await asyncio.sleep(1)
    print(args)

if __name__ == "__main__":
    import sys
    asyncio.run(main(sys.argv))

```
