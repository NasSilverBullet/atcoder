package main

import (
	"fmt"
	"os"

	"github.com/NasSilverBullet/atcoder/golang"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)

}

func run() error {
	if err := golang.Arc004_1(); err != nil {
		return err
	}
	return nil
}
