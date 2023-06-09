package _105

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderpos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inorderpos[inorder[i]] = i
	}
	return build(preorder, inorder, inorderpos, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(pre []int, in []int, inorderpos map[int]int, preleft int, preright int, inleft int, inright int) *TreeNode {
	if preleft > preright {
		return nil
	}
	root := &TreeNode{
		Val: pre[preleft],
	}
	inindex := inorderpos[pre[preleft]]
	root.Left = build(pre, in, inorderpos, preleft+1, preleft+inindex-inleft, inleft, inindex-1)
	root.Right = build(pre, in, inorderpos, preleft+inindex-inleft+1, preright, inindex+1, inright)
	return root
}
