package same_direct_double

//给你一个整数数组 nums 和一个整数 k 。
//
//一个元素 x 在数组中的 频率 指的是它在数组中的出现次数。
//
//如果一个数组中所有元素的频率都 小于等于 k ，那么我们称这个数组是 好 数组。
//
//请你返回 nums 中 最长好 子数组的长度。
//
//子数组 指的是一个数组中一段连续非空的元素序列

func MaxSubarrayLength(nums []int, k int) int {
	zidian := make(map[int]int, len(nums))
	l, r := 0, 0
	ans := 0
	for r < len(nums) {
		if zidian[nums[r]] < k {
			zidian[nums[r]]++
			ans = max(r-l+1, ans)
			r++
		} else {
			zidian[nums[l]]--
			l++
		}
	}
	return ans
}
