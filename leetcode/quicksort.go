package leetcode

import "math/rand"

// 快选
func selectKthLargest(nums []int, l, r, k int) int {
	if l >= r {
		return nums[l]
	}
	q := partition(nums, l, r)
	if k == q {
		return nums[q]
	} else if k > q {
		return selectKthLargest(nums, q+1, r, k)
	} else {
		return selectKthLargest(nums, 0, q-1, k)
	}
}

// 快排
func quicksort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quicksort(nums, p+1, r)
	quicksort(nums, 0, p-1)
}

func partition(nums []int, l, r int) int {
	k := l + rand.Intn(r-l+1)
	nums[l], nums[k] = nums[k], nums[l]
	i, j := l, r
	for i < j {
		for i < j && nums[j] >= nums[l] {
			j--
		}
		for i < j && nums[i] <= nums[l] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[l] = nums[l], nums[i]
	return i
}
