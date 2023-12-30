# `grep`


## RegExps

- https://unix.stackexchange.com/q/181254/


```shell
# Grab urls
grep -Eo "(http|https)://[a-zA-Z0-9./?=_%:-]*" source.html | sort -u

# same but piped
grep -Eoi '<a [^>]+>' source.html |
  grep -Eo 'href="[^\"]+"' |
  grep -Eo '(http|https)://[^/"]+'

# PCRE
grep -Po '(?<=href=")[^"]*(?=")'

```
- `grep -P` lookaround expressions
- `grep -E`is the same as egrep
- `grep -o` only outputs what has been grepped
- `(http|https)` is an either / or
- `a-z` is all lower case
- `A-Z` is all upper case
- `.` is dot
- `/` is the slash
- `?` is ?
- `=` is equal sign
- `_` is underscore
- `%` is percentage sign
- `:` is colon
- `-` is dash
- `*` is repeat the [...] group
- `sort -u`  will sort & remove any duplicates
