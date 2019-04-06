package main

import (
  "fmt"
  "math"
)

var N, A, B, C int
var L []int

func main() {
  fmt.Scan(&N, &A, &B, &C)
  L = make([]int, N)
  for i := 0; i < N; i++ {
    fmt.Scan(&L[i])
  }
  fmt.Println(search(0, 0, 0, 0))
}

func search(cur, a, b, c int) int {
  if cur == N {
    if getMin(a, b, c) <= 0 {
      return math.MaxInt32
    }
    return int(math.Abs(float64(a-A)) + math.Abs(float64(b-B)) + math.Abs(float64(c-C)) - 30)
  }

  ret0 := search(cur+1, a, b, c)
  ret1 := search(cur+1, a+L[cur], b, c) + 10
  ret2 := search(cur+1, a, b+L[cur], c) + 10
  ret3 := search(cur+1, a, b, c+L[cur]) + 10
  return getMin(ret0, ret1, ret2, ret3)
}

func getMin(nums ...int) (min int) {
  min = nums[0]
  for i := 0; i < len(nums); i++ {
    min = int(math.Min(float64(min), float64(nums[i])))
  }
  return
}
