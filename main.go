package main

import (
	"fmt"
	"go_leetcode/same_direct_double"
)

func main() {
	s := []int{1, 2, 3, 1, 2, 3, 1, 2}
	fmt.Printf("最多k重复的子数组数目为:%v", same_direct_double.MaxSubarrayLength(s, 2))
}
