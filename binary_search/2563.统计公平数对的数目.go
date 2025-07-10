package binary_search

import "sort"

//给你一个下标从 0 开始、长度为 n 的整数数组 nums ，和两个整数 lower 和 upper ，返回 公平数对的数目 。
//
//如果 (i, j) 数对满足以下情况，则认为它是一个 公平数对 ：
//
//0 <= i < j < n，且
//lower <= nums[i] + nums[j] <= upper

func binary(nums []int, target int) int {
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

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var ans int64
	for j, x := range nums {
		//注意大小
		left := binary(nums[:j], lower-x)
		right := binary(nums[:j], upper-x+1)
		ans = ans + int64(right-left)
	}
	return ans

}
