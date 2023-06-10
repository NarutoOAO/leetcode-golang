package _054

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	m := len(matrix)
	if m == 0 {
		return ans
	}
	n := len(matrix[0])
	if n == 0 {
		return ans
	}
	sum := m * n
	count := 0
	top, bottom, left, right := 0, m-1, 0, n-1
	for count < sum {
		i, j := left, right
		for i <= j && count < sum {
			ans = append(ans, matrix[top][i])
			i++
			count++
		}
		i, j = top+1, bottom
		for i <= j && count < sum {
			ans = append(ans, matrix[i][right])
			i++
			count++
		}
		i, j = right-1, left
		for i >= j && count < sum {
			ans = append(ans, matrix[bottom][i])
			i--
			count++
		}
		i, j = bottom-1, top+1
		for i >= j && count < sum {
			ans = append(ans, matrix[i][left])
			i--
			count++
		}
		top++
		bottom--
		left++
		right--
	}
	return ans
}
