package l0238

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	ret[0] = 1
	for i := 1; i < n; i++ {
		ret[i] = nums[i-1] * ret[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		ret[i] *= right
		right *= nums[i]
	}

	return ret
}

func productExceptSelf1(nums []int) []int {
	n := len(nums)
	left := make([]int, n+1)
	right := make([]int, n+1)
	left[0] = 1
	right[0] = 1

	for i, num := range nums {
		left[i+1] = left[i] * num
	}

	for i := n - 1; i >= 0; i-- {
		right[n-i] = right[n-i-1] * nums[i]
	}

	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = left[i] * right[n-1-i]
	}

	return ret
}
