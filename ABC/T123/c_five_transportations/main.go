package main

import (
  "fmt"
  "math"
)

func main() {
  var n int
  fmt.Scan(&n)
  A := make([]int, 5)
  min := math.MaxInt64
  for i := 0; i < 5; i++ {
    fmt.Scan(&A[i])
    if A[i] < min {
      min = A[i]
    }
  }
  fmt.Println((n+min-1)/min + 4)
}
