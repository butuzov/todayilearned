# [sort](https://golang.org/pkg/sort/)

* [Performance: Slices of structs vs. slices of pointers to structs](https://stackoverflow.com/questions/27622083)
* [Performance: Sorting Slice vs Sorting Type (of Slice) with Sort implementation](https://stackoverflow.com/questions/54276285)

### Tests Results

 Type                                | Implementation                    | Sort Speed
-------------------------------------|-----------------------------------|-------------
Slice Of Structs                     | Sort.Slice                        | slow
Struct (Slice Of Structs)            | Sort interface implementation     | fast
Slice Of Pointer to Structs          | Sort.Slice                        | fast
Struct(Slice Of Pointer to Structs)  | Sort interface implementation                       | middle

```bash
    > go test -bench=.
    goos: darwin
    goarch: amd64
    BenchmarkSlice-4                   	 3000000	       542 ns/op
    BenchmarkStruct-4                  	 5000000	       318 ns/op
    BenchmarkSlicePointers-4           	 5000000	       280 ns/op
    BenchmarkStructOfSlicePointers-4   	 5000000	       321 ns/op
```
