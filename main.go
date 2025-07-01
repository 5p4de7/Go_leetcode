package main

import (
	"fmt"
	"go_leetcode/sliding_window"
)

func main() {
	s := []int{61, 23, 38, 23, 56, 40, 82, 56, 82, 82, 82, 70, 8, 69, 8, 7, 19, 14, 58, 42, 82, 10, 82, 78, 15, 82}
	fmt.Printf("子数组数目为:%v\n", sliding_window.CountSubarrays(s, 2))
	s1 := "你还没"
	for _, x := range s1 {
		fmt.Printf("%v,%c\n", s1[0], x)
	}
}
