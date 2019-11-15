package main

import "fmt"

func main() {
  var a, b, c int
  fmt.Scan(&a, &b, &c)
  listen := 0
  if b/a < c {
    listen = b / a
  } else {
    listen = c
  }
  fmt.Println(listen)
}
