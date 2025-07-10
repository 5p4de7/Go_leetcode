package binary_search

import "sort"

//给你一个正整数数组 price ，其中 price[i] 表示第 i 类糖果的价格，另给你一个正整数 k 。
//
//商店组合 k 类 不同 糖果打包成礼盒出售。礼盒的 甜蜜度 是礼盒中任意两种糖果 价格 绝对差的最小值。
//
//返回礼盒的 最大 甜蜜度。

func panduan3(box []int, k int, n int) bool {
	if len(box) < k {
		return false
	}
	cnt := 0
	pre := box[0]
	for _, p := range box[1:] {
		if p >= pre+n {
			cnt++
			pre = p
		}
	}
	return cnt >= k-1
}

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	l, r := 0, (price[len(price)-1]-price[0])/(k-1)+1
	mid := l + (r-l)/2
	for l <= r {
		if panduan3(price, k, mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
		mid = l + (r-l)/2
	}
	return l - 1
}
