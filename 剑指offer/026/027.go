package _26

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur1(root.Left, root.Right)
}

func recur1(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if (l != nil && r == nil) || (l == nil && r != nil) || (l.Val != r.Val) {
		return false
	}
	return recur(l.Left, r.Right) && recur(l.Right, r.Left)
}
