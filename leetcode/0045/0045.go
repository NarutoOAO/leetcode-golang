package _045

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	needChoose := 0
	canReach := 0
	step := 0
	for i, v := range nums {
		if i+v > canReach {
			canReach = i + v
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			step++
			needChoose = canReach
		}
	}
	return step
}
