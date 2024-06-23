package main

import "fmt"

type chain struct {
	key  int
	val  int
	pre  *chain
	next *chain
}

type head struct {
	head *chain
}

func newChainHead() *head {
	return &head{
		head: &chain{
			key:  -1,
			val:  -1,
			pre:  nil,
			next: nil,
		},
	}
}

func (h *head) set(key, val int) {
	pre := h.head
	cur := pre.next
	for cur != nil {
		if cur.key == key {
			cur.val = val
			return
		}
		if cur.key > key {
			cur.pre = &chain{
				key:  key,
				val:  val,
				pre:  cur.pre,
				next: cur,
			}
			cur.pre.pre.next = cur.pre
			return
		}
		pre = cur
		cur = cur.next
	}
	pre.next = &chain{
		key:  key,
		val:  val,
		pre:  pre,
		next: nil,
	}
	fmt.Println("set new")
	return
}

func (h *head) del(key int) {
	for cur := h.head.next; cur != nil; cur = cur.next {
		if cur.key == key {
			cur.pre.next = cur.next
			if cur.next != nil {
				cur.next.pre = cur.pre
			}
			return
		}
	}
	return
}

func (h *head) get(key int) int {
	for cur := h.head.next; cur != nil; cur = cur.next {
		if cur.key == key {
			return cur.val
		}
	}
	return -1
}

func (h *head) list() {
	fmt.Printf("------------\nkey\tval\n")
	for cur := h.head.next; cur != nil; cur = cur.next {
		fmt.Printf("%d\t%d\n", cur.key, cur.val)
	}
	fmt.Printf("------------\n")
}

//type MyHashMap struct {
//	record *head
//}
//
//func Constructor() MyHashMap {
//	ret := MyHashMap{
//		record: newChainHead(),
//	}
//	return ret
//}
//
//func (this *MyHashMap) Put(key int, value int) {
//	this.record.set(key, value)
//	this.record.list()
//}
//
//func (this *MyHashMap) Get(key int) int {
//	return this.record.get(key)
//}
//
//func (this *MyHashMap) Remove(key int) {
//	this.record.del(key)
//	this.record.list()
//}

type MyHashMap struct {
	record []*head
}

func Constructor() MyHashMap {
	ret := MyHashMap{
		record: make([]*head, 100),
	}
	for i := range ret.record {
		ret.record[i] = newChainHead()
	}
	return ret
}

func (this *MyHashMap) Put(key int, value int) {
	whichHead := key % len(this.record)
	this.record[whichHead].set(key, value)
}

func (this *MyHashMap) Get(key int) int {
	whichHead := key % len(this.record)
	return this.record[whichHead].get(key)
}

func (this *MyHashMap) Remove(key int) {
	whichHead := key % len(this.record)
	this.record[whichHead].del(key)
	this.record[whichHead].list()
}
