package l0041

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; {
		if nums[i] > n || nums[i] <= 0 {
			nums[i] = 0
			i++
			continue
		}

		if nums[i]-1 == i {
			i++
			continue
		}
		if nums[nums[i]-1] == nums[i] {
			nums[i] = 0
			i++
			continue
		}

		nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]

	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}

	return n + 1
}
