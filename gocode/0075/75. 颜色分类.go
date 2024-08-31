package l0075

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	const (
		red = iota
		white
		blue
	)

	for left < right {
		if nums[left] == red {
			left++
			continue
		}
		if nums[right] == blue {
			right++
			continue
		}
	}

}
