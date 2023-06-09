package _109

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val, Left: nil, Right: nil}
	}
	pre, mid := midNode(head)
	if pre != nil {
		pre.Next = nil
	}
	if mid == head {
		head = nil
	}
	return &TreeNode{Val: mid.Val, Left: sortedListToBST(head), Right: sortedListToBST(mid.Next)}
}

func midNode(head *ListNode) (pre *ListNode, mid *ListNode) {
	if head == nil || head.Next == nil {
		return nil, head
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return pre, slow
}
