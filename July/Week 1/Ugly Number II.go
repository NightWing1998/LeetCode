// Link - https://leetcode.com/problems/ugly-number-ii/

func min(arr []int) int {
	n := len(arr)
	m := arr[0]
	for i := 1; i < n; i++ {
		if m > arr[i] {
			m = arr[i]
		}
	}
	return m
}

func nthUglyNumber(n int) int {
	factors := []int{2, 3, 5}
	k := 3
	starts := make([]int, k)
	nums := make([]int, 1)
	nums[0] = 1
	for i := 0; i < n-1; i++ {
		posNums := make([]int, k)
		for j := 0; j < k; j++ {
			posNums[j] = factors[j] * nums[starts[j]]
		}
		newNum := min(posNums)
		nums = append(nums, newNum)
		for j := 0; j < k; j++ {
			if posNums[j] == newNum {
				starts[j] += 1
			}
		}
	}
	return nums[len(nums)-1]
}

// Time - 12ms (beats 37.47%)
// Memory - 64.MB (beats -)