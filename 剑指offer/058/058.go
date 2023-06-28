package _58

func reverseLeftWords(s string, n int) string {
	s1 := s[:n]
	s2 := s[n:]
	return s2 + s1
}
