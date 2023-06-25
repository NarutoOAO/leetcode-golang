package 周赛1

const MOD int = 1e9 + 7

// 将数组为1的下标加到tmp数组里,遍历tmp数组，乘每个下标的间距即可
func numberOfGoodSubarraySplits(nums []int) int {
	tmp := make([]int, 0)
	for i, v := range nums {
		if v == 1 {
			tmp = append(tmp, i)
		}
	}
	if len(tmp) == 0 {
		return 0
	}
	ans := 1
	for i := 1; i < len(tmp); i++ {
		ans = ans * (tmp[i] - tmp[i-1]) % MOD
	}
	return ans
}
