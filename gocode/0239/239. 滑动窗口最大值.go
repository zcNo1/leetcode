package l0239

func maxSlidingWindow(nums []int, k int) []int {
	return maxSlid(nums, k)[k-1:]
}

func maxSlid(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	ret := make([]int, len(nums))
	half := k / 2
	var left, right []int

	if half == k-half {
		left = maxSlid(nums, half)
		right = left
	} else {
		left = maxSlid(nums, half)
		right = maxSlid(nums, k-half)
	}

	for i := k - 1; i < len(nums); i++ {
		ret[i] = getMax(left[i], right[i-half])
	}

	return ret
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
