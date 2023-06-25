package leetcode

// 判断二数是否互质，最小公约数为1
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
