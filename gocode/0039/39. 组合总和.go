package l0039

func combinationSum(candidates []int, target int) [][]int {
	var ret [][]int

	for i, candidate := range candidates {
		// if current num is equal than target, record and continue
		if candidate == target {
			ret = append(ret, []int{candidate})
			continue
		}
		nextTarget := target - candidate
		// smaller than 2, stop dfs
		if nextTarget < 2 {
			continue
		}
		newRets := combinationSum(candidates[i:], nextTarget)

		for _, newRet := range newRets {
			ret = append(ret, append([]int{candidate}, newRet...))
		}
	}
	return ret
}
