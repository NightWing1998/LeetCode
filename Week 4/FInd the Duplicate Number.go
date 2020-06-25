func findDuplicate(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		index := nums[i] % n
		nums[index] += n
	}
	maxNum, maxIndex := 0, -1
	for i := 0; i < n; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIndex = i
		}
		nums[i] = nums[i] % n
	}
	return maxIndex
}