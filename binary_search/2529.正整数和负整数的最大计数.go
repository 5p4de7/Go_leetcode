package binary_search

import "sort"

// 给你一个按 非递减顺序 排列的数组 nums ，返回正整数数目和负整数数目中的最大值。
//
// 换句话讲，如果 nums 中正整数的数目是 pos ，而负整数的数目是 neg ，返回 pos 和 neg二者中的最大值。
// 注意：0 既不是正整数也不是负整数。
func lowerBound2(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := l + (r-l)/2
	for l <= r {
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = l + (r-l)/2
	}
	return l
}
func maximumCount(nums []int) int {
	n := lowerBound2(nums, 0)
	p := lowerBound2(nums, 1)
	if n > len(nums)-p {
		return n
	} else {
		return len(nums) - p
	}

}

func ans_maximumCount(nums []int) int {
	neg := sort.SearchInts(nums, 0)
	// 第一个 > 0 的位置，等价于第一个 >= 1 的位置
	pos := len(nums) - sort.SearchInts(nums, 1)
	return max(neg, pos)
}
