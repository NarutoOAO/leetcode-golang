package _046

func permute(nums []int) [][]int {
	used := make([]bool, len(nums))
	p := []int{}
	res := [][]int{}
	genearatePermutation(nums, 0, p, &res, &used)
	return res
}

func genearatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		tmp := make([]int, len(p))
		copy(tmp, p)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, nums[i])
			genearatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}
