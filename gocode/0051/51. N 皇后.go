package l0051

const (
	OK    = -1
	QUEEN = -2
)

func solveNQueens(n int) [][]string {
	record := make([][]int, n)
	for i := 0; i < len(record); i++ {
		record[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			record[i][j] = OK
		}
	}

	var rets [][]string

	var dfs func(int, int)
	dfs = func(startI, startJ int) {
		if record[startI][startJ] != OK {
			return
		}
		if startI == n-1 {
			record[startI][startJ] = QUEEN
			rets = append(rets, output(record))
			record[startI][startJ] = OK
			return
		}

		changeColor(record, startI, startJ, n, OK, startI)
		record[startI][startJ] = QUEEN
		//fmt.Println(startI, startJ, "change: ", record)
		for j := 0; j < n; j++ {
			dfs(startI+1, j)
		}
		changeColor(record, startI, startJ, n, startI, OK)
		record[startI][startJ] = OK
		//fmt.Println(startI, startJ, "rollback: ", record)
	}

	for j := 0; j < n; j++ {
		dfs(0, j)
	}

	return rets
}

func output(record [][]int) []string {
	ret := make([]string, len(record))
	for i := 0; i < len(record); i++ {
		str := ""
		for j := 0; j < len(record[0]); j++ {
			if record[i][j] == QUEEN {
				str += "Q"
			} else {
				str += "."
			}
		}
		ret[i] = str
	}

	return ret
}

func changeColor(record [][]int, startI, startJ, n, oriColor, dstColor int) {
	for i := 0; i < n; i++ {
		if record[i][startJ] == oriColor {
			record[i][startJ] = dstColor
		}
	}

	for i, j := startI, startJ; i < n && j >= 0; i, j = i+1, j-1 {
		if record[i][j] == oriColor {
			record[i][j] = dstColor
		}
	}

	for i, j := startI, startJ; i < n && j < n; i, j = i+1, j+1 {
		if record[i][j] == oriColor {
			record[i][j] = dstColor
		}
	}
}
