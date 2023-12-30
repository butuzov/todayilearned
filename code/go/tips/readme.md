# Tips & Tricks

## Unterstanding Compilators and Go (Go's Go) Code

by Egon Elbre
- write a small compiler from scratch (e.g. something along the lines of https://compilerbook.com/) -- jumping straight into large compiler can be difficult
- start reading through each patch (or reading random changes on https://go-review.googlesource.com/q/status:open+-is:wip) -- these help to understand relation between "issue <-> code that needs to change"
- try debugging and understanding things with "NeedsInvestigation" https://github.com/golang/go/issues?q=is%3Aissue+is%3Aopen+label%3ANeedsInvestigation -- don't worry when the investigation doesn't lead anywhere, the important part is getting more familiarity with things
- try fixing things with "help wanted" label https://github.com/golang/go/issues?q=is%3Aopen+is%3Aissue+label%3A%22help+wanted%22 -- same here, don't worry if you fail to fix
