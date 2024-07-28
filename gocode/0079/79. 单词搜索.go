package l0079

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	wordByte := []byte(word)
	record := make([][]bool, len(board))
	for i := 0; i < len(record); i++ {
		record[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, wordByte, record, i, j) {
				return true
			}
		}
	}

	return false
}

var directions = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func dfs(board [][]byte, word []byte, record [][]bool, idx, idy int) bool {
	// has found
	if len(word) == 0 {
		return true
	}

	// over edge
	if idx < 0 || idx >= len(board) || idy < 0 || idy >= len(board[0]) {
		return false
	}

	// has dfs
	if record[idx][idy] == true {
		return false
	}

	// is not fit word
	if board[idx][idy] != word[0] {
		return false
	}

	// record
	record[idx][idy] = true
	word = word[1:]
	for _, direction := range directions {
		if dfs(board, word, record, idx+direction[0], idy+direction[1]) {
			return true
		}
	}

	// not fit,rollback
	record[idx][idy] = false
	return false
}
