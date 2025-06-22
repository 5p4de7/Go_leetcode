package main

import (
	"go_leetcode/max_water"
)

func main() {
	s := "A."
	if max_water.IsPalindrome(s) {
		println("yes")
	} else {
		println("no")
	}
}
