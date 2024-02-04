<!-- menu: "`jq`" -->
# jq

> A jq program is a "filter": it takes an input, and produces an output. There are a lot of builtin filters for extracting a particular field of an object, or converting a number to a string, or various other standard tasks.

## `./jq` Online

- https://jqplay.org/
- https://jqlang.github.io/jq/
- https://manpages.ubuntu.com/manpages/xenial/man1/jq.1.html
- https://github.com/jqlang/jq/wiki/Cookbook

## Example

More Simple steps to begin with...

```shell
> jq -V
jq-1.7

# Let's grab a sources
> curl https://api.github.com/users/butuzov/received_events > o.json

# We can pipe source into jq with a few ways
> cat "o.json" | jq 'length'
30
> jq -f o.json 'length'
30
> jq 'length' o.json
30
# Using null input + piping into another jq to minify it.
> jq --null-input --argjson test null '.test=$test'  | jq -r tostring

# How this help structured:
# each call to jquery will precent ...| and it means we runnin cat 'o.json'
> cat 'o.json'

# writing line below is same as cat o.json | jq -M '.[]', if other input not provided
> ...| jq '.'
# json output
{...}

```
### Kubernetes Examples

```bash
alias k="kubectl $@"
# initial
k create deployment nginx --image=nginx --dry-run -o yaml | k create -f -
k scale deployment nginx --replicas=4
k get pods

# k get pods -o json

# how many pods do i have?
k get pods -o json | jq '.items[]' | jq 'length'

# output names
k get pods -o json | jq '.items[].metadata.name'
# output names as json list
k get pods -o json | jq '[.items[].metadata.name]'
# how many pods do i have?
k get pods -o json | jq '[.items[].metadata.name]' | jq 'length'
# less piping plese
k get pods -o json | jq '[.items[].metadata.name] | length'

# new json creation
k get pods -o json | jq '{items:[{name:.items[].metadata.name, termination: .items[].spec.terminationGracePeriodSeconds}]}'

# pre filtered
k get pods -o json | jq '.items[] | {name:.metadata.name, termination: .spec.terminationGracePeriodSeconds}'
# more filters (greater or equal then 30)
k get pods -o json | jq '.items[] | {name:.metadata.name, termination: .spec.terminationGracePeriodSeconds} | select(.termination|tonumber >= 30)'
# more filters (not null)
k get pods -o json | jq '.items[] | {name:.metadata.name, termination: .spec.terminationGracePeriodSeconds} | select(.termination != null)'

```

## Invoking jq

- Unix shells: `jq '.["foo"]'`
- Powershell: `jq '.[\"foo\"]'`
- Windows command shell: `jq ".[\"foo\"]"`


## Basic filters

```shell
export j='-M o.json'
> jq '.' $j                             # Identity
[...]
> echo 1 | jq '[., tojson, tostring]' | jq -r tostring
[1,"1","1"]
> echo 1 | jq '. < 0.12345678901234567890123456788'
false
> echo [20] | jq 'map([., . == 1]) | tojson'
"[[20,false]]"
>

```
## TODO: Object Identifier-Index

## Types and Values

```shell
> echo -n '[1, "1"]' | jq '.[] | tonumber'
1
1
> echo -n '[1, "1"]' | jq '.[] | tostring'
"1"
"1"
> echo -n '[0, false, [], {}, null, "hello"]' | jq 'type'
array
> echo -n '[0, false, [], {}, null, "hello"]' | jq -c 'map(type)'
["number","boolean","array","object","null","string"]

# Other...
$json=$(curl -s https://api.github.com/users/butuzov/received_events)
# Array To Multi Document Json
echo $json | jq '.[] | select(.)'
echo $json | jq '.[] | select(.type == "CreateEvent") | [.]'
# Subselecting Element from Selected Object into Array
echo $json | jq '.[] | select(.type == "CreateEvent").actor | [.]'
```

## Builtin operators and functions

### @CSV

This is copy of https://e.printstacktrace.blog/how-to-convert-json-to-csv-from-the-command-line/

```shell
# json to csv
> jq -r 'map({id,title,url,company,location}) | (first | keys_unsorted) as $keys | map([to_entries[] | .value]) as $rows | $keys,$rows[] | @csv' jobs.json > jobs.csv
```

1. We run `jq -r` to output raw strings (without double quotes.)
1. In the filter part, we pipe multiple filters together, starting with `map({id,title,url,company,location})`. This filter instructs `jq` which keys we want to extract from the input JSON file.
1. Then we use `(first | keys_unsorted)` as `$keys` filter which takes the first object, extracts its keys and stores them under the `$keys` variable as an array.
1. Next, we use `map([to_entries[] | .value]) as $rows` filter which converts every key-value entry like `"foo": "bar"` into an object like `{"key":"foo","value":"bar"}` so we can extract only values as an array and store it in a $rows variable.
1. Once we do it, we can use `$keys,$rows[]` filter to put keys and rows together and then pipe it with the `@csv` filter to convert JSON objects to CSV rows.

### CVS parsing

```shell
# csv to json
> ./install.sh
> jq --slurp --raw-input --raw-output \
   'split("\n") | .[1:] | map(split(",")) |
    map({
        "PassengerId": .[0],
        "Survived": "\(.[1])",
        "Pclass": .[2],
      })' \
  example_titanic.csv
```

### Other builtins

```shell
# print sorted keys
curl -s http://localhost:9200/_aliases | jq 'keys'
# print keys
curl -s http://localhost:9200/_aliases | jq 'keys_unsorted'
# object/array length
curl -s http://localhost:9200/_aliases | jq 'length'
# check for key in array
curl -s http://localhost:9200/_aliases | jq 'has("nifi-req-2024.01.11")'
# in set
echo -n '["foo", "bar"]' | jq '.[] | in({"foo": 42})'
```

## Conditionals and Comparisons

```shell
> json='[{"a":1,"b":10,"c":9},{"a":2,"b":25,"c":5}]'
> echo -n $json | jq -c '.[] | select((.a <= 1)).a'
1
> echo -n $json | jq -c '.[] | select((.a <= 1) and (.c=9)).b'
10
> echo -n $json | jq -c '.[] | select((.a <= 1) or (.c%3>0)).b'
10
25
```

## Regular expressions

The jq regex filters are defined so that they can be used using one of these patterns:

```
STRING | FILTER(REGEX)
STRING | FILTER(REGEX; FLAGS)
STRING | FILTER([REGEX])
STRING | FILTER([REGEX, FLAGS])
```
where:

- `STRING`, `REGEX`, and `FLAGS` are jq strings and subject to jq string interpolation;
- `REGEX`, after string interpolation, should be a valid regular expression;
- `FILTER` is one of `test`, `match`, or `capture``, as described below.

`FLAGS` is a string consisting of one of more of the supported flags:

- `g` - Global search (find all matches, not just the first)
- `i` - Case insensitive search
- `m` - Multi line mode (. will match newlines)
- `n` - Ignore empty matches
- `p` - Both s and m modes are enabled
- `s` - Single line mode (^ -> \A, $ -> \Z)
- `l` - Find longest possible matches
- `x` - Extended regex format (ignore whitespace and comments)


```shell
# Test
> jq -n '"a" | test("a")'
true
> jq -n '"a" | test("a", "x")'
true
false
> echo -n '["xabcd", "ABC"]' | jq '.[] | test("a b c # spaces are ignored"; "ix")'
true
true

# Match
echo -n '"abc abc"' | jq 'match("(abc)+"; "g")'
{
  "offset": 0,
  "length": 3,
  "string": "abc",
  "captures": [
    {
      "offset": 0,
      "length": 3,
      "string": "abc",
      "name": null
    }
  ]
}
{
  "offset": 4,
  "length": 3,
  "string": "abc",
  "captures": [
    {
      "offset": 4,
      "length": 3,
      "string": "abc",
      "name": null
    }
  ]
}
> echo -n 'abc' | jq '[ match("."; "g")] | length'
3

# Capture
> echo -n '"xyzzy-14"' | jq 'capture("(?<a>[a-z]+)-(?<n>[0-9]+)")'
{
  "a": "xyzzy",
  "n": "14"
}

# Split
> echo -n '"ab,cd, ef"' | jq 'split(", *"; null)'
["ab","cd","ef"]
# Splits
> echo -n '"ab,cd, ef"' | jq 'splits(", *")'
"ab"
"cd"
"ef"

# Sub
> echo -n '"ab,cd, ef"' |jq 'sub("[^a-z]*(?<x>[a-z]+)"; "Z\(.x)"; "g")'
"ZabZcdZef"

> echo -n '"Ab"' | jq '[sub("(?<a>.)"; "\(.a|ascii_upcase)", "\(.a|ascii_downcase)")]'
[
  "Ab",
  "ab"
]

# GSUB
> echo -n '"Ab"' | jq 'gsub("(?<x>.)[^a]*"; "+\(.x)-")'
"+A-"
```


## `TODO`:  Advanced features

### Environment

```shell
> jq -n 'env'
...dumps json of env...

> jq -n 'env | keys | length'
45

> jq -n 'env | with_entries(select ((.key|startswith("TERM_")) or .key == "DOCKER_CONTAINER_VERSION_TAG"))'
{
  "TERM_SESSION_ID": "w0t0p0:D97EB7AC-19B8-4FEF-BB65-E892404B3F6C",
  "TERM_PROGRAM_VERSION": "3.4.23",
  "TERM_PROGRAM": "iTerm.app"
}
```

### Math

- One-іnput C functions: `acos`, `acosh`, `asin`, `asinh`, `atan`, `atanh`, `cbrt`, `ceil`, `cos`, `cosh`, `erf`, `erfc`, `exp`, `exp10`, `exp2`, `expm1`, `fabs`, `floor`, `gamma`, `j0`, `j1`, `lgamma`, `log`, `log10`, `log1p`, `log2`, `logb`, `nearbyint`, `pow10`, `rint`, `round`, `significand`, `sin`, `sinh`, `sqrt`, `tan`, `tanh`, `tgamma`, `trunc`, `y0`, `y1`.
- Two-input C math functions: `atan2`, `copysign`, `drem`, `fdim`, `fmax`, `fmin`, `fmod frexp`, `hypot`, `jn`, `ldexp`, `modf`, `nextafter`, `nexttoward`, `pow`, `remainder`, `scalb`, `scalbln`, `yn`.
- Three-input C math functions: `fma`.

```shell
> echo "[1,2,3,4,5]" | jq  '. | add'
15
> echo "[1,2,3,4,5]" | jq  '. | add | sqrt'
3.872983346207417
```

## `TODO`: I/O
## `TODO`: Streaming
## `TODO`: Assignment
## `TODO`: Modules

### Colors

You can customize colored output by setting: [`JQ_COLORS`](https://jqlang.github.io/jq/manual/#colors) (e.g `"JQ_COLORS=1;30:0;39:0;39:0;39:0;32:1;39:1;39"`)

