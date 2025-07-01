package sliding_window

// 给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。
func LongestOnes(nums []int, k int) int {
	l, r := 0, 0
	n := len(nums)
	cnt := 0
	ans := 0
	for r < n {
		if nums[r] == 0 {
			cnt++
			for cnt > k {
				if nums[l] == 0 {
					cnt--
				}
				l++
			}
		}
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}
