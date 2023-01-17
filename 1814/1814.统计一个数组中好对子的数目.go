package main

import (
	"log"
)

func main() {
	log.Println(countNicePairs([]int{42, 11, 1, 97}))
	log.Println(countNicePairs([]int{13, 10, 35, 24, 76}))
	log.Println(countNicePairs([]int{442111244, 357716602, 131050131, 251794140, 4046404, 373969224, 1082801, 468525864, 384140537, 492968345}))
}

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
