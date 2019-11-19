package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)

}

func run() error {
	cmd := exec.Command("sh", "-c", "~/.pyenv/versions/pypy3.6-7.2.0/bin/pypy3 ~/go/src/github.com/NasSilverBullet/atcoder/abc051_b.py")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if _, err := io.WriteString(stdin, "2 2"); err != nil {
		return err
	}

	if err := stdin.Close(); err != nil {
		return err
	}

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	fmt.Print(string(out))
	return nil
}
