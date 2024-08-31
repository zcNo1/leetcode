package l0322

import "math"

func coinChange(coins []int, amount int) int {
	record := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		min := math.MaxInt
		for j := 0; j < len(coins); j++ {
			if coins[j] > i {
				continue
			}
			if record[i-coins[j]] == -1 {
				continue
			}
			if record[i-coins[j]] < min {
				min = record[i-coins[j]]
			}
		}
		if min == math.MaxInt {
			record[i] = -1
		} else {
			record[i] = min + 1
		}
	}

	return record[amount]
}
