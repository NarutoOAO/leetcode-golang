package _543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	res := l + r
	return max2(res, diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func max2(x, y, z int) int {
	if x >= y && x >= z {
		return x
	}
	if y >= z && y >= x {
		return y
	}
	return z
}
