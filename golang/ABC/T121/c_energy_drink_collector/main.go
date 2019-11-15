package main

import (
  "fmt"
  "sort"
)

func main() {
  var n, m int
  fmt.Scan(&n, &m)
  shops := make([][]int, n)
  for i := 0; i < n; i++ {
    shops[i] = make([]int, 2)
    fmt.Scan(&shops[i][0], &shops[i][1])
  }
  sort.Slice(shops, func(i, j int) bool { return shops[i][0] < shops[j][0] })
  var sum, take int
  for i := 0; m > 0; i++ {
    if shops[i][1] > m {
      take = m
    } else {
      take = shops[i][1]
    }
    sum += shops[i][0] * take
    m -= take
  }
  fmt.Println(sum)
}
