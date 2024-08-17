package l0155

import "math"

type MyListNode struct {
	Val  int
	Next *MyListNode
}

type MinStack struct {
	minHead *MyListNode
	head    *MyListNode
}

func Constructor() MinStack {
	return MinStack{
		minHead: &MyListNode{
			Next: &MyListNode{
				Val: math.MaxInt,
			},
		},
		head: &MyListNode{},
	}
}

func (this *MinStack) Push(val int) {
	cur := &MyListNode{
		Val:  val,
		Next: this.head.Next,
	}
	this.head.Next = cur

	if val > this.minHead.Next.Val {
		val = this.minHead.Next.Val
	}
	curMin := &MyListNode{
		Val:  val,
		Next: this.minHead.Next,
	}
	this.minHead.Next = curMin
}

func (this *MinStack) Pop() {
	this.head.Next = this.head.Next.Next
	this.minHead.Next = this.minHead.Next.Next
}

func (this *MinStack) Top() int {
	return this.head.Next.Val
}

func (this *MinStack) GetMin() int {
	return this.minHead.Next.Val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
