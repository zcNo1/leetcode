package l0034

func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	middle := (right-left)/2 + left

	start, end := -1, -1
	for left <= right {
		middle = (right-left)/2 + left
		if nums[middle] == target {
			start, end = middle, middle
			break
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	// not found
	if left > right {
		return []int{start, end}
	}

	// find start
	left, right = 0, start-1
	for left <= right {
		middle = (right-left)/2 + left
		if nums[middle] == target {
			start = middle
			right = start - 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	// find end
	left, right = end+1, len(nums)-1
	for left <= right {
		middle = (right-left)/2 + left
		if nums[middle] == target {
			end = middle
			left = end + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return []int{start, end}
}
