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