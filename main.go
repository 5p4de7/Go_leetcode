package main

import (
	"fmt"
	"go_leetcode/two_sum_three_sum"
)

func main() {
	a := []int{-1, 1, 2, 3, 1}
	b := two_sum_three_sum.CountPairs(a, 2)
	fmt.Printf("%v\n", b)
}
