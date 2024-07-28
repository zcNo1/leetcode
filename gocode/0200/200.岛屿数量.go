package l0200

func transfer(grid [][]byte, record [][]int, idx, idy int) {
	if idx < 0 || idx >= len(grid) || idy < 0 || idy >= len(grid[0]) {
		return
	}
	if record[idx][idy] == 1 {
		return
	}
	record[idx][idy] = 1

	if grid[idx][idy] == '0' {
		return
	}

	transfer(grid, record, idx+1, idy)
	transfer(grid, record, idx-1, idy)
	transfer(grid, record, idx, idy+1)
	transfer(grid, record, idx, idy-1)
}

func numIslands(grid [][]byte) (ret int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	record := make([][]int, len(grid))
	for i := 0; i < len(record); i++ {
		record[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if record[i][j] == 1 {
				continue
			}
			if grid[i][j] == '1' {
				transfer(grid, record, i, j)
				ret++
			}
			record[i][j] = 1
		}
	}

	return ret
}
