package l0033

import "fmt"

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	middle := left + (right-left)/2

	for left <= right {
		middle = left + (right-left)/2
		fmt.Println(left, right, middle)

		if nums[middle] == target {
			return middle
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}

		if nums[left] < nums[right] {
			if nums[middle] < target {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if nums[middle] < target {
				if nums[middle] < nums[right] {
					if nums[left] > target {
						left = middle + 1
					} else {
						right = middle - 1
					}
				} else {
					left = middle + 1
				}
			} else {
				if nums[middle] < nums[right] {
					right = middle - 1
				} else {
					if nums[left] > target {
						left = middle + 1
					} else {
						right = middle - 1
					}
				}
			}
		}
	}

	return -1
}
