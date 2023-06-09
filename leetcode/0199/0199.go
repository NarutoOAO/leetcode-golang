package _199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) != 0 {
		qlen := len(q)
		for i := 0; i < qlen; i++ {
			node := q[0]
			if i == qlen-1 {
				ans = append(ans, node.Val)
			}
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return ans
}
