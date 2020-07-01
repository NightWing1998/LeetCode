// Link - https://leetcode.com/problems/largest-divisible-subset/submissions/

import "sort"

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	div := make([]int, len(nums))
	prev := make([]int, len(nums))
	maxDivIndex := 0
	for i := 0; i < len(nums); i++ {
		prev[i] = -1
		div[i] = 1
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && div[i] < div[j]+1 {
				div[i] = div[j] + 1
				prev[i] = j
			}
		}
		if div[maxDivIndex] < div[i] {
			maxDivIndex = i
		}
	}

	res := make([]int, div[maxDivIndex])

	iter := maxDivIndex
	for i := div[maxDivIndex] - 1; iter >= 0; i, iter = i-1, prev[iter] {
		res[i] = nums[iter]
	}

	return res
}

// Time:
// 	Usage : 40ms
// 	Beats : 89.54%
// Memory:
// 	Usage : 2.8MB
// 	Beats : 82.35%