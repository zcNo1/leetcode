package l0031

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			return
		}
	}

	n := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[n-i] = nums[n-i], nums[i]
	}
	return
}
