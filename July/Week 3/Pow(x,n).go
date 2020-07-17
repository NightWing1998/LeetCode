// Link - https://leetcode.com/problems/powx-n/

import "math"

func myPow(x float64, n int) float64 {
	// why re-invent the wheel?
	return math.Pow(x,float64(n));
}

// Time - 0ms(beats 100%)
// Memory - 2MB(beats 69.86%)
