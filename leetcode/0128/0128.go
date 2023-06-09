package _128

func longestConsecutive(nums []int) int {
	res := 0
	m := make(map[int]int)
	for _, num := range nums {
		if m[num] == 0 {
			left, right := 0, 0
			left = m[num-1]
			right = m[num+1]
			sum := left + right + 1
			m[num] = sum
			res = max(sum, res)
			m[num-left] = sum
			m[num+right] = sum
		} else {
			continue
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
