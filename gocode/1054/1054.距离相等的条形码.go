package l1054

func rearrangeBarcodes(barcodes []int) []int {
	cnts := make(map[int]int, 0)

	for _, barcode := range barcodes {
		cnts[barcode]++
	}

	ret := make([]int, len(barcodes))
	idx := 0
	idy := 1
	for i, cnt := range cnts {
		for cnt > 0 && cnt <= len(barcodes)/2 && idy < len(barcodes) {
			ret[idy] = i
			cnt--
			idy += 2
		}
		for cnt > 0 {
			ret[idx] = i
			cnt--
			idx += 2
		}
	}

	return ret
}
