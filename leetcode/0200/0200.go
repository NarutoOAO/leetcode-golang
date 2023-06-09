package _200

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(&grid, i, j)
			}
		}
	}
	return ans
}

func dfs(grid *[][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) {
		return
	}
	if (*grid)[i][j] == '0' {
		return
	}
	(*grid)[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)

}
