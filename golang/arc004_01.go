package golang

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Arc004_1() error {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		return err
	}

	s := make([][]int, n)
	for i := 0; n > i; i++ {
		sc.Scan()
		astr := strings.Split(sc.Text(), " ")

		x, err := strconv.Atoi(astr[0])
		if err != nil {
			return err
		}

		y, err := strconv.Atoi(astr[1])
		if err != nil {
			return err
		}

		s[i] = []int{x, y}
	}

	fmt.Println(runArc004_1(s))
	return nil
}

func runArc004_1(ss [][]int) float64 {
	var max float64
	for _, s1 := range ss {
		for _, s2 := range ss {
			xd, yd := float64(s1[0]-s2[0]), float64(s1[1]-s2[1])
			dist := math.Sqrt(xd*xd + yd*yd)
			if dist > max {
				max = dist
			}
		}
	}
	return max
}
