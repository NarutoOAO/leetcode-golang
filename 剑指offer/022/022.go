package _22

func validateStackSequences(pushed []int, popped []int) bool {
	s, i := make([]int, 0), 0
	for _, v := range pushed {
		s = append(s, v)
		for len(s) > 0 && popped[i] == s[len(s)-1] {
			s = s[:len(s)-1]
			i++
		}
	}
	if len(s) > 0 {
		return false
	}
	return true
}
