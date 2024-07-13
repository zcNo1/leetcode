package l0199

import . "gocode/pkg/tree"

type treeListNode struct {
	Val  *TreeNode
	Next *treeListNode
}

type Queue struct {
	fakeHead *treeListNode
	tail     *treeListNode
}

func NewQueue() *Queue {
	node := &treeListNode{}
	return &Queue{fakeHead: node, tail: node}
}

func (q *Queue) Empty() bool {
	return q.fakeHead == q.tail
}

func (q *Queue) Push(val *TreeNode) {
	q.tail.Next = &treeListNode{Val: val}
	q.tail = q.tail.Next
	return
}

func (q *Queue) Pop() *TreeNode {
	if q.Empty() {
		return nil
	}
	val := q.fakeHead.Next.Val
	q.fakeHead.Next = q.fakeHead.Next.Next
	if q.fakeHead.Next == nil {
		q.tail = q.fakeHead
	}
	return val
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int

	queue := NewQueue()
	queue.Push(root)
	cur := 1
	next := 0

	for !queue.Empty() {
		node := queue.Pop()
		if node.Left != nil {
			queue.Push(node.Left)
			next++
		}
		if node.Right != nil {
			queue.Push(node.Right)
			next++
		}
		cur--
		if cur == 0 {
			cur = next
			next = 0
			ret = append(ret, node.Val)
		}

	}
	return ret
}
