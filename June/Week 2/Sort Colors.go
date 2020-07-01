// Link - https://leetcode.com/problems/sort-colors/

func sortColors(nums []int) {
	var count = make(map[int]int)
	for _, color := range nums {
		_, exists := count[color]
		if exists {
			count[color]++
		} else {
			count[color] = 1
		}
	}
	redCount, redExists := count[0]
	whiteCount, whiteExists := count[1]
	blueCount, blueExists := count[2]
	start := 0
	for ; redExists && start < redCount && start < len(nums); start++ {
		nums[start] = 0
	}
	for ; whiteExists && start < redCount+whiteCount && start < len(nums); start++ {
		nums[start] = 1
	}
	for ; blueExists && start < redCount+whiteCount+blueCount && start < len(nums); start++ {
		nums[start] = 2
	}
}

// Time:
// 	Usage : 0
// 	Beats : 100%
// Memory:
// 	Usage : 2MB
// 	Beats : 98.5%