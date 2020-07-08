// Link - https://leetcode.com/problems/island-perimeter

func islandPerimeter(grid [][]int) int {
	n := len(grid)
	if n <= 0 {
		return 0
	}
	m := len(grid[n-1])
	if m <= 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				sum := 0
				calculatePerimeter(grid, n, m, i, j, &sum)
				return sum
			}
		}
	}
	return 0
}

func calculatePerimeter(grid [][]int, n, m, i, j int, sum *int) {
	if i >= n || i < 0 || j >= m || j < 0 || grid[i][j] != 1 {
		return
	}

	// fmt.Println(i,j,*sum);

	*sum += 4       // considering square has no neighbours
	grid[i][j] = -1 // indicates node is visited

	// iterate for neighbours
	if i+1 < n && grid[i+1][j] != 0 {
		*sum -= 1 // if neighbour then decrease 1 from perimeter
		if grid[i+1][j] == 1 {
			calculatePerimeter(grid, n, m, i+1, j, sum) // iterate if square not visited
		}
	}

	if i-1 >= 0 && grid[i-1][j] != 0 {
		*sum -= 1
		if grid[i-1][j] == 1 {
			calculatePerimeter(grid, n, m, i-1, j, sum)
		}
	}

	if j+1 < m && grid[i][j+1] != 0 {
		*sum -= 1
		if grid[i][j+1] == 1 {
			calculatePerimeter(grid, n, m, i, j+1, sum)
		}
	}

	if j-1 >= 0 && grid[i][j-1] != 0 {
		*sum -= 1
		if grid[i][j-1] == 1 {
			calculatePerimeter(grid, n, m, i, j-1, sum)
		}
	}

}

// Time - 148ms (beats - 5.59%)
// Memory - 7.3MB (beats - 8.69%)