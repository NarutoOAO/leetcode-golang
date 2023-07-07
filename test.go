package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(pathSum(root, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	recur(root, target, []int{})
	return res
}

func recur(root *TreeNode, target int, tmp []int) {
	target -= root.Val
	if target < 0 {
		return
	}
	tmp = append(tmp, root.Val)
	if root.Left == nil && root.Right == nil && target == 0 {
		tmp1 := make([]int, len(tmp))
		copy(tmp1, tmp)
		res = append(res, tmp1)
		return
	}
	if root.Left != nil {
		recur(root.Left, target, tmp)
	}
	if root.Right != nil {
		recur(root.Right, target, tmp)
	}

	//tmp = tmp[:len(tmp)-1]
}
