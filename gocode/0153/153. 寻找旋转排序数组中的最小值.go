package l0153

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	middle := left + (right-left)/2
	for left < right {
		if nums[left] <= nums[right] {
			return nums[left]
		}
		middle = left + (right-left)/2
		if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return right
}
