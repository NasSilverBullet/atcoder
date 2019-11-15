package main

import "fmt"

func main() {
  var (
    n, q int
    s    string
  )
  fmt.Scan(&n, &q, &s)
  t := make([]int, n)
  for i := 1; i < n; i++ {
    t[i] = t[i-1]
    if s[i] == 'C' && s[i-1] == 'A' {
      t[i]++
    }
  }
  for i := 0; i < q; i++ {
    var l, r int
    fmt.Scan(&l, &r)
    fmt.Println(t[r-1] - t[l-1])
  }
}
