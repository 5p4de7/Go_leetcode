package binary_search

import "sort"

//珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 h 小时后回来。
//
//珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
//
//珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
//
//返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。

func binary4(piles []int, h int) int {
	l, r := 0, piles[len(piles)-1]
	mid := l + (r-l)/2
	pan := func(k int, piles []int) bool {
		n := len(piles)
		var total int
		if k == 0 {
			return false
		}
		for _, x := range piles {
			total += (x - 1) / k
		}
		if total+n <= h {
			return true
		} else {
			return false
		}
	}
	for l <= r {
		if !pan(mid, piles) {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = l + (r-l)/2
	}
	return l
}

func minEatingSpeed(piles []int, h int) int {
	sort.Ints(piles)
	ans := binary4(piles, h)
	return ans
}
