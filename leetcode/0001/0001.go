package _001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		j, ok := m[target-nums[i]]
		if ok {
			res = append(res, j)
			res = append(res, i)
			return res
		} else {
			m[nums[i]] = i
		}
	}
	return res
}
