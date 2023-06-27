package _008

import "math"

func myAtoi(s string) int {
	res, flag, i, n := 0, 1, 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	if i < n {
		if s[i] == '+' {
			flag = 1
			i++
		} else if s[i] == '-' {
			i++
			flag = -1
		}
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		res *= 10
		res += int(s[i] - '0')
		if res*flag < math.MinInt32 {
			return math.MinInt32
		} else if res*flag > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return res * flag
}
