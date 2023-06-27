package leetcode

import "math/rand"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	quicksort(intervals, 0, len(intervals)-1)
	res := make([][]int, 0)
	res = append(res, intervals[0])
	cur := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > res[cur][1] {
			res = append(res, intervals[i])
			cur++
		} else {
			res[cur][1] = max(intervals[i][1], res[cur][1])
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func quicksort(intervals [][]int, l, r int) {
	if l >= r {
		return
	}
	p := partition(intervals, l, r)
	quicksort(intervals, l, p-1)
	quicksort(intervals, p+1, r)
}

func partition(intervals [][]int, l, r int) int {
	k := rand.Intn(r-l+1) + l
	intervals[l], intervals[k] = intervals[k], intervals[l]
	i, j := l, r
	for i < j {
		for i < j && intervals[j][0] > intervals[l][0] {
			j--
		}
		for i < j && intervals[i][0] <= intervals[l][0] {
			i++
		}
		intervals[i], intervals[j] = intervals[j], intervals[i]
	}
	intervals[l], intervals[i] = intervals[i], intervals[l]
	return i
}
