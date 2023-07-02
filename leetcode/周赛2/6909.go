package 周赛2

func longestAlternatingSubarray(nums []int, threshold int) int {
	length, flag, tmp, i, j := 0, 0, 0, 0, 0
	for i < len(nums) {
		if nums[i]%2 != 0 || nums[i] > threshold {
			i++
			continue
		}
		flag = 1
		tmp = 1
		for j = i + 1; j < len(nums); j++ {
			if nums[j]%2 == flag && nums[j] <= threshold {
				flag = 1 - flag
				tmp++
				if tmp > length {
					length = tmp
				}
			} else {
				break
			}
		}
		if tmp > length {
			length = tmp
		}
		i = j
	}
	return length
}
