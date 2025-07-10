package binary_search

//

func binary3(nums []int) int {
	l, r := 0, len(nums)-1
	mid := l + (r-l)/2
	for l <= r {
		if nums[mid] < len(nums)-mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = l + (r-l)/2
	}
	return len(nums) - l
}
func hIndex(citations []int) int {
	var ans int
	ans = binary3(citations)
	return ans
}
