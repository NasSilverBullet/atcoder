package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "regexp"
  "strings"
)

func main() {
  str := readLinesfromStdin()
  contest, number, task := getDirNames(str)
  fmt.Println(contest, number, task)
}

func readLinesfromStdin() string {
  io.WriteString(os.Stdout, "e.g. ) ABC122 A. Double Helix\n")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  return scanner.Text()
}

func getDirNames(rawStr string) (contest string, number string, task string) {
  rep := regexp.MustCompile(`\s+`)
  compStr := rep.Split(rawStr, -1)
  contest = compStr[0][:3]
  number = "T" + compStr[0][3:]
  nAlNum := regexp.MustCompile(`\W`)
  task = nAlNum.ReplaceAllString(strings.ToLower(strings.Join(compStr[1:], "_")), "")
  return
}
