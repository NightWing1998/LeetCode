import "math"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func dfs(dungeon [][]int, i, j, n, m int, table [][]int) int {
	if table[i][j] != -1 {
		return table[i][j]
	}
	if i == n-1 && j == m-1 {
		if dungeon[i][j] > 0 {
			return 1
		}
		return -dungeon[i][j] + 1
	}
	if i == n || j == m {
		return math.MaxInt32
	}
	table[i][j] = max(1, min(dfs(dungeon, i+1, j, n, m, table), dfs(dungeon, i, j+1, n, m, table))-dungeon[i][j])
	return table[i][j]
}

func calculateMinimumHP(dungeon [][]int) int {
	n := len(dungeon)
	if n == 0 {
		return 0
	}
	m := len(dungeon[n-1])
	if m == 0 {
		return 0
	}
	table := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		table[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			table[i][j] = -1
		}
	}
	return dfs(dungeon, 0, 0, n, m, table)
}