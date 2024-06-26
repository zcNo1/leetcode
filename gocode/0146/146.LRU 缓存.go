package l0146

import "fmt"

type DoubleListNode struct {
	key  int
	val  int
	pre  *DoubleListNode
	next *DoubleListNode
}

type DoubleList struct {
	head *DoubleListNode
	tail *DoubleListNode
}

func NewDoubleList() *DoubleList {
	ret := &DoubleList{
		head: &DoubleListNode{},
		tail: &DoubleListNode{},
	}
	ret.head.next = ret.tail
	ret.tail.pre = ret.head

	return ret
}

func (d *DoubleList) AddTail(node *DoubleListNode) {
	node.pre = d.tail.pre
	node.pre.next = node

	d.tail.pre = node
	node.next = d.tail
}

func (d *DoubleList) DeleteNode(node *DoubleListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (d *DoubleList) DeleteHead() *DoubleListNode {
	if d.head.next == d.tail {
		return nil
	}
	head := d.head.next
	d.head.next = head.next
	head.next.pre = d.head

	return head
}

func printNode(head *DoubleListNode) {
	fmt.Println("--------------------")
	cnt := 0
	for head != nil {
		fmt.Printf("cnt: %d, key: %d, val: %d\n", cnt, head.key, head.val)
		cnt++
		head = head.next
	}
	fmt.Println("--------------------")
}

type LRUCache struct {
	where    map[int]*DoubleListNode
	sortList *DoubleList
	capacity int
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{
		where:    make(map[int]*DoubleListNode, capacity),
		sortList: NewDoubleList(),
		capacity: capacity,
	}

	return ret
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.where[key]; ok {
		this.sortList.DeleteNode(node)
		this.sortList.AddTail(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.where[key]; ok {
		node.val = value
		this.sortList.DeleteNode(node)
		this.sortList.AddTail(node)
		return
	}
	if len(this.where) == this.capacity {
		node := this.sortList.DeleteHead()
		delete(this.where, node.key)
		this.where[key] = node
		node.key = key
		node.val = value
		this.sortList.AddTail(node)
		return
	}
	node := &DoubleListNode{
		key: key,
		val: value,
	}
	this.sortList.AddTail(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
