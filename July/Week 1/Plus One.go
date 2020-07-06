// Link - https://leetcode.com/problems/plus-one/submissions/

func plusOne(digits []int) []int {
	add := 1
	n := len(digits) - 1
	for ; add > 0 && n >= 0; n-- {
		digits[n] += add
		add = digits[n] / 10
		digits[n] %= 10
	}
	if add > 0 {
		return append([]int{add}, digits...)
	}
	return digits
}

// Time - 0ms(beats -)
// Memory - 2.1MB (beats 11.23%)