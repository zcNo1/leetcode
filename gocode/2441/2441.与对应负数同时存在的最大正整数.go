package l2441

import "sort"

func findMaxK(nums []int) int {
	sort.Ints(nums)
	r := len(nums) - 1
	l := 0
	for l < r && nums[l] < 0 && nums[r] > 0 {
		if nums[l]*-1 == nums[r] {
			return nums[r]
		}
		if nums[l]*-1 > nums[r] {
			l++
		} else {
			r--
		}
	}
	return -1
}
