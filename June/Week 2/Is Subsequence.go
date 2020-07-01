// Link - https://leetcode.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
	ch := 0
	for i := 0; i < len(t) && ch < len(s); i++ {
		if t[i] == s[ch] {
			ch++
		}
	}
	if ch == len(s) {
		return true
	}
	return false
}

// Time:
// 	Usage : 0ms
// 	Beats : 100%
// Memory:
// 	Usage : 2.1MB
// 	Beats : 61.96%