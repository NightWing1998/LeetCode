// Link - https://leetcode.com/problems/h-index-ii/

func hIndex(citations []int) int {
	n := len(citations)
	h, start, end := 0, 0, n-1
	for start <= end {
		mid := (start + end) / 2
		if citations[mid] == n-mid {
			h = n - mid
			break
		} else if citations[mid] > n-mid {
			h = n - mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return h
}

// Time:
// 	Usage : 16ms
// 	Beats : 85.37%
// Memory:
// 	Usage : 6.3MB
// 	Beats : 41.18%