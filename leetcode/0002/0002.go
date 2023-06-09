package _002

import "leetcode-golang/leetcode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *leetcode.ListNode, l2 *leetcode.ListNode) *leetcode.ListNode {
	l := &leetcode.ListNode{Val: 0}
	cur := l
	carry := 0
	for l1 != nil || l2 != nil {
		tmp := carry
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		carry = tmp / 10
		cur.Next = &leetcode.ListNode{Val: tmp % 10}
		cur = cur.Next
	}
	if carry == 1 {
		cur.Next = &leetcode.ListNode{Val: 1}
	}
	return l.Next
}
