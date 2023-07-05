package _12

// 回溯
var words []byte

func exist(board [][]byte, word string) bool {
	words = []byte(word)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j, k int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != words[k] {
		return false
	}
	if k == len(words)-1 {
		return true
	}
	board[i][j] = 0
	ans := dfs(board, i+1, j, k+1) || dfs(board, i-1, j, k+1) || dfs(board, i, j+1, k+1) || dfs(board, i, j-1, k+1)
	board[i][j] = words[k]
	return ans
}
