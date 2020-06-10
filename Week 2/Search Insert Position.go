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