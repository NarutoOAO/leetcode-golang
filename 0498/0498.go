package _498

func findDiagonalOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	x, y, d, row, col, i := 0, 0, 0, len(matrix), len(matrix[0]), 0
	direction := [][]int{{-1, 1}, {1, -1}}
	sum := row * col
	res := make([]int, sum)
	for i < sum {
		for x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
			res[i] = matrix[x][y]
			i++
			x += direction[d][0]
			y += direction[d][1]
		}
		d = (d + 1) % 2
		if y == len(matrix[0]) {
			x += 2
			y--
		}
		if x == len(matrix) {
			y += 2
			x--
		}
		if x < 0 {
			x = 0
		}
		if y < 0 {
			y = 0
		}
	}
	return res
}
