package _798

import (
	"sort"
)

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	num := 1
	if coins[0] != 1 {
		return 1
	}
	for _, coin := range coins[1:] {
		if coin > num+1 {
			break
		}
		num += coin
	}

	return num + 1
}
