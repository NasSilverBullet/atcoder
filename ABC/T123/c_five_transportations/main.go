package main

import (
  "fmt"
  "math"
)

func main() {
  var n, a int
  fmt.Scan(&n)
  min := math.MaxInt64
  for i := 0; i < 5; i++ {
    fmt.Scan(&a)
    if a < min {
      min = a
    }
  }
  fmt.Println((n+min-1)/min + 4)
}
