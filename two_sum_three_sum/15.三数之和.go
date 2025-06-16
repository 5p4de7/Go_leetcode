package two_sum_three_sum

import (
	"slices"
	"sort"
)

// *给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//
// *注意：答案中不可以包含重复的三元组。

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i, lens := 0, len(nums); i < lens; i++ {
		if i > 0 {
			if nums[i] == nums[i-1] { //跳过重复数字
				println("跳过")
				continue
			}
		}
		left := i + 1
		right := lens - 1
		tem := 0
		for left < right {
			tem = nums[left] + nums[right]
			if tem == -nums[i] {
				if left-i == 1 {
					println("found 1")
					s := []int{nums[i], nums[left], nums[right]}
					ans = append(ans, s)
				} else if left-i > 1 && nums[left] != nums[left-1] {
					println("found 1")
					s := []int{nums[i], nums[left], nums[right]}
					ans = append(ans, s)
				} //*此处的判断显得很复杂，究其原因是在第一个的位置不能减一
				left++
			} else if tem > -nums[i] {
				right--
			} else {
				left++
			}
		}

	}
	return ans
}
func Ans_threeSum(nums []int) (ans [][]int) {
	slices.Sort(nums)
	n := len(nums)
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] { // 跳过重复数字
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 { // 优化一
			break
		}
		if x+nums[n-2]+nums[n-1] < 0 { // 优化二
			continue
		}
		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else { // 三数之和为 0
				ans = append(ans, []int{x, nums[j], nums[k]})
				//*值得注意的是在完成操作之后再判断后来的值，可以避开第一个的位置以免不能出现j-1的情况
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 跳过重复数字
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // *跳过重复数字这一步可以省略
			}
		}
	}
	return
}
