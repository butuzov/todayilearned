# `json`

## Reading

* https://dave.cheney.net/high-performance-json.html

## Tooling

* https://jsonformatter.org/json-to-go
* https://mholt.github.io/json-to-go/


## Libraries

* https://github.com/Jeffail/gabs

## Recipes

### `ordered json`

See next example of `inlining`, with help of sorted slics we can have sorted keys in json.

{{% list "unordered.go" %}}

### `,inline`
{{% list "skip_json_fileds.go" %}}

### `,omitempty`

{{% list "omitempty.go" %}}
```shell
>> {"A":"","b":"","e":""}
```
### `-`

`json:"-"` Helps to skip Records population

{{% list "skip_json_fileds.go" %}}


### htmlquoting

{{% list "htmlquoting.go" %}}

### json.Raw

{{% list "raw.go" %}}
