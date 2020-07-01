// Link - https://leetcode.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	if target > nums[end] {
		return end + 1
	}
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}
	return start
}

// Time:
// 	Usage : 4ms
// 	Beats : 88.73%
// Memory:
// 	Usage : 3.1MB
// 	Beats : 10.53%