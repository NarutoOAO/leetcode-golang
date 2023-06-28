package _040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	tmp, res := make([]int, 0), make([][]int, 0)
	sort.Ints(candidates)
	find(candidates, target, tmp, 0, &res)
	return res
}

func find(candidates []int, target int, tmp []int, index int, res *[][]int) {
	if target == 0 {
		b := make([]int, len(tmp))
		copy(b, tmp)
		*res = append(*res, b)
	}
	for i := index; i < len(candidates); i++ {
		//本次不取，下次取
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		if target-candidates[i] >= 0 {
			tmp = append(tmp, candidates[i])
			find(candidates, target-candidates[i], tmp, i+1, res)
			tmp = tmp[:len(tmp)-1]
		}
	}
}
