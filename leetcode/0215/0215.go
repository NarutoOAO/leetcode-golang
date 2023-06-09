package _215

import "math/rand"

func findKthLargest(nums []int, k int) int {
	return selectKthLargest(nums, 0, len(nums)-1, len(nums)-k)
}

func selectKthLargest(nums []int, l, r, k int) int {
	if l >= r {
		return nums[l]
	}
	q := partion(nums, l, r)
	if k == q {
		return nums[q]
	} else if k > q {
		return selectKthLargest(nums, q+1, r, k)
	} else {
		return selectKthLargest(nums, 0, q-1, k)
	}
}

func partion(nums []int, l, r int) int {
	k := l + rand.Intn(r-l+1)
	nums[k], nums[r] = nums[r], nums[k]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
