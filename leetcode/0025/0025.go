package _025

import "leetcode-golang/leetcode"

func reverseKGroup(head *leetcode.ListNode, k int) *leetcode.ListNode {
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

func reverse(head, last *leetcode.ListNode) *leetcode.ListNode {
	pre := last
	for head != last {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}
