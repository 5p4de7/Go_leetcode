package binary_search

import (
	"slices"
	"sort"
)

// 给你两个正整数数组 spells 和 potions ，长度分别为 n 和 m ，其中 spells[i] 表示第 i 个咒语的能量强度，potions[j] 表示第 j 瓶药水的能量强度。
//
// 同时给你一个整数 success 。一个咒语和药水的能量强度 相乘 如果 大于等于 success ，那么它们视为一对 成功 的组合。
//
// 请你返回一个长度为 n 的整数数组 pairs，其中 pairs[i] 是能跟第 i 个咒语成功组合的 药水 数目。
func binary1(nums []int, target int64, chengzi int) int {
	l, r := 0, len(nums)-1
	mid := l + (r-l)/2
	for l <= r {
		if int64(chengzi*nums[mid]) < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = l + (r-l)/2
	}
	return l
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	m := len(potions)
	ans := []int{}
	var po int
	for _, x := range spells {
		po = binary1(potions, success, x)
		ans = append(ans, m-po)
	}
	return ans
}

func ans_successfulPairs(spells, potions []int, success int64) []int {
	slices.Sort(potions)
	for i, x := range spells {
		spells[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/x+1)
	}
	return spells
}
