package main

import "log"

func main() {
	log.Println(minMaxGame([]int{1, 3, 5, 2, 4, 8, 2, 2}))
	log.Println(minMaxGame([]int{3}))
}

func minMaxGame(nums []int) int {
	for len(nums) != 1 {
		newNums := make([]int, len(nums)/2)
		for i := 0; i < len(nums); i += 2 {
			if i%4 == 0 {
				newNums[i/2] = min(nums[i], nums[i+1])
				continue
			}
			newNums[i/2] = max(nums[i], nums[i+1])
		}
		nums = newNums
		//log.Println(nums)
	}

	return nums[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
