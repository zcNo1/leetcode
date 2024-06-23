package _210

type point struct {
	x int
	y int
}

type snake struct {
	head point
	tail point
}

const (
	OPEN = iota
	CLOSE
)

func minimumMoves(grid [][]int) int {
	var max = -1
	n := len(grid)
	start := snake{head: point{0, 1}, tail: point{0, 0}}

	move(grid, start, n, false, 0, &max)

	return max
}

func move(grid [][]int, s snake, n int, flag bool, cnt int, max *int) {
	//log.Println(s, cnt)
	if cnt >= *max {
		return
	}
	if s.head.x == n-1 && s.tail.x == n-1 && s.head.y == n-1 && s.tail.y == n-2 {
		if *max == -1 {
			*max = cnt
		} else {
			*max = minNum(*max, cnt)
		}
	}

	if s.head.x+1 < n && s.tail.x+1 < n && grid[s.head.x+1][s.head.y] == OPEN && grid[s.tail.x+1][s.tail.y] == OPEN {
		tmp := s
		tmp.head.x++
		tmp.tail.x++
		move(grid, tmp, n, false, cnt+1, max)
	}

	if s.head.y+1 < n && s.tail.y+1 < n && grid[s.head.x][s.head.y+1] == OPEN && grid[s.tail.x][s.tail.y+1] == OPEN {
		tmp := s
		tmp.head.y++
		tmp.tail.y++
		move(grid, tmp, n, false, cnt+1, max)
	}

	if !flag && s.tail.x+1 < n && s.head.x == s.tail.x && grid[s.tail.x+1][s.tail.y] == OPEN {
		tmp := s
		tmp.head.x = tmp.tail.x + 1
		tmp.head.y = tmp.tail.y
		move(grid, tmp, n, true, cnt+1, max)
	}

	if !flag && s.tail.y+1 < n && s.head.y == s.tail.y && grid[s.tail.x][s.tail.y+1] == OPEN {
		tmp := s
		tmp.head.x = tmp.tail.x
		tmp.head.y = tmp.tail.y + 1
		move(grid, tmp, n, true, cnt+1, max)
	}

	return
}

func minNum(a, b int) int {
	if a > b {
		return b
	}

	return a
}
