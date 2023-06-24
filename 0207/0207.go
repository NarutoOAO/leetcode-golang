package _207

func canFinish(numCourses int, prerequisites [][]int) bool {
	in := make([]int, numCourses)
	out := make([][]int, numCourses)
	res := make([]int, 0)
	for _, v := range prerequisites {
		in[v[0]]++
		out[v[1]] = append(out[v[1]], v[0])
	}
	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			res = append(res, i)
		}
	}
	for i := 0; i < len(res); i++ {
		for _, v := range out[res[i]] {
			in[v]--
			if in[v] == 0 {
				res = append(res, v)
			}
		}
	}
	if len(res) == numCourses {
		return true
	}
	return false
}
