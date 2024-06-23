package main

import (
	"log"
)

func main() {
	//log.Println(countDifferentSubsequenceGCDs([]int{6, 12}) == 3)
	//log.Println(countDifferentSubsequenceGCDs([]int{6, 12, 8}) == 5)
	//log.Println(countDifferentSubsequenceGCDs([]int{6, 12, 8, 14}) == 7)
	//log.Println(countDifferentSubsequenceGCDs([]int{6, 10, 3}) == 5)
	log.Println(countDifferentSubsequenceGCDs([]int{5, 15, 40, 5, 6}) == 7)
}

func countDifferentSubsequenceGCDs(nums []int) int {
	record := make(map[int]int)

	record[1] = 2
	for _, num := range nums {
		if _, ok := record[num]; ok {
			record[num] += 2
			continue
		}
		//max := 0
		record[0] = 0
		for i := 2; i*i <= num; i++ {
			if num%i != 0 {
				continue
			}
			record[num/i]++
			record[i]++
			//if _, ok := record[num/i]; ok {
			//	if num/i > max {
			//		record[num/i]++
			//		record[max]--
			//		max = num / i
			//	}
			//} else {
			//	record[num/i] = 1
			//}
			//if _, ok := record[i]; ok {
			//	if i > max {
			//		record[i]++
			//		record[max]--
			//		max = i
			//	}
			//} else {
			//	record[i] = 1
			//}
		}
		//record[num] += 2
		log.Printf("num: %d, array: %v", num, record)
	}
	log.Println(record)

	cnt := 0
	for _, r := range record {
		if r > 1 {
			//log.Println(n)
			cnt++
		}
	}

	return cnt
}
