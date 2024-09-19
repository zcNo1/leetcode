package l0287

func findDuplicate(nums []int) int {
	sum := 0
	for i, i2 := range nums {
		sum += i2 - i
	}

	return sum
}
