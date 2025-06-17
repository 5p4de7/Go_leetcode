package two_sum_three_sum

import (
	"sort"
)

//* 给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。

func TriangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	n := len(nums)
	for a := 0; a < n-2; a++ {
		for b := a + 1; b < n-1; b++ {
			c := n - 1
			for b < c {
				if nums[a]+nums[b] <= nums[c] {
					c--
				} else {
					ans = ans + c - b
					break
				}
			}
		}
	}
	return ans
}
func Ans_triangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	n := len(nums)
	for c := 2; c < n; c++ {
		a, b := 0, c-1
		for a < b {
			if nums[a]+nums[b] > nums[c] {
				ans = ans + b - a
				b--
			} else {
				a++
			}
		}
	}
	return ans
}
