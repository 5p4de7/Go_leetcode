package max_water

import "strings"

//如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
//
//字母和数字都属于字母数字字符。
//
//给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
//

func IsPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	s = strings.ToLower(s)
	for left < right {
		for !(s[left] <= '9' && s[left] >= '0') && !(s[left] <= 'z' && s[left] >= 'a') && left < right {
			left++
		}
		for !(s[right] <= '9' && s[right] >= '0') && !(s[right] <= 'z' && s[right] >= 'a') && left < right {
			println(2)
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
