package main

import "fmt"

func main() {
  var a, b, k int
  fmt.Scan(&a, &b, &k)
  var ans int
  for i := 100; k > 0; i-- {
    if a%i == 0 && b%i == 0 {
      ans = i
      k--
    }
  }
  fmt.Println(ans)
}
