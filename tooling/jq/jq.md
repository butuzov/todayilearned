# `jq`

* [Working with JSON in bash using `jq`](https://medium.com/cameron-nokes/working-with-json-in-bash-using-jq-13d76d307c4)
* [Custom Functions](https://medium.com/google-cloud/kubernetes-engine-master-node-versions-b5ecd9ed0b35)

### Functions

```bash
# get keys of the passed json, a, and b
echo '{ "a": 1, "b": 2 }' | jq 'keys | .[]'

# length of the object (3)
echo '[1,2,3]' | jq 'length'
```

## General 

```
# Grabbing Chack Noris Joke 
curl -LO http://api.icndb.com/jokes/random
cp random randjoke.json

# getting joke
cat randjoke.json  | jq '.value.joke'

# or reading file (not quoted)
jq '.value.joke' -r randjoke.json
```

## General (kubernetes examples)

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

### Accessing elements

```bash
# get the data
curl 'http://stash.compciv.org/congress-twitter/json/nancy-pelosi.json' > test.json

# all of that
cat test.json | jq '.'
cat test.json | jq '.name'
cat test.json | jq '.name.last'

# access to array data
cat test.json | jq '.terms[].'
cat test.json | jq '.terms[].state'
#first.
cat test.json | jq '.terms[0].state'
# secodn from the end
cat test.json | jq '.terms[-2].state'
# last el
cat test.json | jq '.terms[14].state'
# null
cat test.json | jq '.terms[19].state'
```

# output 

```
cat test.json | jq --raw-output '.terms[14].state'
```

# unprocessed
https://medium.com/free-code-camp/how-to-transform-json-to-csv-using-jq-in-the-command-line-4fa7939558bf

```python
jq -r '(map(keys) | add | unique) as $cols | map(. as $row | $cols | map($row[.])) as $rows | $cols, $rows[] | @csv'
```

```python
jq -r '(map(keys) | add | unique) as $cols | map(. as $row | $cols | map($row[.])) as $rows | $cols, $rows[] | @csv'
```

      File "<ipython-input-2-4be1047b68ec>", line 1
        jq -r '(map(keys) | add | unique) as $cols | map(. as $row | $cols | map($row[.])) as $rows | $cols, $rows[] | @csv'
                                                                                                                           ^
    SyntaxError: invalid syntax

* https://stedolan.github.io/jq/
* https://stedolan.github.io/jq/tutorial/
