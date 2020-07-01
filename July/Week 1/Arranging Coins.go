// Link - https://leetcode.com/problems/arranging-coins/

import "math"

func arrangeCoins(n int) int {
	// find i such that i*(i+1) <= 2*n --> quadratic equation = i^2 + i - 2n
	// Therefore, i > 0 && i <= Floor((-1 + ((1+8n)^1/2)))/2 --> return
	return int(math.Floor((-1 + math.Sqrt(float64(1+8*n))))) / 2
}

// Time: 4ms(beats 76.92%)
// Memory: 2.2MB(beats 100%)