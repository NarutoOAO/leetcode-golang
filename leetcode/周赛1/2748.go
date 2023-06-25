package 周赛1

import "strconv"

func countBeautifulPairs(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		a := strconv.Itoa(nums[i])
		x := int(a[0] - '0')
		for j := i + 1; j < len(nums); j++ {
			b := strconv.Itoa(nums[j])
			y := int(b[len(b)-1] - '0')
			if gcd(x, y) == 1 {
				res++
			}
		}
	}
	return res
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
