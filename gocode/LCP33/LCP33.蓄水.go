package LCP33

import "math"

func storeWater(bucket []int, vat []int) int {
	maxk := 0
	for i := 0; i < len(vat); i++ {
		maxk = getMax(maxk, vat[i])
	}

	if maxk == 0 {
		return 0
	}

	minCnt := math.MaxInt
	for k := 1; k < minCnt; k++ {
		addCntTmp := 0
		for i := 0; i < len(bucket); i++ {
			addCntTmp += getMax(getCeil(vat[i], k)-bucket[i], 0)
		}
		minCnt = getMin(minCnt, k+addCntTmp)
	}

	return minCnt
}

func getCeil(a, b int) int {
	if a%b == 0 {
		return a / b
	}
	return (a + b) / b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
