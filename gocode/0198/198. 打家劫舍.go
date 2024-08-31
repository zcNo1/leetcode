package l0198

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	record := make([]int, len(nums)+1)

	ret := nums[0]
	record[0] = 0
	record[1] = nums[0]
	record[2] = nums[1]
	if nums[1] > nums[0] {
		ret = nums[1]
	}

	for i := 2; i < len(nums); i++ {
		if record[i-1] > record[i-2] {
			record[i+1] = nums[i] + record[i-1]
		} else {
			record[i+1] = nums[i] + record[i-2]
		}
		if record[i+1] > ret {
			ret = record[i+1]
		}
	}

	return ret
}
