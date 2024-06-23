package l1090

import (
	"log"
	"sort"
)

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	cnt := make(map[int][]int)
	for i := 0; i < len(values); i++ {
		cnt[labels[i]] = append(cnt[labels[i]], values[i])
	}

	log.Println(cnt)

	var allVal []int
	for _, ints := range cnt {
		sortMax(ints)
		max := len(ints)
		if max > useLimit {
			max = useLimit
		}
		log.Println(ints, max)
		allVal = append(allVal, ints[:max]...)
		log.Println(allVal)
	}
	sortMax(allVal)
	log.Println(allVal)

	ret := 0
	for i := 0; i < numWanted && i < len(allVal); i++ {
		ret += allVal[i]
	}

	return ret
}

func sortMax(allVal []int) {
	sort.Slice(allVal, func(i, j int) bool {
		if allVal[i] > allVal[j] {
			return true
		}
		return false
	})
}
