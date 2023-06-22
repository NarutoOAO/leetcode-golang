package _031

func nextPermutation(nums []int) {
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(&nums, i+1, len(nums)-1)
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		i++
		j--
	}
}
