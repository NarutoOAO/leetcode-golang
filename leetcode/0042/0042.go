package _042

func trap(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	maxleft, maxright := 0, 0
	for left < right {
		if height[left] <= height[right] {
			if height[left] > maxleft {
				maxleft = height[left]
			} else {
				res += maxleft - height[left]
			}
			left++
		} else {
			if height[right] > maxright {
				maxright = height[right]
			} else {
				res += maxright - height[right]
			}
			right--
		}
	}
	return res
}
