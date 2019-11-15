package main

import "fmt"

func main() {
  var n int
  fmt.Scan(&n)
  a := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&a[i])
  }
  mod := a[0]
  for _, v := range a {
    mod = gcd(mod, v)
  }
  fmt.Println(mod)
}

func gcd(a, b int) int {
  var mod int
  for a%b > 0 {
    mod = a % b
    a = b
    b = mod
  }
  return b
}
