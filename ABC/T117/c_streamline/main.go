package main

import (
  "fmt"
  "math"
  "sort"
)

func main() {
  var n, m int
  fmt.Scan(&n, &m)
  x := make([]int, m)
  for i := 0; i < m; i++ {
    fmt.Scan(&x[i])
  }

  if n >= m {
    fmt.Println(0)
    return
  }

  sort.Ints(x)
  diff := make([]int, m-1)
  for i := 0; i < m-1; i++ {
    diff[i] = x[i+1] - x[i]
  }
  for i := 0; i < m-1; i++ {
    diff[i] = int(math.Abs(float64(diff[i])))
  }
  sort.Ints(diff)

  ans := x[len(x)-1] - x[0]
  for _, v := range diff[m-(n-1)-1:] {
    ans -= v
  }
  fmt.Println(ans)
}
