package l0046

import "fmt"

func permute(nums []int) [][]int {
	var rets, lastRets [][]int
	lastRets = append(lastRets, []int{nums[0]})

	for i := 1; i < len(nums); i++ {
		fmt.Println("-----------------")
		fmt.Println(nums[i])
		for _, lastRet := range lastRets {
			fmt.Println(lastRet)
			newRet := append(lastRet, nums[i])
			var dstRet []int
			copy(dstRet, newRet)
			rets = append(rets, dstRet)
			for j := len(lastRet) - 1; j >= 0; j-- {
				newRet[j], newRet[j+1] = newRet[j+1], newRet[j]
				var dstRet []int
				copy(dstRet, newRet)
				rets = append(rets, dstRet)
			}
			fmt.Println(rets)
		}
		lastRets = rets
		rets = [][]int{}
	}

	return lastRets
}
