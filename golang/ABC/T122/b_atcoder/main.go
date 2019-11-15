package main

import "fmt"

func main() {
  bs := map[byte]int{
    'A': 0,
    'C': 1,
    'G': 2,
    'T': 3,
  }
  var s string
  fmt.Scan(&s)
  max := 0
  for i := range s {
    count := 0
    for j := i; j < len(s); j++ {
      if _, ok := bs[s[j]]; ok {
        count++
      } else {
        i = j
        break
      }
    }
    if max < count {
      max = count
    }
  }
  fmt.Println(max)
}
