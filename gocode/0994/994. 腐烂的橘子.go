package l0994

import "fmt"

const (
	noOrange   = 0
	goodOrange = 1
	badOrange  = 2
)

type index struct {
	x int
	y int
}

func isGoodOrange(grid [][]int, cur index) bool {
	if cur.x < 0 || cur.x >= len(grid) || cur.y < 0 || cur.y >= len(grid[0]) {
		return false
	}

	if grid[cur.x][cur.y] == goodOrange {
		return true
	}

	return false
}

func orangesRotting(grid [][]int) int {
	var badOranges []index
	cnt := 0

	for i, ints := range grid {
		for i2, i3 := range ints {
			if i3 == badOrange {
				badOranges = append(badOranges, index{i, i2})
				cnt++
			}
		}
	}

	nextCnt := 0
	minutes := 0
	for len(badOranges) > 0 {
		cur := badOranges[0]
		badOranges = badOranges[1:]
		cnt--

		next := index{cur.x + 1, cur.y}
		if isGoodOrange(grid, next) {
			grid[next.x][next.y] = badOrange
			badOranges = append(badOranges, next)
			nextCnt++
		}
		next = index{cur.x - 1, cur.y}
		if isGoodOrange(grid, next) {
			grid[next.x][next.y] = badOrange
			badOranges = append(badOranges, next)
			nextCnt++
		}
		next = index{cur.x, cur.y + 1}
		if isGoodOrange(grid, next) {
			grid[next.x][next.y] = badOrange
			badOranges = append(badOranges, next)
			nextCnt++
		}
		next = index{cur.x, cur.y - 1}
		if isGoodOrange(grid, next) {
			grid[next.x][next.y] = badOrange
			badOranges = append(badOranges, next)
			nextCnt++
		}

		if cnt <= 0 {
			cnt = nextCnt
			nextCnt = 0
			minutes++
		}

		fmt.Println(cnt, badOranges)
	}

	for _, ints := range grid {
		for _, i3 := range ints {
			if i3 == goodOrange {
				return -1
			}
		}
	}

	// no bad and no good
	if minutes == 0 {
		return 0
	}

	return minutes - 1
}
