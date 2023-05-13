package _124

func longestWPI(hours []int) int {
	var max, up, down, left, right int

	for right < len(hours) {
		if hours[right] > 8 {
			up++
		} else {
			down++
		}

		if up > down {
			max = getMax(max, right-left)
		} else {
			left++
		}
	}

}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
