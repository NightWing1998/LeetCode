// Link - https://leetcode.com/problems/permutation-sequence/

import (
	"strings"
)

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	var table = make([]int, 9)
	table[0] = 1
	nums[0] = 1
	for i := 1; i < n; i++ {
		nums[i] = i + 1
		table[i] = i * table[i-1]
	}
	k--
	var ans strings.Builder
	for i := n - 1; i >= 0; i-- {
		j := k / table[i]
		ans.WriteByte(byte(nums[j]) + '0')
		k -= j * table[i]

		nums = append(nums[:j], nums[j+1:]...)
	}
	return ans.String()
}

// Time:
// 	Usage : 0ms
// 	Beats : 100%
// Memory:
// 	Usage : 2MB
// 	Beats : 82.93%