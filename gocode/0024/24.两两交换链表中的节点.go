package l0024

import . "gocode/pkg/list"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	tmp := swapPairs(next.Next)
	next.Next = head
	head.Next = tmp

	return next
}
