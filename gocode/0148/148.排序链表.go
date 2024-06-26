package l0148

import (
	. "gocode/pkg/list"
	"sort"
)

// TODO 归并排序
func sortList(head *ListNode) *ListNode {
	var sortSlice []*ListNode
	cur := head
	for cur != nil {
		sortSlice = append(sortSlice, cur)
		cur = cur.Next
	}

	sort.Slice(sortSlice, func(i, j int) bool {
		return sortSlice[i].Val < sortSlice[j].Val
	})

	newHead := &ListNode{}
	cur = newHead
	for _, node := range sortSlice {
		cur.Next = node
		cur = cur.Next
	}
	cur.Next = nil

	return newHead.Next
}
