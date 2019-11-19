package golang

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
	dir := contest + "/" + number + "/" + task
	if _, err := os.Stat(dir); err != nil {
		fmt.Printf("%v\n", err)
		fmt.Printf("mkdir %s\n", dir)
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println(err)
	}
	os.Chdir(dir)
	fmt.Printf("cd %s\n", dir)
	fileName := "main.go"
	_, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
		fmt.Printf("touch %s\n", fileName)
	}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer file.Close()
	fmt.Printf("%s/%s is created!\n", dir, fileName)
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
