package _129

type colorRecord struct {
	path map[int]int
	maps map[int][]int
}

func buildColorRecord(n int, edges [][]int) *colorRecord {
	var c colorRecord
	c.maps = getMap(edges)
	c.path = make(map[int]int)

	return &c
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	res := make([]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	res[0] = 0

	red := buildColorRecord(n, redEdges)
	blue := buildColorRecord(n, blueEdges)
	goColor(red, blue, res)

	red = buildColorRecord(n, redEdges)
	blue = buildColorRecord(n, blueEdges)
	goColor(blue, red, res)

	return res
}

func getMap(Edges [][]int) map[int][]int {
	maps := make(map[int][]int)
	for _, edge := range Edges {
		maps[edge[0]] = append(maps[edge[0]], edge[1])
	}
	return maps
}

func goColor(color1, color2 *colorRecord, res []int) {
	var color *colorRecord
	rec := []int{0}
	for cur := 1; len(rec) > 0; cur++ {
		if cur%2 == 1 {
			color = color1
		} else {
			color = color2
		}
		tmp := rec
		rec = []int{}
		for _, from := range tmp {
			if color.path[from] == 1 {
				continue
			}
			color.path[from] = 1
			for _, to := range color.maps[from] {
				if res[to] == -1 {
					res[to] = cur
				} else {
					if cur < res[to] {
						res[to] = cur
					}
				}

				rec = append(rec, to)
			}
		}
	}
}
