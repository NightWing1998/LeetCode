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