package main

func findOriginalArray(changed []int) []int {
	record := [100001]int{}
	for _, i2 := range changed {
		record[i2]++
	}

	var ret []int
	if record[0]%2 != 0 {
		return ret
	}

	for j := 0; j < record[0]/2; j++ {
		ret = append(ret, 0)
	}

	for i := len(record) - 1; i > 0; i-- {
		if record[i] == 0 {
			continue
		}

		if i%2 != 0 {
			return []int{}
		}

		if record[i] > record[i/2] {
			return []int{}
		}

		for j := 0; j < record[i]; j++ {
			ret = append(ret, i/2)
		}
		record[i/2] -= record[i]
		record[i] = 0
	}

	return ret
}

func main() {

}
