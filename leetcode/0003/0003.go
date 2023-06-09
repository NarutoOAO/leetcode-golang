package _003

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var a [300]int
	left, right, result := 0, -1, 0
	for left < len(s) {
		if right+1 < len(s) && a[s[right+1]-'a'] == 0 {
			a[s[right+1]-'a']++
			right++
		} else {
			a[s[left]-'a']--
			left++
		}
		if (right - left + 1) > result {
			result = right - left + 1
		}
	}
	return result
}
