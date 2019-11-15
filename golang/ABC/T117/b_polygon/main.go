package main

import "fmt"

func main() {
  var n int
  fmt.Scan(&n)

  var num, sum, max int
  for i := 0; i < n; i++ {
    fmt.Scan(&num)
    sum += num
    if num > max {
      max = num
    }
  }
  if sum-max > max {
    fmt.Println("Yes")
    return
  }
  fmt.Println("No")
}
