package l0739

func dailyTemperatures(temperatures []int) []int {
	if temperatures == nil || len(temperatures) == 0 {
		return []int{}
	}

	ret := make([]int, len(temperatures))
	var cur []int

	for i := len(ret) - 1; i >= 0; i-- {
		for len(cur) > 0 && temperatures[i] > temperatures[cur[len(cur)-1]] {
			cur = cur[:len(cur)-1]
		}
		if len(cur) == 0 {
			cur = append(cur, i)
			ret[i] = 0
		} else {
			ret[i] = cur[len(cur)-1] - i
			cur = append(cur, i)
		}
	}

	return ret
}
