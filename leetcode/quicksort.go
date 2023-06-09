package leetcode

import "math/rand"

// 快选
func selectKthLargest(nums []int, l, r, k int) int {
	if l >= r {
		return nums[l]
	}
	q := partion1(nums, l, r)
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
	p := partion1(nums, l, r)
	partion1(nums, p+1, r)
	partion1(nums, 0, p-1)
}

// [0:i] 的数一定小于nums[r]
func partion1(nums []int, l, r int) int {
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

func partion2(arr []int, l, r int) int {
	k := rand.Intn(r-l+1) + l
	arr[l], arr[k] = arr[k], arr[l]
	i := l
	j := r
	for i < j {
		for i < j && arr[j] >= arr[l] {
			j--
		}
		for i < j && arr[i] <= arr[l] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[l], arr[i] = arr[i], arr[l]
	return i
}
