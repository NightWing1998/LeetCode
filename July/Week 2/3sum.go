// Link - https://leetcode.com/problems/3sum/

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return true
		}
		return false
	})
	result := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			j, k := i+1, n-1
			for j < k {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					for j < k && nums[j] == nums[j+1] {
						j++
					}
					for j < k && nums[k] == nums[k-1] {
						k--
					}
					j++
					k--
				} else if sum > 0 {
					k--
				} else {
					j++
				}
			}
		}
	}
	return result
}

// Time - 28ms(beats 94.6%)
// Memory - 6.9MB(beats 67.31%)