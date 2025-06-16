package two_sum_three_sum

import "sort"

//* 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//* 0 <= a, b, c, d < n
//* a、b、c 和 d 互不相同
//* nums[a] + nums[b] + nums[c] + nums[d] == target
//* 你可以按 任意顺序 返回答案

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	n := len(nums)
	for a := 0; a < n-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < n-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			c := b + 1
			d := n - 1
			for c < d {
				cur := nums[a] + nums[b] + nums[c] + nums[d]
				if cur < target {
					c++
				} else if cur > target {
					d--
				} else {
					ans = append(ans, []int{nums[a], nums[b], nums[c], nums[d]})
					for c++; c < d && nums[c] == nums[c-1]; c++ {
					}
				}

			}
		}
	}
	return ans
}
