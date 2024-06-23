package l1072

func maxEqualRowsAfterFlips(matrix [][]int) int {
	cnts := make(map[string]int)
	for i := 0; i < len(matrix); i++ {
		str := make([]byte, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j]^matrix[i][0] == 0 {
				str[j] = '0'
			} else {
				str[j] = '1'
			}
		}
		cnts[string(str)]++
	}

	res := 0
	for _, cnt := range cnts {
		if cnt > res {
			res = cnt
		}
	}

	return res
}
