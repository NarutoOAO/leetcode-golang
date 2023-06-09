package _139

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == false {
				continue
			}
			s1 := s[j:i]
			if m[s1] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
