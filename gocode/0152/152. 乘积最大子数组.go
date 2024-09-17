package l0152

func maxProduct(nums []int) int {
	prez := 1
	pref := 1
	nextz := 1
	nextf := 1

	ret := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nextz = prez * nums[i]
			nextf = 1
			if pref < 0 {
				nextf = pref * nums[i]
			}
			ret = getMax(ret, nextz)
		} else if nums[i] < 0 {
			nextz = 1
			if pref < 0 {
				nextz = pref * nums[i]
				ret = getMax(ret, nextz)
			}
			nextf = prez * nums[i]
		} else {
			nextz = 1
			nextf = 1
			ret = getMax(ret, 0)
		}
		prez = nextz
		pref = nextf
	}

	return ret
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
