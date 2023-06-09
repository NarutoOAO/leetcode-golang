package _033

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	low, right := 0, len(nums)-1
	for low <= right {
		mid := low + (right-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[low] {
			if target < nums[mid] && target >= nums[low] {
				right = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[right] {
			if target > nums[mid] && target <= nums[right] {
				low = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			} else if nums[mid] == nums[right] {
				right--
			}
		}
	}
	return -1
}
