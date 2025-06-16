package two_sum_three_sum

import "sort"

// *给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 target
// *请你返回满足 0 <= i < j < n 且 nums[i] + nums[j] < target 的下标对 (i, j) 的数目。

func CountPairs(nums []int, target int) int {
	sort.Ints(nums)
	ans := 0
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			ans = ans + (right - left) //*关键在于找到所有的符合条件的对数
			left++
		} else {
			right--
		}
	}
	return ans
}
