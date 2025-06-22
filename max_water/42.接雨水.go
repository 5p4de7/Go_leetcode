package max_water

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
func trap(height []int) int {
	ans := 0
	left := 0
	right := len(height) - 1
	pre, suf := height[left], height[right]
	for left <= right {
		if pre < suf {
			if height[left] > pre {
				pre = height[left]
			}
			ans = ans + pre - height[left]
			left++
		} else {
			if height[right] > suf {
				suf = height[right]
			}
			ans = ans + suf - height[right]
			right--
		}
	}
	return ans

}
