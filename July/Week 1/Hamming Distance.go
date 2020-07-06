// Link - https://leetcode.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	x, count := x^y, 0
	for ; x > 0; x >>= 1 {
		if x&1 == 1 {
			count++
		}
	}
	return count
}

// Time - 0ms (beats -)
// Memory - 2MB (beats 22.34%)