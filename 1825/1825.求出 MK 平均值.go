package _1825

import (
	"sort"
)

type MKAverage struct {
	m    int
	k    int
	cnt  int
	nums []int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m:    m,
		k:    k,
		cnt:  0,
		nums: make([]int, m),
	}
}

func (this *MKAverage) AddElement(num int) {
	this.nums[this.cnt%this.m] = num
	this.cnt++
}

func (this *MKAverage) CalculateMKAverage() int {
	if this.cnt < this.m {
		return -1
	}

	newNums := make([]int, len(this.nums))
	copy(newNums, this.nums)
	sort.Ints(newNums)
	sum := 0
	for i := this.k; i < this.m-this.k; i++ {
		sum += newNums[i]
	}

	return sum / (this.m - 2*this.k)
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
