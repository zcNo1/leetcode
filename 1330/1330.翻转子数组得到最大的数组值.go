package l1330

import "math"

func maxValueAfterReverse(nums []int) int {
	ret := 0
	min := math.MaxInt
	max := 0
	for i := 1; i < len(nums); i++ {
		ret += getSubAbs(nums[i], nums[i-1])
		max = getMax(max, nums[i]+nums[i-1])
		min = getMin(min, nums[i]+nums[i-1])
	}

	return ret
}

func getSubAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
