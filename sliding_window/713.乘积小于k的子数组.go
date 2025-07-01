package sliding_window

// 给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。

func NumSubarrayProductLessThanK(nums []int, k int) int {
	//n := len(nums)
	left := 0
	cur := 1
	ans := 0
	for right, _ := range nums {
		cur = cur * nums[right]
		for cur >= k {
			cur = cur / nums[left]
			left++
		}
		ans = ans + right - left + 1

	}

	return ans
}
