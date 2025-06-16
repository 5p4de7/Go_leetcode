package two_sum_three_sum

import (
	"fmt"
	"sort"
)

//* 给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。

func cN2Cal(left, right int) int {
	n := right - left + 1

	ans := int(0.5 * float64(n) * float64(n-1))
	return ans
}
func TriangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	n := len(nums)
	for a := 0; a < n-2; a++ {
		b := a + 1
		c := n - 1
		left := b
		right := c
		cur1, cur2 := nums[a]+nums[b], nums[a]+nums[b]
		if cur1 > nums[c] {
			ans = ans + cN2Cal(b, c)
			fmt.Printf("ans is %v\n", b)
			fmt.Printf("结束\n")
			continue

		}
		for left < c && cur1 <= nums[c] {
			left++
			//println(left)
			//println("结束")
			cur1 = nums[a] + nums[left]
		}
		for b < right && cur2 <= nums[right] {
			right--
			//println(right)
			cur2 = nums[a] + nums[b]
		}
		if left == c || right == b {
			continue
		}
		fmt.Printf("left: %d,right: %d\n", left, right)
		fmt.Printf("b:%d,c:%d\n", b, c)
		var bonus int
		if left < right {
			bonus = cN2Cal(left, c) + cN2Cal(b, right) - cN2Cal(left, right)
		} else {
			bonus = cN2Cal(left, c) + cN2Cal(b, right) + cN2Cal(right, left)
		}

		fmt.Printf("bonus: %d\n", bonus)
		ans = ans + bonus
		fmt.Printf("ans is %v\n", b)
	}
	return ans
}
