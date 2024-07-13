package l0074

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	left, right := 0, len(matrix)*len(matrix[0])-1
	var middle, row, col int

	for left <= right {
		middle = left + (right-left)/2
		row = middle / len(matrix[0])
		col = middle % len(matrix[0])
		fmt.Println(row, col, left, right, middle)
		fmt.Println(matrix[row][col])
		fmt.Println("-----------------")
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return false
}
