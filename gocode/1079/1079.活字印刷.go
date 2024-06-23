package l1079

var res int

func numTilePossibilities(tiles string) int {
	cntTile := make(map[rune]int)

	for _, tile := range tiles {
		cntTile[tile]++
	}

	res = 0

	dfs(cntTile)
	return res
}

func dfs(cntTile map[rune]int) {
	for tile, cnt := range cntTile {
		if cnt > 0 {
			res++
			cntTile[tile]--
			dfs(cntTile)
			cntTile[tile]++
		}
	}
}
