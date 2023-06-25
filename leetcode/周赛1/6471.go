package 周赛1

import "math/bits"

// bits.OnesCount 统计二进制中1的数量 先减去后面的num2 * i，再判断能不能拆成i个2的n次方，最少能拆成二进制中1的数量，最多能拆成当前数
func makeTheIntegerZero(num1 int, num2 int) int {
	for i := 0; i < 1000; i++ {
		x := num1 - i*num2
		if x < 0 {
			continue
		}
		if i >= bits.OnesCount(uint(x)) && i <= x {
			return i
		}
	}
	return -1
}
