package _143

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	row := len(text1)
	col := len(text2)
	grid := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		grid[i] = make([]int, col+1)
	}
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if text1[i-1] == text2[j-1] {
				grid[i][j] = grid[i-1][j-1] + 1
			} else {
				grid[i][j] = max(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[row][col]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
