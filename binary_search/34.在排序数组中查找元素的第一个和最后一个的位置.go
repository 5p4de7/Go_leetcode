package binary_search

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。
//
// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
func lowerBound1(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func searchRange(nums []int, target int) []int {
	ans1 := []int{-1, -1}
	if len(nums) == 0 {
		return ans1
	}
	l := lowerBound1(nums, target)
	r := lowerBound1(nums, target+1) - 1

	ans2 := []int{l, r}
	if r < l || l == len(nums) || r == -1 {
		return ans1
	}
	return ans2

}
