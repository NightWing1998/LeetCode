func solve(board [][]byte) {
	n := len(board)
	if n == 0 {
		return
	}
	m := len(board[n-1])
	if m == 0 {
		return
	}
	for i := 0; i < n; i++ {
		dfs(board, i, 0, n, m)
		dfs(board, i, m-1, n, m)
	}
	for j := 0; j < m; j++ {
		dfs(board, 0, j, n, m)
		dfs(board, n-1, j, n, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j, rows, cols int) {
	if i < 0 || j < 0 || i >= rows || j >= cols || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'T'
	dfs(board, i-1, j, rows, cols)
	dfs(board, i, j-1, rows, cols)
	dfs(board, i+1, j, rows, cols)
	dfs(board, i, j+1, rows, cols)
}