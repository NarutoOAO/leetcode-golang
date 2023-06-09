package _165

import (
	"strconv"
	"strings"
)

// 使用strings.Split方法，得到一个字符串数组，再使用strconv包中的Atoi方法可以得到每位上的数字，逐一比较，如果有差距则返回-1或1
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
