package 背包

import "math"

// 给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
// 计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//
// 你可以认为每种硬币的数量是无限的。
func coinChange2(coins []int, amount int) (ans int) {
	n := len(coins)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, amount+1) //*这里不用:=是因为前面以及初始化储存空间了
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(coinPosition, target int) (res int) {
		if coinPosition < 0 {
			if target == 0 {
				return 0
			}
			return math.MaxInt / 2
		}
		p := &memo[coinPosition][target]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if target < coins[coinPosition] {
			return dfs(coinPosition-1, target)
		}
		return min(dfs(coinPosition-1, target), dfs(coinPosition, target-coins[coinPosition])+1)

	}
	ans = dfs(n-1, amount)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

//*记忆化搜索
//*dfs
//*动态规划
//todo 加入贪婪算法
