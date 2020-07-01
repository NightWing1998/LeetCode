// Link - https://leetcode.com/problems/find-the-duplicate-number/

func findDuplicate(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		index := nums[i] % n
		nums[index] += n // add len to the index pinted by num[i], this results in repeated num[i] to have max value
	}
	maxNum, maxIndex := 0, -1
	// find max value and index
	for i := 0; i < n; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIndex = i
		}
		nums[i] = nums[i] % n
	}
	return maxIndex
}

// Time:
// 	Usage : 4ms
//  Beats : 99.31%
// Memory:
//  Usage : 3.8MB
// 	Beats : 100%