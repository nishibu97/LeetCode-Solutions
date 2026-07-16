package main

// 回文を探す問題。
func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }

    start, end := 0,0
    for i := 0; i < len(s); i++ {
        // 奇数長: "aba"
		l1, r1 := expandAroundCenter(s, i, i)
		// 偶数長: "abba"
		l2, r2 := expandAroundCenter(s, i, i+1)

        if r1-l1 > end-start {
            start, end = l1, r1
        }
        if r2-l2 > end-start {
            start, end = l2, r2
        }
    }
    return s[start: end+1]
}

// 中央から左右にポインタを移動し探索
func expandAroundCenter(s string, left, right int) (int, int) {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return left + 1, right - 1
}