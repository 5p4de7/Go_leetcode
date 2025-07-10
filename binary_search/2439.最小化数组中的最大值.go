package binary_search

import "slices"

//给你一个下标从 0 开始的数组 nums ，它含有 n 个非负整数。
//
//每一步操作中，你需要：
//
//选择一个满足 1 <= i < n 的整数 i ，且 nums[i] > 0 。
//将 nums[i] 减 1 。
//将 nums[i - 1] 加 1 。
//你可以对数组执行 任意 次上述操作，请你返回可以得到的 nums 数组中 最大值 最小 为多少。

func panduan2(nums []int, limit int) bool {
	flag := len(nums) - 1
	for flag > 0 {
		if nums[flag] > limit {
			nums[flag-1] += nums[flag] - limit
		}
		flag--
	}
	if nums[0] > limit {
		return false
	} else {
		return true
	}
}

func minimizeArrayValue(nums []int) int {
	l, r := 0, slices.Max(nums)
	mid := l + (r-l)/2
	for l <= r {
		nc := make([]int, len(nums))
		_ = copy(nc, nums)
		if !panduan2(nc, mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = l + (r-l)/2
	}
	return l
}
