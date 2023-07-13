package main

import "fmt"

func main() {
	fmt.Println(maxNonDecreasingLength([]int{11, 7, 7, 9}, []int{19, 19, 1, 7}))
}

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	preMax, sum, n, res := 0, 0, len(nums1), 0
	for i := 0; i < n; i++ {
		if min(nums1[i], nums2[i]) >= preMax {
			preMax = min(nums1[i], nums2[i])
			sum++
		} else if max(nums1[i], nums2[i]) >= preMax {
			preMax = max(nums1[i], nums2[i])
			sum++

		} else {
			sum = 1
			preMax = min(nums1[i], nums2[i])
		}
		res = max(res, sum)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
