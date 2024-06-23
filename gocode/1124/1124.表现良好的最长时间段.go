package l1224

func longestWPI(hours []int) int {
	tiredDays := make([]int, len(hours)+1)
	tiredDay := 0
	min := 0
	for i, hour := range hours {
		today := i + 1
		if hour > 8 {
			tiredDays[today] = tiredDays[i] + 1
			tiredDay++
		} else {
			tiredDays[today] = tiredDays[i] - 1
			if min > tiredDays[today] {
				min = tiredDays[today]
			}
		}
	}

	if tiredDay == 0 {
		return 0
	}

	var l, r, max int
	for i := 0; i >= min; i-- {
		r = 0
		for j := len(tiredDays) - 1; j > 0; j-- {
			if tiredDays[j] > i {
				r = j
				break
			}
		}

		for l = 0; l < r; l++ {
			if tiredDays[l] == i {
				max = getMax(max, r-l)
			}
		}
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
