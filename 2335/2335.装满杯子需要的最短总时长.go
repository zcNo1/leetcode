package _335

import "sort"

func fillCups(amount []int) int {
	flag := true
	ret := 0
	for flag {
		sort.Ints(amount)
		if amount[2] > 0 && amount[1] > 0 {
			ret++
			amount[2]--
			amount[1]--
			continue
		}
		ret += amount[2]
		flag = false
	}

	return ret
}
