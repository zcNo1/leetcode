package l1079

import "sort"


AAB 1 + C 2 1 + C 3 1 =1+ 2 +3=6
A 1
AA 1
AAA 1
AAAB 1 + C 2 1 + C 3 1 + C 4 1
AAABB 1 + C 2 1 * C 3 1 + C 3 1 * C 4 1 + C 4 1 * C 5 1


AAABBC 1 + C 2 1 * C 3 1
   1 2 3 4
A

n个字符=

func numTilePossibilities(tiles string) int {
	tilesByte := []byte(tiles)
	sort.Slice(tilesByte, func(i, j int) bool {
		if i < j {
			return true
		}
		return false
	})


	return 0
}
