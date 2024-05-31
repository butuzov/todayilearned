# `ruff`

- https://docs.astral.sh/ruff/

## Config

`.ruff.toml`, `pyproject.toml` or `ruff.toml` in the file's directory or any parent directory.

```toml
# Support Python 3.10+.
target-version = "py310"
# Set the maximum line length to 79.
line-length = 79

[lint]
# Add the `line-too-long` rule to the enforced rule set.
extend-select = ["E501"]

select = [
    "B",       # Bugbear
    "E",       # pycodestyle, standard
    "F",       # Enable all pyflakes rules
    "UP",      # Enable all pyupgrade rules by default
    "RUF100",  # Ban unused `# noqa` comments
    "PGH004",  # Ban blanket `# noqa` comments (only ignore specific error codes)
]

ignore = [
    # Unnecessary parentheses to functools.lru_cache: just leads to unnecessary churn.
    # https://github.com/python/cpython/pull/104684#discussion_r1199653347.
    "UP011",
    # Use format specifiers instead of %-style formatting.
    # Doesn't always make code more readable.
    "UP031",
    # Use f-strings instead of format specifiers.
    # Doesn't always make code more readable.
    "UP032",
    # Use PEP-604 unions rather than tuples for isinstance() checks.
    # Makes code slower and more verbose. https://github.com/astral-sh/ruff/issues/7871.
    "UP038",
    "E501", # line-too-long
    "F403", # undefined-local-with-import-star
    "F405", # undefined-local-with-import-star-usage
]

unfixable = [
    # The autofixes sometimes do the wrong things for these;
    # it's better to have to manually look at the code and see how it needs fixing
    "F841",  # Detects unused variables
    "F601",  # Detects dictionaries that have duplicate keys
    "F602",  # Also detects dictionaries that have duplicate keys
]

```

## Checks

- https://docs.astral.sh/ruff/rules/

## VSCode Integration

```json
{
  "ruff.lint.run": "onSave",
  "ruff.showNotifications": "onError",
  "ruff.lint.args": [
    "check", "--config", "/path/ruff.toml"
  ],

}
```
