package l0279

func numSquares(n int) int {
	record := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := 100000
		for j := 1; j*j <= i; j++ {
			if record[i-j*j] < min {
				min = record[i-j*j]
			}
		}
		record[i] = min + 1
	}

	return record[n]
}
