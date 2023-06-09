package _103

import "leetcode-golang/leetcode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *leetcode.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	tmp := []int{}
	q := []*leetcode.TreeNode{}
	q = append(q, root)
	reverse := 0
	curnum := 1
	nextlevelnum := 0
	for len(q) != 0 {
		if curnum > 0 {
			node := q[0]
			if node.Left != nil {
				q = append(q, node.Left)
				nextlevelnum++
			}
			if node.Right != nil {
				q = append(q, node.Right)
				nextlevelnum++
			}
			curnum--
			tmp = append(tmp, node.Val)
			q = q[1:]
		}
		if curnum == 0 {
			if reverse == 1 {
				for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
					tmp[i], tmp[j] = tmp[j], tmp[i]
				}
			}
			res = append(res, tmp)
			tmp = []int{}
			curnum = nextlevelnum
			nextlevelnum = 0
			if reverse == 1 {
				reverse = 0
			} else {
				reverse = 1
			}
		}
	}
	return res
}
