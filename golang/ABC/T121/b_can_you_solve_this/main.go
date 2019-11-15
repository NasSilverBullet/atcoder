package main

import "fmt"

func main() {
  var n, m, c int
  fmt.Scan(&n, &m, &c)
  b := make([]int, m)
  for i := 0; i < m; i++ {
    fmt.Scan(&b[i])
  }
  a := make([]int, m)
  count := 0
  for i := 0; i < n; i++ {
    sum := c
    for j := 0; j < m; j++ {
      fmt.Scan(&a[j])
      sum += a[j] * b[j]
    }
    if sum > 0 {
      count++
    }
  }
  fmt.Println(count)
}
