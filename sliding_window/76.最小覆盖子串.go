package sliding_window

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
// 注意：
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
func isCovered(cntS, cntT []int) bool {
	for i := 'A'; i <= 'Z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	return true
}

func minWindow(s, t string) string {
	var cntS, cntT [128]int
	for _, c := range t {
		cntT[c]++
	}

	ansLeft, ansRight := -1, len(s)
	left := 0
	for right, c := range s { // 移动子串右端点
		cntS[c]++                         // 右端点字母移入子串
		for isCovered(cntS[:], cntT[:]) { // 涵盖
			if right-left < ansRight-ansLeft { // 找到更短的子串
				ansLeft, ansRight = left, right // 记录此时的左右端点
			}
			cntS[s[left]]-- // 左端点字母移出子串
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

// 优化思路,使用less字符来表示剩下的字母，减少检查时间
func ans_minWindow(s, t string) string {
	cnt := [128]int{}
	less := 0
	for _, c := range t {
		if cnt[c] == 0 {
			less++ // 有 less 种字母的出现次数 < t 中的字母出现次数
		}
		cnt[c]++
	}

	ansLeft, ansRight := -1, len(s)
	left := 0
	for right, c := range s { // 移动子串右端点
		cnt[c]-- // 右端点字母移入子串
		if cnt[c] == 0 {
			// 原来窗口内 c 的出现次数比 t 的少，现在一样多
			less--
		}
		for less == 0 { // 涵盖：所有字母的出现次数都是 >=
			if right-left < ansRight-ansLeft { // 找到更短的子串
				ansLeft, ansRight = left, right // 记录此时的左右端点
			}
			x := s[left] // 左端点字母
			if cnt[x] == 0 {
				// x 移出窗口之前，检查出现次数，
				// 如果窗口内 x 的出现次数和 t 一样，
				// 那么 x 移出窗口后，窗口内 x 的出现次数比 t 的少
				less++
			}
			cnt[x]++ // 左端点字母移出子串
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

//作者：灵茶山艾府
//链接：https://leetcode.cn/problems/minimum-window-substring/solutions/2713911/liang-chong-fang-fa-cong-o52mn-dao-omnfu-3ezz/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
