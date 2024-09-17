package l0763

func partitionLabels(s string) []int {
	maxWhere := make(map[byte]int)
	for i, i2 := range []byte(s) {
		maxWhere[i2] = i
	}

	var rets []int
	var i, left, right int

	for i < len(s) {
		right = maxWhere[s[left]]
		//fmt.Println(i, right)
		for i <= right {
			if maxWhere[s[i]] > right {
				right = maxWhere[s[i]]
			}
			i++
		}
		rets = append(rets, right-left+1)
		//fmt.Println(i, right)
		left = right + 1
		i = left
	}

	return rets
}
