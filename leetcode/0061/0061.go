package _061

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	k = k % count
	if k == 0 {
		return head
	}
	head1 := &ListNode{Val: 0, Next: head}
	cur = head1
	for i := 0; i < count-k; i++ {
		cur = cur.Next
	}
	tmp := head1.Next
	head1.Next = cur.Next
	cur.Next = nil
	cur = head1
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = tmp
	return head1.Next
}
