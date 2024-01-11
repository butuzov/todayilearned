<!-- menu: "`jq`" -->
# jq

> A jq program is a "filter": it takes an input, and produces an output. There are a lot of builtin filters for extracting a particular field of an object, or converting a number to a string, or various other standard tasks.

## `./jq` Online

- https://jqplay.org/
- https://jqlang.github.io/jq/

## Example

https://github.com/jqlang/jq/wiki/Cookbook#filter-objects-based-on-the-contents-of-a-key

```bash
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


<!-- - https://medium.com/cameron-nokes/working-with-json-in-bash-using-jq-13d76d307c4
- https://medium.com/google-cloud/kubernetes-engine-master-node-versions-b5ecd9ed0b35
- https://github.com/butuzov/todayilearned/blob/v2/tooling/jq/jq.md
- https://medium.com/free-code-camp/how-to-transform-json-to-csv-using-jq-in-the-command-line-4fa7939558bf
- https://stackoverflow.com/questions/18592173/
- https://stackoverflow.com/questions/28164849/
- https://stackoverflow.com/questions/44656515/how-to-remove-double-quotes-in-jq-output-for-parsing-json-files-in-bash
- https://stackoverflow.com/questions/33057420/jq-select-multiple-conditions/33059058#33059058 -->




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
TODO: Object Identifier-Index
```

## `TODO`: Types and Values

```shell

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

## `TODO`:  Conditionals and Comparisons

## `TODO`:  Regular expressions

## `TODO`:  Advanced features

### Math

```shell
# addition
echo "[1,2,3,4,5]" | jq '.'

```

## `TODO`: I/O
## `TODO`: Streaming
## `TODO`: Assignment
## `TODO`: Modules

### Colors

You can customize colored output by setting: [`JQ_COLORS`](https://jqlang.github.io/jq/manual/#colors) (e.g `"JQ_COLORS=1;30:0;39:0;39:0;39:0;32:1;39:1;39"`)

## Other...

```shell
curl https://api.github.com/users/butuzov/received_events


# Array To Multi Document Json
jq '.[] | select(.)'
jq '.[] | select(.type == "CreateEvent") | [.]'
# Subselecting Element from Selected Object into Array
jq '.[] | select(.type == "CreateEvent").actor | [.]'
```
