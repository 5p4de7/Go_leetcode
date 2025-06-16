package two_sum_three_sum

import (
	"math"
	"sort"
)

// *给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// *返回这三个数的和。
//
// *假定每组输入只存在恰好一个解。

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var ans int
	ans = math.MaxInt
	dis := math.MaxInt
	n := len(nums)
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			cur := nums[left] + nums[right]
			if cur == target-nums[i] {
				ans = cur + nums[i]
				break
			} else if int(math.Abs(float64(target-nums[i]-cur))) < dis {
				ans = cur + nums[i]
				dis = int(math.Abs(float64(target - nums[i] - cur)))
			}
			if cur < target-nums[i] {
				left++
			} else {
				right--
			}
		}

	}
	return ans
}
