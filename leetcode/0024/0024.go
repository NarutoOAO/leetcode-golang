package _024

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head1 := &ListNode{Val: 0}
	cur := head1
	for head != nil && head.Next != nil {
		tmp := head.Next.Next
		cur.Next = head.Next
		cur.Next.Next = head
		head.Next = tmp
		cur = cur.Next.Next
		head = tmp
	}
	return head1.Next
}
