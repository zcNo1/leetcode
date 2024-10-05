package l0031

func nextPermutation(nums []int) {
	maxN := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			maxN = i
			break
		}
	}
	i := maxN
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	if maxN == 0 {
		return
	}

	for i := maxN; i < len(nums); i++ {
		if nums[i] > nums[maxN-1] {
			nums[i], nums[maxN-1] = nums[maxN-1], nums[i]
			break
		}
	}

	return
}
