package _025

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(head, last *ListNode) *ListNode {
	pre := last
	for head != last {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}
