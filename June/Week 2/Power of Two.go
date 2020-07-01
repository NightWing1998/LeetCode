// Link - https://leetcode.com/problems/power-of-two/

import "math"

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	power := math.Log2(float64(n))
	if power == math.Floor(power) {
		return true
	}
	return false
}

// Time:
// 	Usage : 0ms
// 	Beats : 100%
// Memory:
// 	Usage : 2.2MB
// 	Beats : 19.48%