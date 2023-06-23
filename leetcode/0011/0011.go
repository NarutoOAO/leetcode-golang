package _011

func maxArea(height []int) int {
	res, i, j, tmp := 0, 0, len(height)-1, 0
	for i < j {
		if height[i] < height[j] {
			tmp = height[i] * (j - i)
			if tmp > res {
				res = tmp
			}
			i++
		} else {
			tmp = height[j] * (j - i)
			if tmp > res {
				res = tmp
			}
			j--
		}
	}
	return res
}
