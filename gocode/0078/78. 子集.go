package l0078

func deepAppend(nums []int, val int) []int {
	ret := make([]int, len(nums))
	copy(ret, nums)
	return append(ret, val)
}

func subsets(nums []int) [][]int {
	if nums == nil {
		return nil
	}

	var rets, lastRets [][]int
	lastRets = append(lastRets, []int{})

	for i := 0; i < len(nums); i++ {
		// 遍历之前数字形成的数组
		for _, lastRet := range lastRets {
			rets = append(rets, deepAppend(lastRet, nums[i]))
		}
		lastRets = append(lastRets, rets...)
		rets = [][]int{}
	}

	return lastRets
}
