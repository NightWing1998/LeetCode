// Link - https://leetcode.com/problems/unique-binary-search-trees/

func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	solution := make([]int, n+1)
	solution[0], solution[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			solution[i] += solution[j] * solution[i-j-1]
		}
	}
	return solution[n]
}

// Time:
//  Usage: 0ms
// 	Beats: 100%
// Memory:
//  Usage: 2MB
// 	Beats: 28.43%