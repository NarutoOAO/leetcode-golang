package _143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	head1 := slow.Next
	slow.Next = nil
	head1 = reverse(head1)
	cur := head
	for head1 != nil {
		tmp1 := cur.Next
		tmp2 := head1.Next
		cur.Next = head1
		head1.Next = tmp1
		head1 = tmp2
		cur = tmp1
	}
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
