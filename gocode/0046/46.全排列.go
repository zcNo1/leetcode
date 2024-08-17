package l0046

func deepInsert(nums []int, val, index int) []int {
	if index == 0 {
		return append([]int{val}, nums...)
	}

	ret := make([]int, len(nums)+1)
	if index >= len(nums) {
		copy(ret, nums)
		return append(ret, val)
	}

	for i := 0; i < index; i++ {
		ret[i] = nums[i]
	}
	ret[index] = val
	for i := index + 1; i < len(ret); i++ {
		ret[i] = nums[i-1]
	}

	return ret
}

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	var rets, lastRets [][]int
	lastRets = append(lastRets, []int{nums[0]})

	for i := 1; i < len(nums); i++ {
		// 遍历之前数字形成的数组
		for _, lastRet := range lastRets {
			// 将新数字从第一位一直插入到最后一位，并生成新的数组
			for j := 0; j < len(lastRet); j++ {
				rets = append(rets, deepInsert(lastRet, nums[i], j))
			}
			// 最后一位，复用之前的数组，减少内存消耗
			rets = append(rets, append(lastRet, nums[i]))
		}
		lastRets = rets
		rets = [][]int{}
	}

	return lastRets
}
