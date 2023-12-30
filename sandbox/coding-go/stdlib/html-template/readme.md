# `html/template`

- https://golang.org/pkg/text/template/
- https://golang.org/pkg/html/template/

```go
import "log"
import "os"
```

```go
import "html/template"
```

```go
var tpl = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
------------------------------------------------------
`
```

```go
t := template.Must(template.New("tpl").Parse(tpl))
```

```go
data := []struct {
	Name, Gift string
	Attended   bool
}{
	{"Aunt Mildred", "bone china tea set", true},
	{"Uncle John", "moleskin pants", false},
	{"Cousin Rodney", "", false},
}
```

```go
for _, r := range data {
	err := t.Execute(os.Stdout, r)
	if err != nil {
		log.Println("executing template:", err)
	}
tpl}

> 
> Dear Aunt Mildred,
> 
> It was a pleasure to see you at the wedding.
> Thank you for the lovely bone china tea set.
> 
> Best wishes,
> Josie
> ------------------------------------------------------
> 
> Dear Uncle John,
> 
> It is a shame you couldn't make it to the wedding.
> Thank you for the lovely moleskin pants.
> 
> Best wishes,
> Josie
> ------------------------------------------------------
> 
> Dear Cousin Rodney,
> 
> It is a shame you couldn't make it to the wedding.
> 
> Best wishes,
> Josie
> ------------------------------------------------------
```

## Binary comparison operators (defined as functions):

* `eq` - Returns the boolean truth of `arg1` == `arg2`
* `ne` - Returns the boolean truth of `arg1` != `arg2`
* `lt` - Returns the boolean truth of `arg1` < `arg2`
* `le` - Returns the boolean truth of `arg1` <= `arg2`
* `gt` - Returns the boolean truth of `arg1` > `arg2`
* `ge` - Returns the boolean truth of `arg1` >= `arg2`

```go
tpl  := `{{if or (gt .Hour 21) (le .Hour 7) }}Good Night{{else}}Good Other Time of the 24 hours{{end}}`
t := template.Must(template.New("tpl").Parse(tpl))

t.Execute(os.Stdout, struct{
    Hour    int 
}{4})

> Good Night
```

### Descreet Functions 

* `and` 
* `not`
* `or`

```go
tpl  := `{{if and .IsRed .IsRound}}{{.Name}}{{end}}`
t := template.Must(template.New("tpl").Parse(tpl))

t.Execute(os.Stdout, struct{
    Name    string 
    IsRed   bool
    IsRound bool
}{"Red Ball", true, true})

> Red Ball
```

```go
tpl  := `{{if or .IsRed .IsRound}}{{.Name}}{{end}}`
t := template.Must(template.New("tpl").Parse(tpl))

t.Execute(os.Stdout, struct{
    Name    string 
    IsRed   bool
    IsRound bool
}{"Blue Ball", false, true})

> Blue Ball
```

## Other Functions 

* `call` 
* `html` 
* `index` 
* `slice` 
* `js` 
* `len`
* `print` 
* `printf` 
* `println`
* `urlquery` 

## Range Over

```go
tpl  := `{{range .}} - {{.Name}} 
{{end}}`
t := template.Must(template.New("tpl").Parse(tpl))

t.Execute(os.Stdout, []struct{
    Name    string
}{
    {"Pickachu"},
    {"Ratata"},
    {"Mr. Mime"},
})

> - Pickachu
> - Ratata
> - Mr. Mime
```

## Try 

```go
tpl  := `{{.Name}}`
t := template.Must(template.New("tpl").Parse(tpl))

t.Execute(os.Stdout, struct{
    Name    string
}{ "pickachu"})

> pickachu
```