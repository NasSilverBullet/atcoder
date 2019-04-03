package main

import (
  "bufio"
  "io"
  "os"
  "fmt"
)

func main() {
  str := readLinesfromStdin()
  fmt.Println(str)
}

func readLinesfromStdin() string {
  io.WriteString(os.Stdout, "<Contest name> <Question number>\n")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  return scanner.Text()
}
