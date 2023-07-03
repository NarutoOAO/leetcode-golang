package main

import (
	"fmt"
)

func main() {
	fmt.Println(findRepeatNumber([]int{3, 1, 3}))

}
func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if m[num] > 0 {
			return num
		}
		m[num] = 1
	}
	return 0
}
