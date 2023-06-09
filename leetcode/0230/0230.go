package _230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	ans   int
	count int
)

func kthSmallest(root *TreeNode, k int) int {
	count = 0
	inorder(root, k)
	return ans
}

func inorder(root *TreeNode, k int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left, k)
	}
	count++
	if count == k {
		ans = root.Val
	}
	if root.Right != nil {
		inorder(root.Right, k)
	}
}
