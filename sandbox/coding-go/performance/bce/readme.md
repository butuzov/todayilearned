# Bounds Check Elimination

`eli5`  Bounds Check Elimination, Index Bounds checks...

## Reading List

* https://go101.org/article/bounds-check-elimination.html
* https://docs.google.com/document/d/1vdAEAjYdzjnPA9WDOQ1e4e05cYVMpqSxJYZT33Cqw2g/edit
* [Aliaksandr Valialkin: Использование unsafe в Go: плюсы и минусы: BCE Part](https://youtu.be/rxGgdAGhE9k?t=1909)
* [Agniva De Sarker - Common Patterns for Bounds Check Elimination](https://youtu.be/5toTS6kSWHA)
* https://en.wikipedia.org/wiki/Bounds-checking_elimination


## Checking


```shell
alias bce="go tool compile -d=ssa/check_bce/debug=1 ${1}"
bce example.go
```

Messages:

- `IsInBounds`
- `IsSliceInBounds`


### Using `tasks.json` in VSCode

1. Command Pallet
2. Task: Run Task
3. GO: Bound Checks

```json
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "Go: Bounds Check",
            "type": "shell",
            "command": "go tool compile -d=ssa/check_bce/debug=1 ${file} && for _ in '${fileDirname}/*.o'; do unlink $_ ; done",
            "presentation": {
                "echo": true,
                "reveal": "always",
                "focus": true,
                "panel": "shared",
                "showReuseMessage":false,
                "clear": true,
            },
            "problemMatcher": [
                "$go"
            ]
        }
    ]
}
```

### Example Code in this folder.

#### Использование unsafe в Go: плюсы и минусы: BCE Part

* [Aliaksandr Valialkin: Использование unsafe в Go: плюсы и минусы: BCE Part](https://youtu.be/rxGgdAGhE9k?t=1909)

{{% list "valyala.go" %}}

#### Patterns & Examples

{{% list "examples_1.go,examples_2.go,examples_3.go,examples_4.go,pattern_1.go,pattern_2.go,pattern_3.go,pattern_4.go,pattern_5.go" %}}
