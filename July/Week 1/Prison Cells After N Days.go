// Link - https://leetcode.com/problems/prison-cells-after-n-days/

func prisonAfterNDays(cells []int, N int) []int {
	if N%14 != 0 {
		N = N % 14
	} else {
		N = 14
	}
	for ; N > 0; N-- {
		temp := make([]int, 8)
		temp[0], temp[7] = 0, 0
		for i := 1; i < 7; i++ {
			if cells[i-1] == cells[i+1] {
				temp[i] = 1
			} else {
				temp[i] = 0
			}
		}
		cells = temp
	}
	return cells
}

// Time: 0ms (beats 100%)
// Memory: 2.3MB (beats 90%)