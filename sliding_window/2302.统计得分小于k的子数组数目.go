package sliding_window

// 一个数组的 分数 定义为数组之和 乘以 数组的长度。
//
// 比方说，[1, 2, 3, 4, 5] 的分数为 (1 + 2 + 3 + 4 + 5) * 5 = 75 。
// 给你一个正整数数组 nums 和一个整数 k ，请你返回 nums 中分数 严格小于 k 的 非空整数子数组数目。
//
// 子数组 是数组中的一个连续元素序列。
func countSubarrays(nums []int, k int64) int64 {
	var ans int64
	var l int
	var score int64
	score = int64(nums[0])
	if score < k {
		ans++
	}
	for r, x := range nums {
		if r == 0 {
			continue
		}
		score = ((score / int64(r-l)) + int64(x)) * int64(r-l+1)
		for score >= k && l < r {
			score = ((score / int64(r-l+1)) - int64(nums[l])) * int64(r-l)
			l++
		}
		if score < k {
			ans = ans + int64(r-l+1)
		}
	}
	return ans
}
