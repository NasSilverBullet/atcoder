## Sort
```go
// unstable sort
sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
// stable sort
sort.SliceStable(A, func(i, j int) bool { return A[i] < A[j] })
// struct の slice を Stable sort
sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
```
