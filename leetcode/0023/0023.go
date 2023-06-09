package _023

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	l := mergeKLists(lists[:len(lists)/2])
	r := mergeKLists(lists[len(lists)/2:])
	return merge(l, r)
}

func merge(l *ListNode, r *ListNode) *ListNode {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.Val < r.Val {
		l.Next = merge(l.Next, r)
		return l
	}
	r.Next = merge(l, r.Next)
	return r
}
