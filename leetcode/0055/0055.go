package _055

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	maxjump := 0
	for i, v := range nums {
		if i > maxjump {
			return false
		}
		if maxjump < i+v {
			maxjump = i + v
		}
	}
	return true
}
