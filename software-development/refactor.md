- [Codebase Refactoring (with help from Go)](https://talks.golang.org/2016/refactor.article) by Russ Cox [`hn`](https://news.ycombinator.com/item?id=13091634)
- [Refactoring Go Code](https://matt.aimonetti.net/posts/2014-04-refactoring-go-code/)
- [Refactoring With Go Fmt](https://spf13.com/post/go-fmt/)
- [Go code refactoring : the 23x performance hunt](https://medium.com/@val_deleplace/go-code-refactoring-the-23x-performance-hunt-156746b522f7)
- [gorename: easy refactoring tool for Golang](https://texlution.com/post/gorename/)

# Refactoring Techniques [refactoring.guru](https://refactoring.guru/refactoring/techniques)

- Add Parameter
- Bidirectional 
- Change Bidirectional Association to  Unidirectional
- Change Reference to Value
- Change Unidirectional Association to Bidirectional
- Change Value to Reference
- Clauses
- Collapse Hierarchy
- Consolidate Conditional Expression
- Consolidate Duplicate Conditional Fragments
- Constant
- Decompose Conditional
- Duplicate Observed Data
- Dynamic Method Definition
- Eagerly Initialized Attribute
- Encapsulate Collection
- Encapsulate Downcast
- Encapsulate Field
- Extract Class
- Extract Interface
- Extract Method
- Extract Module
- Extract Subclass
- Extract Superclass
- Extract Surrounding Method
- Extract Variable
- Form Template Method
- Hide Delegate
- Hide Method
- Inline Class
- Inline Method
- Inline Module
- ~~Inline Temp~~
- Introduce Assertion
- Introduce Class Annotation
- Introduce Expression Builder
- Introduce Foreign Method
- Introduce Gateway
- Introduce Local Extension
- Introduce Named Parameter
- Introduce Null Object
- Introduce Parameter Object
- Isolate Dynamic Receptor
- Lazily Initialized Attribute
- Method Definition
- Move Eval from Runtime to Parse Time
- Move Field
- Move Method
- Parametize Method
- Preserve Whole Object
- Pull Down Field
- Pull Down Method
- Pull Up Constructor Body
- Pull Up Field
- Pull Up Method
- Recompose Conditional
- Remove Assignments to Parameters
- Remove Assignments to Parameters
- Remove Control Flag
- Remove Middle Man
- Remove Named Parameter Remove Parameter
- Remove Setting Method
- Remove Unused Default Parameter
- ~~Rename Method~~
- Replace Abstract Superclass with Module
- Replace Array with Object
- Replace Conditional with Polymorphism
- Replace Constructor with Factory Method
- Replace Data Value with Object
- Replace Delegation With Hierarchy
- Replace Delegation With Inheritance
- Replace Dynamic Receptor with Dynamic
- Replace Error Code with Exception
- Replace Exception with Test
- Replace Hash with Object
- Replace Inheritance With Delegation
- Replace Loop with Collection Closure Method
- Replace Magic Number with Symbolic
- Replace Method With Method Object
- Replace Nested Conditional with Guard
- Replace Parameter with Explicit Methods
- Replace Parameter with Method
- Replace Record with Data Class
- Replace Subclass with Fields
- Replace Temp with Chain
- Replace Temp with Query
- Replace Type Code with Class
- Replace Type Code with Module Extention
- Replace Type Code with Polymorphism
- Replace Type Code with State/Strategy
- Replace Type Code with Subclasses
- Self Encapsulate Field
- Separate Query from Modifier
- Split Temporary Variable
- Substitute Algorithm

## Change Value to Reference

todo


## Inline Temp

```go
# Before 
func HasDicount(order Order) bool {
    int basePrice = order.BasePrice()
    return basePrice > 1000
}

# After 
func HasDicount(order Order) bool { 
    return order.BasePrice() > 1000
}
```

## Rename Method

```go
# Before 
func (t *this) rmdir(location string) error {
    os.RemoveAll(location)
}

# After
func (t *this) removeDirectory(location strin) error {
   os.RemoveAll(location)
}
```

```go
https://refactoring.fm/archive
```