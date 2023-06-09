package _165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	n1 := strings.Split(version1, ".")
	n2 := strings.Split(version2, ".")
	for i := 0; i < len(n1) || i < len(n2); i++ {
		x1, x2 := 0, 0
		if i < len(n1) {
			x1, _ = strconv.Atoi(n1[i])
		}
		if i < len(n2) {
			x2, _ = strconv.Atoi(n2[i])
		}
		if x1 > x2 {
			return 1
		} else if x1 < x2 {
			return -1
		}
	}
	return 0
}
