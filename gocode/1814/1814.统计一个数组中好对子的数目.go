package l1814

import (
	"log"
)

func countNicePairs(nums []int) int {
	cntSub := make(map[int]int)
	for _, num := range nums {
		sub := getNumSub(num)
		//log.Println(num, sub)
		cntSub[sub]++
	}
	log.Println(nums, cntSub)
	ret := 0
	for _, cnt := range cntSub {
		n := cnt * (cnt - 1) / 2
		ret = (ret + n) % (1000000007)
	}
	return ret
}

func getNumSub(num int) int {
	return num - getRevereNum(num)
}

func getRevereNum(num int) int {
	//log.Println(num)
	newNum := 0
	for num != 0 {
		newNum = newNum*10 + num%10
		num /= 10
	}
	//log.Println(newNum)

	return newNum
}
