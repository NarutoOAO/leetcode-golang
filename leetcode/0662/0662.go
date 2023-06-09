package _662

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	q := []*TreeNode{}
	res := 1
	q = append(q, &TreeNode{Val: 1, Left: root.Left, Right: root.Right})
	for len(q) > 0 {
		var left, right *int
		qlen := len(q)
		for i := 0; i < qlen; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				newVal := node.Val * 2
				if left == nil {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
				q = append(q, &TreeNode{Val: newVal, Left: node.Left.Left, Right: node.Left.Right})
			}
			if node.Right != nil {
				newVal := node.Val*2 + 1
				if left == nil {
					left = &newVal
				}
				if right == nil || *right < newVal {
					right = &newVal
				}
				q = append(q, &TreeNode{Val: newVal, Left: node.Right.Left, Right: node.Right.Right})
			}
			if right != nil || left != nil {
				if *right-*left+1 > res {
					res = *right - *left + 1
				}
			}
		}
	}
	return res
}
