package sliding_window

//给你一个下标从 0 开始的字符串 s ，这个字符串只包含 0 到 9 的数字字符。
//
//如果一个字符串 t 中至多有一对相邻字符是相等的，那么称这个字符串 t 是 半重复的 。例如，"0010" 、"002020" 、"0123" 、"2002" 和 "54944" 是半重复字符串，而 "00101022" （相邻的相同数字对是 00 和 22）和 "1101234883" （相邻的相同数字对是 11 和 88）不是半重复字符串。
//
//请你返回 s 中最长 半重复 子字符串 的长度。

func LongestSemiRepetitiveSubstring(s string) int {
	l := 0
	exi := 0
	ans := 1
	for r, _ := range s {
		if r == 0 {
			continue
		}
		if s[r] == s[r-1] {
			exi++

			for exi > 1 {
				l++
				if s[l] == s[l-1] {
					exi--
				}
			}
		}
		ans = max(ans, r-l+1)
		r++
	}
	return ans
}
