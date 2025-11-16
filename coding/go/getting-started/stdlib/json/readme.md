# `json`

## Reading

- https://dave.cheney.net/high-performance-json.html
- TODO(butuzov): `json/v2`

## Tooling

- https://jsonformatter.org/json-to-go
- https://mholt.github.io/json-to-go/

## Libraries

- https://github.com/Jeffail/gabs

## Recipes

### `ordered json`

See next example of `inlining`, with help of sorted slics we can have sorted keys in json.

{{% list "json/unordered.go" %}}

### Skipping if empty (zero value) `,omitempty`

{{% list "json/omitempty.go" %}}

```shell
>> {"A":"","b":"","e":""}
```

### Skipping with `-`

`json:"-"` Helps to skip Records population

{{% list "json/skip_json_fileds.go" %}}

### htmlquoting

{{% list "json/htmlquoting.go" %}}

### json.Raw

{{% list "json/raw.go" %}}
