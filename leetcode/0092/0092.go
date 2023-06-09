package _092

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	end := newHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	for i := 0; i < right; i++ {
		end = end.Next
	}
	head1 := pre.Next
	end1 := end.Next
	end.Next = nil
	pre.Next = reverse(pre.Next)
	head1.Next = end1
	return newHead.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
