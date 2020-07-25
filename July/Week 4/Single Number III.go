// Link - https://leetcode.com/problems/single-number-iii

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	mask := (xor & (xor - 1)) ^ xor
	n1 := 0
	for _, num := range nums {
		if num&mask == 0 {
			n1 ^= num
		}
	}
	return []int{n1, xor ^ n1}
}

// Time - 4ms(beats 98.98%)
// Memory - 4.2MB(beats 100%)