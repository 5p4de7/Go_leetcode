package main

import (
	"fmt"
	"go_leetcode/two_sum_three_sum"
)

func main() {
	a := []int{24, 3, 82, 22, 35, 84, 19}
	b := two_sum_three_sum.TriangleNumber(a)
	fmt.Printf("ans is %v\n", b)
}
