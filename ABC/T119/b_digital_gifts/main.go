package main

import "fmt"

func main() {
  var n int
  fmt.Scan(&n)
  currency := map[string]float64{
    "JPY": 1.0,
    "BTC": 380000.0,
  }
  var (
    x, sum float64
    u      string
  )
  for i := 0; i < n; i++ {
    fmt.Scan(&x, &u)
    sum += x * currency[u]
  }
  fmt.Println(sum)
}
