package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumImbalanceNumbers([]int{3, 1, 4}))

}
func sumImbalanceNumbers(nums []int) int {
	res, cnt, n := 0, 0, len(nums)
	for i, v1 := range nums {
		vis := make(map[int]int)
		vis[v1] = 1
		cnt = 0
		for j := i + 1; j < n; j++ {
			v2 := nums[j]
			if vis[v2] == 0 {
				cnt += 1 - vis[v2-1] - vis[v2+1]
				vis[v2] = 1
			}
			res += cnt
		}
	}
	return res
}
