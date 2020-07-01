// Link - https://leetcode.com/problems/single-number-ii/

func singleNumber(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		if _, exist := count[num]; exist {
			count[num]++
		} else {
			count[num] = 1
		}
	}
	for num, cnt := range count {
		if cnt == 1 {
			return num
		}
	}
	return -1
}

// Time:
//  Usage: 4ms
// 	Beats: 96.47%
// Memory:
//  Usage: 4.3MB
// 	Beats: 10.34%