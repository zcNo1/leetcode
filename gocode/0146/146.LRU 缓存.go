package l0146

import "fmt"

type DoubleListNode struct {
	key  int
	val  int
	pre  *DoubleListNode
	next *DoubleListNode
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
	oldHead  *DoubleListNode
	NewTail  *DoubleListNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{
		where:    make(map[int]*DoubleListNode, capacity),
		oldHead:  &DoubleListNode{},
		NewTail:  &DoubleListNode{},
		capacity: capacity,
	}

	ret.oldHead.next = ret.NewTail
	ret.NewTail.pre = ret.oldHead

	return ret
}

func (l *LRUCache) addNewTail(node *DoubleListNode) {
	node.pre = l.NewTail.pre
	node.pre.next = node

	l.NewTail.pre = node
	node.next = l.NewTail
}

func (l *LRUCache) DeleteOldHead() {
	if l.oldHead.next == l.NewTail {
		return
	}
	l.oldHead.next = l.oldHead.next.next
	l.oldHead.next.pre = l.oldHead
}

func (l *LRUCache) moveNewTail(node *DoubleListNode) {
	if l.NewTail.pre == node {
		return
	}

	node.next.pre = node.pre
	node.pre.next = node.next
	l.addNewTail(node)
}

func (l *LRUCache) Get(key int) int {
	//defer printNode(l.oldHead.next)
	if node, ok := l.where[key]; ok {
		l.moveNewTail(node)
		return node.val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	//defer printNode(l.oldHead.next)
	if node, ok := l.where[key]; ok {
		node.val = value
		l.moveNewTail(node)
		return
	}
	if len(l.where) == l.capacity {
		node := l.oldHead.next

		delete(l.where, node.key)
		l.where[key] = node

		node.key = key
		node.val = value

		l.moveNewTail(node)
		return
	}
	node := &DoubleListNode{
		key: key,
		val: value,
	}
	l.where[key] = node
	l.addNewTail(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
