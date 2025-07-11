package max_water

//给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//返回容器可以储存的最大水量。
//
//说明：你不能倾斜容器。

func MaxArea(height []int) int {
	ans := 0
	left := 0
	right := len(height) - 1
	cur := 0
	for right > left {
		cur = min(height[left], height[right]) * (right - left)
		println(cur)
		if cur > ans {
			ans = cur
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}
