package main

const kLen = 32

type MyHashSet struct {
	record []int
}

func Constructor() MyHashSet {
	mySet := MyHashSet{}
	mySet.record = make([]int, 40000)
	return mySet
}

func (this *MyHashSet) Add(key int) {
	numIdx := key / kLen
	where := key % kLen
	this.record[numIdx] |= 1 << where
}

func (this *MyHashSet) Remove(key int) {
	numIdx := key / kLen
	where := key % kLen
	this.record[numIdx] &= ^(1 << where)
}

func (this *MyHashSet) Contains(key int) bool {
	numIdx := key / kLen
	where := key % kLen
	return (this.record[numIdx]>>where)&1 == 1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func main() {

}
