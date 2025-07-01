package sliding_window

import "slices"

// 给你一个整数数组 nums 和一个 正整数 k 。
//
// 请你统计有多少满足 「 nums 中的 最大 元素」至少出现 k 次的子数组，并返回满足这一条件的子数组的数目。
//
// 子数组是数组中的一个连续元素序列。

func CountSubarrays(nums []int, k int) int64 {
	m := slices.Max(nums)
	l := 0
	cnt := 0
	var ans int64
	ans = 0
	for _, x := range nums {
		if x == m {
			cnt++
			for cnt == k {
				if nums[l] == m {
					cnt--
				}
				l++
			}
		}
		ans = ans + int64(l)
		println(l)
	}
	return ans
}
func wrong_countSubarrays(nums []int, k int) int64 {
	m := slices.Max(nums)
	l := 0
	cnt := 0
	var ans int64
	ans = 0
	for _, x := range nums {
		if x == m {
			cnt++
			for cnt == k {
				if nums[l] == m {
					cnt--
					ans = ans + int64(l+1)

				}
				l++
			}
		}
	}
	return ans
}

//*体会一下错误答案的错因
