package main

func findChampion(grid [][]int) int {
	var i, j int
	for i = 1; i < len(grid); i++ {
		for j = 0; j < i; j++ {
			if grid[i][j] != 1 {
				break
			}
			if j == i-1 {
				return i
			}
		}
	}
	return 0
}

func main() {

}
